package main

import (
	"database/sql"
	"fmt"
	"log"
)

func fetchPosts(db *sql.DB) ([]PostData, error) {
	sqlStmt := `
        SELECT 
            p.post_id, p.user_id, p.post_title, p.post_category, p.post_text, p.post_date, 
            p.post_likes, p.post_dislikes, p.post_img,
            u.user_name AS post_user_name,
            c.comments_user_id, c.comments_content,
            u2.user_name AS comment_user_name
        FROM 
            posts p
        LEFT JOIN 
            comments c ON p.post_id = c.comments_post_id
        LEFT JOIN
            user u ON p.user_id = u.user_id
        LEFT JOIN
            user u2 ON c.comments_user_id = u2.user_id
        ORDER BY 
            p.post_date DESC
    `

	rows, err := db.Query(sqlStmt)
	if err != nil {
		log.Println("Erreur lors de la récupération des données:", err)
		return nil, fmt.Errorf("erreur lors de la récupération des données: %w", err)
	}
	defer rows.Close()

	var posts []PostData
	var currentPost PostData
	var currentPostID int

	for rows.Next() {
		var postID, userID int
		var postTitle, postText, postDate, postImg string
		var postLikes, postDislikes sql.NullInt64
		var commentsUserID sql.NullInt64
		var commentsContent sql.NullString
		var postCategory, postUserName, commentUserName sql.NullString

		err := rows.Scan(&postID, &userID, &postTitle, &postCategory, &postText, &postDate, &postLikes, &postDislikes, &postImg, &postUserName, &commentsUserID, &commentsContent, &commentUserName)
		if err != nil {
			log.Println("Erreur lors de la lecture des données:", err)
			return nil, fmt.Errorf("erreur lors de la lecture des données: %w", err)
		}

		if currentPostID != postID {
			if currentPostID != 0 {
				posts = append(posts, currentPost)
			}

			currentPostID = postID
			currentPost = PostData{
				PostID:       postID,
				UserID:       userID,
				UserName:     postUserName.String,
				PostTitle:    postTitle,
				PostCategory: postCategory,
				PostText:     postText,
				PostDate:     sql.NullString{String: postDate, Valid: true},
				PostLikes:    postLikes,
				PostDislikes: postDislikes,
				PostImg:      postImg,
				Comments:     []CommentData{},
			}

		}

		// Ajouter le commentaire à la liste des commentaires du post actuel
		if commentsUserID.Valid && commentsContent.Valid {
			currentPost.Comments = append(currentPost.Comments, CommentData{
				UserID:         int(commentsUserID.Int64),
				UserName:       commentUserName.String,
				CommentContent: commentsContent.String,
			})
		}
	}

	// Ajouter le dernier post récupéré
	if currentPostID != 0 {
		posts = append(posts, currentPost)
	}

	if err = rows.Err(); err != nil {
		log.Println("Erreur finale lors de la lecture des données:", err)
		return nil, fmt.Errorf("erreur finale lors de la lecture des données: %w", err)
	}

	return posts, nil
}
