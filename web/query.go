package main

import (
	"encoding/json"
	"github.com/soyking/douban-rent-tools/storage"
	"github.com/soyking/douban-rent-tools/web/expand"
	"io/ioutil"
	"net/http"
)

var ep = expand.NewExpander().
	AddExpand(expand.SubwayCond, expand.SubwayExpand).
	AddExpand(expand.RoomCond, expand.RoomExpand)

func writeErr(w http.ResponseWriter, err error) {
	w.Write([]byte(`{"error":"` + err.Error() + `"}`))
}

type queryResult struct {
	Count  int         `json:"count"`
	Result interface{} `json:"result"`
}

func writeResult(w http.ResponseWriter, count int, result interface{}) {
	b, _ := json.Marshal(queryResult{count, result})
	w.Write(b)
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

	q.Keywords = ep.Expand(q.Keywords)

	count, result, err := store.Query(q)
	if err != nil {
		writeErr(w, err)
		return
	}

	writeResult(w, count, result)
}
