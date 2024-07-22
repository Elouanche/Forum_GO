package main

import (
	"database/sql"
	"fmt"
)

// Ajout de la table reactions
// initDB crée la table user si elle n'existe pas déjà
func initDB(db *sql.DB) error {
	sqlStmt := `
    CREATE TABLE IF NOT EXISTS user (
        user_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        user_email TEXT UNIQUE NOT NULL,
        user_name TEXT UNIQUE NOT NULL,
        user_password TEXT
    );
    
	CREATE TABLE IF NOT EXISTS posts (
        post_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER NOT NULL,
        post_title TEXT,
        post_text TEXT,
		post_date TEXT,
		post_likes INTEGER,
		post_dislikes INTEGER,
        post_img TEXT,
        post_category TEXT,
        FOREIGN KEY (user_id) REFERENCES user(user_id)
    );

    CREATE TABLE IF NOT EXISTS comments (
        comments_id INTEGER PRIMARY KEY AUTOINCREMENT,
        comments_post_id INTEGER,
        comments_user_id INTEGER,
        comments_content TEXT,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );
    

    CREATE TABLE IF NOT EXISTS reactions (
			like_id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			post_id INTEGER NOT NULL,
            reaction_type TEXT CHECK(reaction_type IN ('like', 'dislike')),
			UNIQUE(user_id, post_id),
			FOREIGN KEY (user_id) REFERENCES user(user_id),
			FOREIGN KEY (post_id) REFERENCES posts(post_id)
	);
    `
	_, err := db.Exec(sqlStmt)
	if err != nil {
		return fmt.Errorf("erreur lors de la création des tables: %w", err)
	}
	return nil
}
