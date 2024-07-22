package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	userID, isLoggedIn, err := isAuthenticated(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if !isLoggedIn {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		title := r.FormValue("title")
		categories := r.FormValue("categories")
		description := r.FormValue("description")

		file, handler, err := r.FormFile("photo")
		if err != nil {
			http.Error(w, "Erreur lors du traitement du fichier", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		photoPath := "./static/uploaded/" + handler.Filename
		err = saveFile(file, photoPath)
		if err != nil {
			http.Error(w, "Erreur lors de la sauvegarde du fichier", http.StatusInternalServerError)
			return
		}

		err = createPost(db, userID, title, categories, description, photoPath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/?status=success", http.StatusSeeOther)
		return
	}

	tmpl := template.Must(template.ParseFiles("./static/post.html"))
	err = tmpl.ExecuteTemplate(w, "post.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func isAuthenticated(r *http.Request) (int, bool, error) {
	cookie, err := r.Cookie("user_id")
	if err != nil {
		return 0, false, err
	}

	var userID int
	_, err = fmt.Sscanf(cookie.Value, "%d", &userID)
	if err != nil {
		return 0, false, fmt.Errorf("ID utilisateur invalide")
	}

	return userID, true, nil
}
