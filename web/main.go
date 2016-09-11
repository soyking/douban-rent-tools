package main

import (
	"encoding/json"
	"github.com/soyking/douban-rent-tools/storage"
	"github.com/soyking/douban-rent-tools/storage/mongo"
	"io/ioutil"
	"net/http"
	"strconv"
)

var store storage.Storage

func main() {
	store, _ = mongo.NewMongoDBStorage("", "", "", "db_rent")
	http.HandleFunc("/query", queryHandler)
	http.ListenAndServe(":8080", nil)
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
