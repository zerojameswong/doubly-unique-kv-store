package main

import (
	"database/sql"
	"net/http"
	"zerojameswong/kvstore/handlers"
	"zerojameswong/kvstore/impl"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func main() {
	err := impl.InitDB()
	if err != nil {
		panic(err)
	}

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/stores/", handlers.StoreHandler)
	http.ListenAndServe(":8080", nil)
}
