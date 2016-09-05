package main

import (
	"github.com/soyking/douban-rent-tools/storage"
	"github.com/soyking/douban-rent-tools/storage/mongo"
	"log"
)

var store storage.Storage

func initStorage() {
	var err error
	store, err = mongo.NewMongoDBStorage(mongoDBAddr, mongoDBUsername, mongoDBPassword, mongoDBDatabase)
	if err != nil {
		log.Fatal(err)
	}
}
