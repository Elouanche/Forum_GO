package main

import (
	"log"
	"net/http"
)

func handlerAccueil(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		handleFormData(r)
		return
	}

	cookie, err := r.Cookie("user_id")
	isLoggedIn := err == nil && cookie != nil && cookie.Value != ""

	posts, err := fetchPosts(db)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des publications", http.StatusInternalServerError)
		log.Printf("Erreur lors de la récupération des publications: %v", err)
		return
	}

	data := PageData{
		IsLoggedIn: isLoggedIn,
		Posts:      posts,
	}

	err = tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, "Erreur lors du chargement de la page", http.StatusInternalServerError)
		log.Printf("Erreur lors du chargement de la page: %v", err)
		return
	}
}
