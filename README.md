# auto-model

auto-model 是一款简单的数据表模型生成工具，目前只支持mysql（或mariadb）

## 安装

```shell
git clone https://github.com/west3316/auto-model

cd auto-model

go build
```

## 快速开始

1. 在数据库中建表

```sql
-- 示例表 user_charge
-- demo用到test数据库
ALTER DATABASE `test` CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci; 

CREATE TABLE `user_charge` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '唯一ID',
  `user_id` int(10) unsigned NOT NULL COMMENT '用户ID',
  `create_at` datetime NOT NULL COMMENT '充值时间',
  `value` decimal(20,2) NOT NULL COMMENT '充值金额',
  `result` enum('NoResult','Success','Fail','Locked') NOT NULL COMMENT '充值结果',
  `desc` text DEFAULT NULL COMMENT '备注',
  KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=640 DEFAULT CHARSET=utf8mb4

```

2. 在config.toml中设置运行参数

3. 运行生成命令

```shell
./auto-model.exe --config config.toml
```

示例生成代码 model/user_charge.go，也可以直接通过命令选项来，使用方法参看命令帮助 `./auto-model.exe --help`

## 特点

支持enum枚举类型（少得尴尬



