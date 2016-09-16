package router

import (
	"github.com/soyking/douban-rent-tools/expand"
	"github.com/soyking/douban-rent-tools/flag"
	"net/http"
)

func InitRouter(f *flag.Flag) error {
	store, err := newStorage(f)
	if err != nil {
		return err
	}

	// 关键字扩展
	ep := expand.NewExpander().
		AddExpand(SubwayCond, SubwayExpand).
		AddExpand(RoomCond, RoomExpand)

	http.HandleFunc("/query", queryHandler(store, ep))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/dist/index.html")
	})

	return nil
}
