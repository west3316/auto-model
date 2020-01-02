package gen

// makeFieldTags 构造tags
func makeFieldTags(field string, autoInc bool, tags []string) string {
	var text = "`"

	for _, tag := range tags {
		var f string
		warpper := _tagWarpper[tag]
		if warpper != nil {
			f = warpper(field)
		} else {
			f = field
		}

		text += tag + `:"` + f + `" `
	}
	if len(text) != 0 {
		if autoInc {
			text += `mark:"auto_increment"`
		} else {
			text = text[:len(text)-1]
		}
		text += "`"
	}
	return text
}

var (
	_tagWarpper = map[string]func(string) string{
		"gorm": gormFieldWarpper,
	}
)

func gormFieldWarpper(field string) string {
	return "column:" + field
}
