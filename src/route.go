package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func route(h1, h2, h3, h4, h5, h6, h7 http.HandlerFunc) {
	r := mux.NewRouter()

	r.HandleFunc("/", h1)
	r.HandleFunc("/about", h2)
	r.HandleFunc("/login", h3)
	r.HandleFunc("/create", h4)
	r.HandleFunc("/logout", h5)
	r.HandleFunc("/post", h6)
	r.HandleFunc("/comment", h7)

	// Ajouter la route pour la réaction aux posts
	r.HandleFunc("/posts/{id:[0-9]+}/react", reactToPost).Methods("POST")

	// Servir les fichiers statiques depuis le répertoire "static"
	fs := http.FileServer(http.Dir("static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	ClearConsole()
	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
