package gen

import (
	"log"
	"os/exec"
	"strings"

	"github.com/serenize/snaker"
	"github.com/siddontang/go-mysql/client"
)

var (
	_client *client.Conn
)

// Run _
func Run(args Args) error {
	var err error
	_client, err = client.Connect(args.DB.Address, args.DB.User, args.DB.Password, args.DB.Name)
	if err != nil {
		return err
	}

	err = _client.Ping()
	if err != nil {
		return err
	}

	// 获取需要生成model的表
	exportTables, err := getExportTables(args.DB.Tables)
	if err != nil {
		return err
	}

	for _, table := range exportTables {
		fields, enums, err := getFields(table, args.Tags)
		if err != nil {
			log.Println("查询表“"+table+"”字段失败：", err)
			continue
		}
		data := &Schema{
			PackageName: args.PackageName,
			TableName:   table,
			Fields:      fields,
			StructName:  snaker.SnakeToCamel(table),
			Enums:       enums,
		}
		data.FormatFields()

		filename, err := render(args.DB.Name, args.OutDir, args.Template, data)
		if err != nil {
			log.Println("生成表“"+table+"”模型失败：", err)
			continue
		}

		err = exec.Command("goimports", "-w", filename).Run()
		if err != nil {
			log.Println("格式化“"+table+"”失败：", err)
			continue
		}
	}

	// _client.FieldList()
	return nil
}

func getExportTables(argTables string) ([]string, error) {
	result, err := _client.Execute("SHOW TABLES")
	if err != nil {
		return nil, err
	}

	argTables = strings.Trim(argTables, `'"`)
	var filter map[string]bool
	t := strings.Split(argTables, ",")
	if len(t) != 0 {
		filter = make(map[string]bool, len(t))
		for _, table := range t {
			filter[table] = true
		}
	}

	// 需要生成model的表
	var tables []string
	for _, cells := range result.Values {
		table := string(toString(cells[0]))
		if filter == nil || filter[table] {
			tables = append(tables, table)
		}
	}

	return tables, nil
}

func getFields(table string, tags string) ([]Field, []Enum, error) {
	result, err := _client.Execute("SHOW FULL FIELDS FROM " + table)
	if err != nil {
		return nil, nil, err
	}

	tags = strings.Trim(tags, `'"`)
	theTags := strings.Split(tags, ",")

	const (
		IndexField      = 0
		IndexType       = 1
		IndexCollation  = 2
		IndexNull       = 3
		IndexKey        = 4
		IndexDefault    = 5
		IndexExtra      = 6
		IndexPrivileges = 7
		IndexComment    = 8
	)

	var fields []Field
	var enums []Enum
	for _, cells := range result.Values {
		field := toString(cells[IndexField])
		fields = append(fields, Field{
			Name:       snaker.SnakeToCamel(field),
			Type:       toGoType(toBytes(cells[IndexType]), toString(cells[IndexNull]) == "YES"),
			Comment:    toString(cells[IndexComment]),
			PrimaryKey: toString(cells[IndexKey]) == "PRI",
			Tags:       makeFieldTags(field, toString(cells[IndexExtra]) == "auto_increment", theTags),
		})

		last := fields[len(fields)-1]
		values := enumValues(toString(cells[IndexType]))
		if len(values) != 0 {
			enums = append(enums, Enum{Prefix: last.Name, Comment: last.Comment, Values: values})
		}
	}

	return fields, enums, nil
}
