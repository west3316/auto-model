package gen

import "github.com/serenize/snaker"

// Field 字段
type Field struct {
	Name       string
	Comment    string
	Type       string
	Tags       string
	PrimaryKey bool
}

// Enum 枚举定义
type Enum struct {
	Comment string
	Prefix  string
	Values  []string
}

// Schema 渲染数据
type Schema struct {
	PackageName string
	TableName   string
	StructName  string
	Fields      []Field

	// 运行时生成
	PrimaryKeys []string
	NotPKFields []string

	Enums []Enum
}

// FormatFields 获取字段，按照主键，非主键区分
func (s *Schema) FormatFields() {
	for _, field := range s.Fields {
		if field.PrimaryKey {
			s.PrimaryKeys = append(s.PrimaryKeys, field.Name)
		} else {
			s.NotPKFields = append(s.NotPKFields, field.Name)
		}
	}
}

// SnakeToCamel 命名风格转换
func (s Schema) SnakeToCamel(text string) string {
	return snaker.SnakeToCamel(text)
}

// CamelToSnake 命名风格转换
func (s Schema) CamelToSnake(text string) string {
	return snaker.CamelToSnake(text)
}

// Args 参数
type Args struct {
	DB          DB
	PackageName string
	Template    string
	OutDir      string
	Tags        string
}

// DB 数据库参数
type DB struct {
	Name     string
	Address  string
	User     string
	Password string
	Tables   string
}
