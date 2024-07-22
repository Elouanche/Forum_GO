package main

import (
	"database/sql"
	"fmt"
	"time"
)

func createPost(db *sql.DB, userID int, title, categories, description, photoPath string) error {
	currentDateTime := time.Now().Format("02-01-2006 15:04:05")
	sqlStmt := `INSERT INTO posts (user_id, post_title, post_category, post_text, post_date, post_img) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(sqlStmt, userID, title, categories, description, currentDateTime, photoPath)
	if err != nil {
		return fmt.Errorf("erreur lors de l'insertion des données: %w", err)
	}
	return nil
}
