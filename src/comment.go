package main

import (
	"net/http"
	"strconv"
)

// Gestion utilisateur pour commenter
func commentHandler(w http.ResponseWriter, r *http.Request) {
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
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		postID := r.FormValue("post_id")
		content := r.FormValue("comment_content")

		postIDInt, err := strconv.Atoi(postID)
		if err != nil {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
			return
		}

		err = createComment(db, userID, postIDInt, content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

}
