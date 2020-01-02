package gen

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

var (
	_tpl *template.Template
)

func getTemplate(db, tplFile string) (*template.Template, error) {
	if _tpl != nil {
		return _tpl, nil
	}

	data, err := ioutil.ReadFile(tplFile)
	if err != nil {
		return nil, errors.New("打开模板文件“" + tplFile + "”失败：" + err.Error())
	}

	_tpl = template.New(db)
	_tpl, err := _tpl.Parse(string(data))
	if err != nil {
		return nil, errors.New("解析模板文件“" + tplFile + "”失败：" + err.Error())
	}

	// _tpl = _tpl.Funcs(template.FuncMap{"SnakeToCamel": func(text string) string {
	// 	return snaker.SnakeToCamel(text)
	// }})

	return _tpl, nil
}

func render(db, outDir, tplFile string, data *Schema) (string, error) {
	tpl, err := getTemplate(db, tplFile)
	if err != nil {
		return "", err
	}

	filename := filepath.Join(outDir, data.TableName+".go")
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	defer file.Close()
	if err != nil {
		return "", errors.New("创建文件“" + filename + "”失败：" + err.Error())
	}

	err = tpl.Execute(file, data)
	if err != nil {
		return "", errors.New("渲染文件“" + filename + "”失败：" + err.Error())
	}

	return filename, nil
}
