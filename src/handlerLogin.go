package main

import (
	"fmt"
	"net/http"
	"time"
)

func handlerLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		valid, userID, err := checkUser(username, password)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erreur interne: %v", err), http.StatusInternalServerError)
			return
		}
		if valid {
			cookie := &http.Cookie{
				Name:    "user_id",
				Value:   fmt.Sprintf("%d", userID),
				Path:    "/",
				Expires: time.Now().Add(24 * time.Hour),
			}
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/login?status=failure", http.StatusSeeOther)
		}
		return
	}
	status := r.URL.Query().Get("status")
	data := struct {
		Status string
	}{
		Status: status,
	}
	err := tmpl.ExecuteTemplate(w, "login.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
