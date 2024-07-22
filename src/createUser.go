package main

import (
	"database/sql"
	"fmt"
)

// createUser insère des données dans la table user
func createUser(db *sql.DB, email, username, password string) error {
    hashedPassword := hashPassword(password)
    sqlStmt := `INSERT INTO user (user_email, user_name, user_password) VALUES (?, ?, ?)`
    _, err := db.Exec(sqlStmt, email, username, hashedPassword)
    if err != nil {
        return fmt.Errorf("erreur lors de l'insertion des données: %w", err)
    }
    return nil
}