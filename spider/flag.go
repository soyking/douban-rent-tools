package main

import "flag"

const (
	FLAG_GROUPS_NAME    = "groups"
	FLAG_GROUPS_DEFAULT = "beijingzufang"
	FLAG_GROUPS_USAGE   = "group name, split by ,"

	FLAG_MONGO_ADDR_NAME    = "mg_addr"
	FLAG_MONGO_ADDR_DEFAULT = "127.0.0.1:27017"
	FLAG_MONGO_ADDR_USAGE   = "MongoDB address, split by ,"

	FLAG_MONGO_USERNAME_NAME    = "mg_usr"
	FLAG_MONGO_USERNAME_DEFAULT = ""
	FLAG_MONGO_USERNAME_USAGE   = "MongoDB username"

	FLAG_MONGO_PASSWORD_NAME    = "mg_pwd"
	FLAG_MONGO_PASSWORD_DEFAULT = ""
	FLAG_MONGO_PASSWORD_USAGE   = "MongoDB password"

	FLAG_MONGO_DATABASE_NAME    = "mg_db"
	FLAG_MONGO_DATABASE_DEFAULT = "db_rent"
	FLAG_MONGO_DATABASE_USAGE   = "MongoDB database"

	FLAG_FREQUENCY_NAME    = "freq"
	FLAG_FREQUENCY_DEFUALT = 60
	FLAG_FREQUENCY_USAGE   = "spider frequency(in second)"

	FLAG_GROUPS_CONCURRENCY_NAME    = "g_con"
	FLAG_GROUPS_CONCURRENCY_DEFAULT = 1
	FLAG_GROUPS_CONCURRENCY_USAGE   = "concurrency for groups crawling"

	FLAG_TOPICS_CONCURRENCY_NAME    = "t_con"
	FLAG_TOPICS_CONCURRENCY_DEFAULT = 1
	FLAG_TOPICS_CONCURRENCY_USAGE   = " for topics crawling"
)

var (
	groupNames        string
	mongoDBAddr       string
	mongoDBUsername   string
	mongoDBPassword   string
	mongoDBDatabase   string
	frequency         int
	groupsConcurrency int
	topicsConcurrency int
)

func init() {
	flag.StringVar(&groupNames, FLAG_GROUPS_NAME, FLAG_GROUPS_DEFAULT, FLAG_GROUPS_USAGE)
	flag.StringVar(&mongoDBAddr, FLAG_MONGO_ADDR_NAME, FLAG_MONGO_ADDR_DEFAULT, FLAG_MONGO_ADDR_USAGE)
	flag.StringVar(&mongoDBUsername, FLAG_MONGO_USERNAME_NAME, FLAG_MONGO_USERNAME_DEFAULT, FLAG_MONGO_USERNAME_USAGE)
	flag.StringVar(&mongoDBPassword, FLAG_MONGO_PASSWORD_NAME, FLAG_MONGO_PASSWORD_DEFAULT, FLAG_MONGO_PASSWORD_USAGE)
	flag.StringVar(&mongoDBDatabase, FLAG_MONGO_DATABASE_NAME, FLAG_MONGO_DATABASE_DEFAULT, FLAG_MONGO_DATABASE_USAGE)
	flag.IntVar(&frequency, FLAG_FREQUENCY_NAME, FLAG_FREQUENCY_DEFUALT, FLAG_FREQUENCY_USAGE)
	flag.IntVar(&groupsConcurrency, FLAG_GROUPS_CONCURRENCY_NAME, FLAG_GROUPS_CONCURRENCY_DEFAULT, FLAG_GROUPS_CONCURRENCY_USAGE)
	flag.IntVar(&topicsConcurrency, FLAG_TOPICS_CONCURRENCY_NAME, FLAG_TOPICS_CONCURRENCY_DEFAULT, FLAG_TOPICS_CONCURRENCY_USAGE)
	flag.Parse()

	if frequency <= 0 {
		frequency = FLAG_FREQUENCY_DEFUALT
	}
	if groupsConcurrency <= 0 {
		groupsConcurrency = FLAG_GROUPS_CONCURRENCY_DEFAULT
	}
	if topicsConcurrency <= 0 {
		groupsConcurrency = FLAG_TOPICS_CONCURRENCY_DEFAULT
	}
}
