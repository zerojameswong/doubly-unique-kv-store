package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"zerojameswong/kvstore/impl"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	stores, err := impl.GetStores()
	if err != nil {
		fmt.Println(err)
		return
	}
	tmp := template.Must(template.ParseFiles("templates/home.html"))
	tmp.Execute(w, stores)
}
