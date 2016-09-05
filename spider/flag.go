package main

import "flag"

const (
	FLAG_GROUPS_NAME    = "group"
	FLAG_GROUPS_DEFAULT = "beijingzufang"
	FLAG_GROUPS_USAGE   = "小组名，多个用逗号分隔开"

	FLAG_MONGO_ADDR_NAME    = "mg_addr"
	FLAG_MONGO_ADDR_DEFAULT = "127.0.0.1:27017"
	FLAG_MONGO_ADDR_USAGE   = "MongoDB 地址，多个用逗号分隔开"

	FLAG_MONGO_USERNAME_NAME    = "mg_usr"
	FLAG_MONGO_USERNAME_DEFAULT = ""
	FLAG_MONGO_USERNAME_USAGE   = "MongoDB 用户名"

	FLAG_MONGO_PASSWORD_NAME    = "mg_pwd"
	FLAG_MONGO_PASSWORD_DEFAULT = ""
	FLAG_MONGO_PASSWORD_USAGE   = "MongoDB 密码"

	FLAG_MONGO_DATABASE_NAME    = "mg_db"
	FLAG_MONGO_DATABASE_DEFAULT = "db_rent"
	FLAG_MONGO_DATABASE_USAGE   = "MongoDB 数据库名称"

	FLAG_FREQUENCY_NAME    = "freq"
	FLAG_FREQUENCY_DEFUALT = 60
	FLAG_FREQUENCY_USAGE   = "抓取频率，单位秒"
)

var (
	groupNames      string
	mongoDBAddr     string
	mongoDBUsername string
	mongoDBPassword string
	mongoDBDatabase string
	frequency       int
)

func init() {
	flag.StringVar(&groupNames, FLAG_GROUPS_NAME, FLAG_GROUPS_DEFAULT, FLAG_GROUPS_USAGE)
	flag.StringVar(&mongoDBAddr, FLAG_MONGO_ADDR_NAME, FLAG_MONGO_ADDR_DEFAULT, FLAG_MONGO_ADDR_USAGE)
	flag.StringVar(&mongoDBUsername, FLAG_MONGO_USERNAME_NAME, FLAG_MONGO_USERNAME_DEFAULT, FLAG_MONGO_USERNAME_USAGE)
	flag.StringVar(&mongoDBPassword, FLAG_MONGO_PASSWORD_NAME, FLAG_MONGO_PASSWORD_DEFAULT, FLAG_MONGO_PASSWORD_USAGE)
	flag.StringVar(&mongoDBDatabase, FLAG_MONGO_DATABASE_NAME, FLAG_MONGO_DATABASE_DEFAULT, FLAG_MONGO_DATABASE_USAGE)
	flag.IntVar(&frequency, FLAG_FREQUENCY_NAME, FLAG_FREQUENCY_DEFUALT, FLAG_FREQUENCY_USAGE)
	flag.Parse()

	if frequency <= 0 {
		frequency = FLAG_FREQUENCY_DEFUALT
	}
}
