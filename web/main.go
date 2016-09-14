package main

import (
	"encoding/json"
	"github.com/soyking/douban-rent-tools/storage"
	"github.com/soyking/douban-rent-tools/storage/es"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var store storage.Storage

func main() {
	var err error
	store, err = es.NewElasticSearchStorage("127.0.0.1:9200", "db_rent")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/query", queryHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/src/index.html")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
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
