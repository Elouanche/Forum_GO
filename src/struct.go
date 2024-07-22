package main

import "database/sql"

type CommentData struct {
	UserID         int
	UserName       string
	CommentContent string
}

// Modifier la structure PostData pour inclure les commentaires
type PostData struct {
	PostID       int
	UserID       int
	UserName     string
	PostTitle    string
	PostCategory sql.NullString
	PostText     string
	PostDate     sql.NullString
	PostLikes    sql.NullInt64
	PostDislikes sql.NullInt64
	PostImg      string
	Comments     []CommentData
}

type PageData struct {
	IsLoggedIn bool
	User       UserData // Ajouter ce champ pour stocker les informations de l'utilisateur
	Posts      []PostData
}

type UserData struct {
	UserID   int
	Email    string
	UserName string
}
