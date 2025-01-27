package main

import (
	"blog_api/config"
	"blog_api/models"
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	db := config.Connect(ctx)
	article := models.Article{
		Title:         "Go Programming Basics",
		Content:       "This article introduces basic concepts of Go programming.",
		PublishedDate: time.Now(), 
		Tags:          []string{"Go", "Programming", "Tutorial"},
	}

	models.CreateArticle(ctx, db, article)
	articles := models.FetchAllArticle(ctx, db)
	fmt.Println(articles)
	defer db.Close(ctx)
}
