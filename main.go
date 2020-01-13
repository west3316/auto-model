// package main MySQL 数据表model生成器
package main

import (
	"flag"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/west3316/auto-model/gen"
)

func main() {
	var config string
	var args gen.Args
	flag.StringVar(&config, "config", "", "从配置文件中读取参数，例如：config.toml")
	flag.StringVar(&args.PackageName, "package-name", "model", "包名")
	flag.StringVar(&args.DB.Address, "db-address", "localhost:3306", "数据库主机地址")
	flag.StringVar(&args.DB.User, "db-user", "root", "数据库用户名")
	flag.StringVar(&args.DB.Password, "db-password", "", "数据库密码")
	flag.StringVar(&args.DB.Name, "db-name", "mysql", "数据库名")
	flag.StringVar(&args.DB.Tables, "db-tables", "", `指定表，格式：["table_name1","table_name2","table_nameN"]`)
	flag.StringVar(&args.Template, "template", "example.tpl", "模板文件名称")
	flag.StringVar(&args.Tags, "tags", "db", `golang结构体tags，格式：["gorm","db","json","yaml","toml"]`)
	flag.StringVar(&args.OutDir, "out-dir", "model", "输出目录")

	flag.Parse()

	if config != "" {
		_, err := toml.DecodeFile(config, &args)
		if err != nil {
			log.Fatalln("解析“"+config+"”配置失败：", err)
		}
	}

	if !gen.DirExists(args.OutDir) {
		err := os.MkdirAll(args.OutDir, 0600)
		if err != nil {
			log.Fatalln("创建“"+args.OutDir+"”输出目录失败：", err)
		}
	}

	err := gen.Run(args)
	if err != nil {
		log.Fatalln("生成model文件失败：", err)
	}

	log.Println("model文件生成完毕：", args.OutDir)
}
