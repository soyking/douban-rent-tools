package main

import (
	"github.com/soyking/douban-rent-tools/flag"
	"github.com/soyking/douban-rent-tools/router"
	"log"
	"net/http"
)

const (
	APP_NAME    = "DOUBAN RENT TOOLS - WEB"
	APP_VERSION = "0.0.1"
)

func main() {
	f := flag.ParseFlag()
	err := router.InitRouter(f)
	if err != nil {
		log.Fatal(err)
	}

	println(APP_NAME + "\t" + APP_VERSION)
	log.Printf("listen on " + f.Port)
	log.Fatal(http.ListenAndServe(":"+f.Port, nil))
}
