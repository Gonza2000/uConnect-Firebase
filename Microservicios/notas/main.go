package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Endpoint para obtener todas las notas
	r.HandleFunc("/notes", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Obteniendo todas las notas"))
	}).Methods("GET")

	// Endpoint para crear una nota
	r.HandleFunc("/notes", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Creando una nueva nota"))
	}).Methods("POST")

	log.Println("Microservicio de Notas corriendo en el puerto 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
