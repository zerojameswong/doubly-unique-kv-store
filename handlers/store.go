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

	entriesPerCategory := make(map[string][]model.Entry)
	for _, entry := range entries {
		entriesPerCategory[entry.Category] = append(entriesPerCategory[entry.Category], entry)
	}

	data := struct {
		Store              model.Store
		EntriesPerCategory map[string][]model.Entry
	}{
		*store,
		entriesPerCategory,
	}

	tmp := template.Must(template.ParseFiles("templates/store.html"))
	tmp.Execute(w, data)
}
