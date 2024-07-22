package main

import (
	"net/http"
	"strings"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		status := r.URL.Query().Get("status")
		data := struct {
			Status string
		}{
			Status: status,
		}
		err := tmpl.ExecuteTemplate(w, "create.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		email := r.FormValue("email")
		username := r.FormValue("username")
		password := r.FormValue("password")

		err = createUser(db, email, username, password)
		if err != nil {
			if strings.Contains(err.Error(), "UNIQUE constraint failed: user.user_name") {
				http.Redirect(w, r, "/create?status=username_taken", http.StatusSeeOther)
				return
			} else if strings.Contains(err.Error(), "UNIQUE constraint failed: user.user_email") {
				http.Redirect(w, r, "/create?status=email_taken", http.StatusSeeOther)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

