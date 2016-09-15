package main

import (
	"github.com/soyking/douban-rent-tools/storage"
	"github.com/soyking/douban-rent-tools/storage/es"
	"github.com/soyking/douban-rent-tools/storage/mongo"
	"log"
)

var store storage.Storage

func initStorage() {
	var err error
	if mongoDBOn {
		println("STORAGE MONGODB ( " + mongoDBAddr + " )\n")
		store, err = mongo.NewMongoDBStorage(mongoDBAddr, mongoDBUsername, mongoDBPassword, mongoDBDatabase)
	} else {
		println("STORAGE ELASTICSEARCH ( " + esAddr + " )\n")
		store, err = es.NewElasticSearchStorage(esAddr, esIndex)
	}
	if err != nil {
		log.Fatal(err)
	}
}
