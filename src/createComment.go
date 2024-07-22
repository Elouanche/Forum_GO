package main

import (
	"database/sql"
	"fmt"
)

func createComment(db *sql.DB, userID, postID int, content string) error {
	sqlStmt := `INSERT INTO comments (comments_post_id, comments_user_id, comments_content) VALUES (?, ?, ?)`
	_, err := db.Exec(sqlStmt, postID, userID, content)
	if err != nil {
		return fmt.Errorf("erreur lors de l'insertion du commentaire: %w", err)
	}
	return nil
}
