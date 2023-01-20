package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"zerojameswong/kvstore/impl"
	"zerojameswong/kvstore/model"
)

func StoreHandler(w http.ResponseWriter, r *http.Request) {
	storeIdStr := r.URL.Path[len("/stores/"):]
	storeId, err := strconv.ParseInt(storeIdStr, 10, 64)
	if err != nil {
		log.Println(err)
		return
	}
	entries, err := impl.GetEntriesForStore(storeId)
	if err != nil {
		log.Println(err)
		return
	}
	store, err := impl.GetStoreById(storeId)
	if err != nil {
		log.Println(err)
		return
	}

	data := struct {
		Store   model.Store
		Entries []model.Entry
	}{
		*store,
		entries,
	}

	tmp := template.Must(template.ParseFiles("templates/store.html"))
	tmp.Execute(w, data)
}
