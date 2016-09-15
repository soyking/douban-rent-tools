package main

import (
	"log"
	"net/http"
)

const (
	APP_NAME    = "DOUBAN RENT TOOLS - WEB"
	APP_VERSION = "0.0.1"
)

func main() {
	println(APP_NAME + "\t" + APP_VERSION)
	initStorage()

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/query", queryHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/dist/index.html")
	})

	log.Printf("listen on " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
