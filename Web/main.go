package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("blog.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	t.Execute(w, "home_page")
}

func HandleRec() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", home_page).Methods("GET")
	http.Handle("/", rtr)

	http.ListenAndServe(":8080", nil)

}

func main() {
	HandleRec()
}
