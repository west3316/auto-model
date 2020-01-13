package gen

import (
	"os"
	"regexp"
	"strings"
)

// DirExists 检查目录是否存在
func DirExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// toString mysql返回转为字符串
func toString(v interface{}) string {
	return string(v.([]byte))
}

// toString mysql返回转为字节数组
func toBytes(v interface{}) []byte {
	return v.([]byte)
}

// toGoType mysql类型转为go类型
func toGoType(typ []byte, null bool) string {
	reg := regexp.MustCompile(`^\w+`)
	typ = reg.Find(typ)
	var goType string
	switch string(typ) {
	case "char", "varchar", "text", "enum", "set", "json":
		goType = "string"
	case "binary", "varbinary", "blob":
		goType = "[]byte"
	case "tinyint", "smallint", "mediumint", "int":
		goType = "int"
	case "bigint":
		goType = "int64"
	case "decimal", "float", "double":
		goType = "float64"
	case "real":
		goType = "unsupport"
	case "date", "time", "datetime", "timestamp":
		goType = "time.Time"
	case "bit":
		goType = "bool"
	default:
		goType = "unknown"
	}

	if null {
		goType = "*" + goType
	}
	return goType
}

// enumValues 取得枚举类型的取值范围
func enumValues(typ string) []string {
	if !strings.HasPrefix(typ, "enum") {
		return nil
	}

	return strings.Split(strings.TrimRight(strings.TrimLeft(string(typ), `enum('`), "')"), `','`)
}
