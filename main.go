package main

import (
	"github.com/soyking/douban-rent-tools/expand"
	"github.com/soyking/douban-rent-tools/flag"
	"github.com/soyking/douban-rent-tools/server"
	"log"
)

const (
	APP_NAME    = "DOUBAN RENT TOOLS - WEB"
	APP_VERSION = "0.0.1"
)

func main() {
	println(APP_NAME + "\t" + APP_VERSION)

	f := flag.ParseFlag()
	store, err := newStorage(f)
	if err != nil {
		log.Fatal(err)
	}

	ep := expand.NewExpander().
		AddExpand(SubwayCond, SubwayExpand).
		AddExpand(RoomCond, RoomExpand)

	svr, err := server.NewServer(":"+f.Port, store, ep)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("listen on " + f.Port)
	log.Fatal(svr.ListenAndServe())
}
