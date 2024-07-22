package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func reactToPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	reactionType := r.FormValue("reaction_type")
	if reactionType != "like" && reactionType != "dislike" {
		http.Error(w, "Invalid reaction type", http.StatusBadRequest)
		return
	}

	userID, isLoggedIn, err := isAuthenticated(r)
	if err != nil {
		http.Error(w, "Authentication error", http.StatusUnauthorized)
		return
	}

	if !isLoggedIn {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	log.Printf("Post ID: %d, Reaction Type: %s, User ID: %d", postID, reactionType, userID)
	tx, err := db.Begin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Insérer la réaction dans la table reactions
	_, err = tx.Exec("INSERT INTO reactions (user_id, post_id, reaction_type) VALUES (?, ?, ?)", userID, postID, reactionType)
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error inserting reaction into reactions table: %s", err.Error())
		return
	}

	// Mettre à jour le compteur de likes ou dislikes dans la table posts
	column := "post_likes"
	if reactionType == "dislike" {
		column = "post_dislikes"
	}

	_, err = tx.Exec(fmt.Sprintf("UPDATE posts SET %s = %s + 1 WHERE post_id = ?", column, column), postID)
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error updating post likes/dislikes: %s", err.Error())
		return
	}

	err = tx.Commit()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"success":  true,
		"likes":    getPostLikes(postID),
		"dislikes": getPostDislikes(postID),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Ajoutez ces fonctions pour récupérer les likes et dislikes actuels d'un post
func getPostLikes(postID int) int {
	var likes int
	err := db.QueryRow("SELECT post_likes FROM posts WHERE post_id = ?", postID).Scan(&likes)
	if err != nil {
		log.Printf("Error fetching post likes: %s", err.Error())
		return 0
	}
	return likes
}

func getPostDislikes(postID int) int {
	var dislikes int
	err := db.QueryRow("SELECT post_dislikes FROM posts WHERE post_id = ?", postID).Scan(&dislikes)
	if err != nil {
		log.Printf("Error fetching post dislikes: %s", err.Error())
		return 0
	}
	return dislikes
}
