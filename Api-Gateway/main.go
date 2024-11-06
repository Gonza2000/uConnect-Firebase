package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Rutas para el microservicio de notas
	r.HandleFunc("/notes", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "http://localhost:8081/notes", http.StatusTemporaryRedirect)
	}).Methods("GET")

	r.HandleFunc("/notes", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "http://localhost:8081/notes", http.StatusTemporaryRedirect)
	}).Methods("POST")

	log.Println("API Gateway corriendo en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
