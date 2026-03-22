package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var tmpl = template.Must(template.ParseFiles("core/html/home.html"))

func main() {
	router := mux.NewRouter()

	// Homepage
	router.HandleFunc("/", homeHandler).Methods("GET")

	// Serve CSS files from core/css at /static/
	router.PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("core/css"))))

	// Example API route
	router.HandleFunc("/users", getUsers).Methods("GET")

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("users endpoint"))
}
