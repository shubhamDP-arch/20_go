package models

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type Article struct {
	ID            int      `json:"id"`
	Title         string   `json:"title"`
	Content       string   `json:"content"`
	PublishedDate time.Time   `json:"published_date"`
	Tags          []string `json:"tags"`
}

// func CreateTableArticle(conn *pgx.Conn) {
// 	query := `
// 		CREATE TABLE IF NOT EXISTS articles(
// 			id SERIAL PRIMARY KEY,
// 			title TEXT NOT NULL,
// 			content TEXT NOT NULL,
// 			published_date DATE NOT NULL,
// 			tags TEXT[] NOT NULL
// 		)
// 	`
// 	ctx := context.Background()
// 	_, err := conn.Exec(ctx, query)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Table 'articles' created successfull")
// }

func CreateArticle(ctx context.Context, conn *pgx.Conn, article Article) error {
	query := `
		INSERT INTO articles (title, content, published_date, tags)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
		`
	var id int
	err := conn.QueryRow(ctx, query, article.Title, article.Content, article.PublishedDate, article.Tags).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Article create with ID: %d\n", id)
	return nil
}

func FetchAllArticle(ctx context.Context, conn *pgx.Conn) []Article {
	query := `
		SELECT id, title, content, published_date, tags FROM articles
	`
	rows, err := conn.Query(ctx, query)

	if err != nil {
		panic(err)
	}
	var articles []Article
	for rows.Next() {
		var article Article
		err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.PublishedDate, &article.Tags)
		if err != nil {
			panic(err)
		}
		articles = append(articles, article)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
	defer rows.Close()
	return articles
}
