package flag

import (
	"flag"
)

const (
	// ==== Web Setting ====
	FLAG_PORT_NAME    = "port"
	FLAG_PORT_DEFAULT = "8080"
	FLAG_PORT_USAGE   = "listen port"

	// ==== ElasticSearch Setting ====
	FLAG_ES_ADDR_NAME    = "es_addr"
	FLAG_ES_ADDR_DEFAULT = "127.0.0.1:9200"
	FLAG_ES_ADDR_USAGE   = "es address"

	FLAG_ES_INDEX_NAME    = "es_index"
	FLAG_ES_INDEX_DEFAULT = "db_rent"
	FLAG_ES_INDEX_USAGE   = "es index"

	// ==== MongoDB Setting ====
	FLAG_USE_MONGO_NAME    = "mongo"
	FLAG_USE_MONGO_DEFAULT = false
	FLAG_USE_MONGO_USAGE   = "use mongo storage"

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
)

type Flag struct {
	Port string

	EsAddr  string
	EsIndex string

	MongoDBOn       bool
	MongoDBAddr     string
	MongoDBUsername string
	MongoDBPassword string
	MongoDBDatabase string
}

func ParseFlag() *Flag {
	f := new(Flag)

	flag.StringVar(&f.Port, FLAG_PORT_NAME, FLAG_PORT_DEFAULT, FLAG_PORT_USAGE)

	flag.StringVar(&f.EsAddr, FLAG_ES_ADDR_NAME, FLAG_ES_ADDR_DEFAULT, FLAG_ES_ADDR_USAGE)
	flag.StringVar(&f.EsIndex, FLAG_ES_INDEX_NAME, FLAG_ES_INDEX_DEFAULT, FLAG_ES_INDEX_USAGE)

	flag.BoolVar(&f.MongoDBOn, FLAG_USE_MONGO_NAME, FLAG_USE_MONGO_DEFAULT, FLAG_USE_MONGO_USAGE)
	flag.StringVar(&f.MongoDBAddr, FLAG_MONGO_ADDR_NAME, FLAG_MONGO_ADDR_DEFAULT, FLAG_MONGO_ADDR_USAGE)
	flag.StringVar(&f.MongoDBUsername, FLAG_MONGO_USERNAME_NAME, FLAG_MONGO_USERNAME_DEFAULT, FLAG_MONGO_USERNAME_USAGE)
	flag.StringVar(&f.MongoDBPassword, FLAG_MONGO_PASSWORD_NAME, FLAG_MONGO_PASSWORD_DEFAULT, FLAG_MONGO_PASSWORD_USAGE)
	flag.StringVar(&f.MongoDBDatabase, FLAG_MONGO_DATABASE_NAME, FLAG_MONGO_DATABASE_DEFAULT, FLAG_MONGO_DATABASE_USAGE)

	flag.Parse()
	return f
}
