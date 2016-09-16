package task

import (
	"github.com/soyking/douban-rent-tools/spider/flag"
	"github.com/soyking/douban-rent-tools/storage"
	"github.com/soyking/douban-rent-tools/storage/es"
	"github.com/soyking/douban-rent-tools/storage/mongo"
)

func NewStorage(f *flag.Flag) (storage.StorageSave, error) {
	var store storage.StorageSave
	var err error
	if f.MongoDBOn {
		println("STORAGE MONGODB ( " + f.MongoDBAddr + " )\n")
		store, err = mongo.NewMongoDBStorage(f.MongoDBAddr, f.MongoDBUsername, f.MongoDBPassword, f.MongoDBDatabase)
	} else {
		println("STORAGE ELASTICSEARCH ( " + f.EsAddr + " )\n")
		store, err = es.NewElasticSearchStorage(f.EsAddr, f.EsIndex)
	}
	return store, err
}
