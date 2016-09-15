package main

import (
	"encoding/json"
	"github.com/soyking/douban-rent-tools/storage"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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

func writeErr(w http.ResponseWriter, err error) {
	w.Write([]byte(`{"error":"` + err.Error() + `"}`))
}

func writeResult(w http.ResponseWriter, count int, result interface{}) {
	b, _ := json.Marshal(result)
	w.Write([]byte(`
		{
			"count":` + strconv.Itoa(count) + `,
			"result":` + string(b) + `
		}`,
	))
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeErr(w, err)
		return
	}

	q := new(storage.QueryRequest)
	err = json.Unmarshal(body, q)
	if err != nil {
		writeErr(w, err)
		return
	}

	count, result, err := store.Query(q)
	if err != nil {
		writeErr(w, err)
		return
	}

	writeResult(w, count, result)
}
