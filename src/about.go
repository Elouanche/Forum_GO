package main

import "net/http"

//Permets de charger la page about
func handlerAbout(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		handleFormData(r)
		return
	}

	err := tmpl.ExecuteTemplate(w, "about.html", nil)
	if err != nil {
		http.Error(w, "Fail dans le chargement de la page", http.StatusInternalServerError)
		return
	}
}
