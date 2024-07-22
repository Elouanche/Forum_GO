package main

import (
	"database/sql"
	"fmt"
)

func checkUser(username, password string) (bool, int, error) {
	sqlStmt := `SELECT user_id, user_password FROM user WHERE user_name = ?`
	var userID int
	var storedHashedPassword string
	err := db.QueryRow(sqlStmt, username).Scan(&userID, &storedHashedPassword)
	if err == sql.ErrNoRows {
		return false, 0, nil
	} else if err != nil {
		return false, 0, fmt.Errorf("erreur lors de la vérification des informations de connexion: %w", err)
	}

	if hashPassword(password) != storedHashedPassword {
		return false, 0, nil
	}

	return true, userID, nil
}
