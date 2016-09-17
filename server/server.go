package server

import (
	"errors"
	"github.com/soyking/douban-rent-tools/expand"
	"github.com/soyking/douban-rent-tools/storage"
	"net/http"
)

var (
	ErrorNoStore = errors.New("no store")
)

func NewServer(addr string, store storage.StorageQuery, ep *expand.Expander) (*http.Server, error) {
	if store == nil {
		return nil, ErrorNoStore
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/query", queryHandler(store, ep))
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./static"))))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/dist/index.html")
	})
	return &http.Server{Addr: addr, Handler: mux}, nil
}
