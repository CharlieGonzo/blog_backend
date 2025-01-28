package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type post struct {
	ID          string `json:"id"`
	PostName    string `json:"post_name"`
	Summary     string `json:"summary"`
	PostContent string `json:"post_content"`
	Date        string `json:"date"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connStr := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())
	_, err = conn.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS blogs(id SERIAL PRIMARY KEY, post_name TEXT NOT NULL, summary TEXT NOT NULL, post_text TEXT NOT NULL,date TEXT NOT NULL  );")
	if err != nil {
		panic(err)
	}

	// // insert
	// _, err = conn.Exec(context.Background(), "INSERT INTO blogs (post_name, summary, post_text, date) VALUES ('Intro', 'This is an example summary', 'This is the post''s content' , '01/28/2025' );")
	// if err != nil {
	// 	panic(err)
	// }

	//get blogs
	// rows, err := conn.Query(context.Background(), "SELECT * FROM playing_with_neon")
	// if err != nil {
	// 	panic(err)
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	var id int32
	// 	var name string
	// 	var value float32
	// 	if err := rows.Scan(&id, &name, &value); err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Printf("%d | %s | %f\n", id, name, value)
	// }
	r := gin.Default()
	r.POST("/createPost", func(c *gin.Context) {
		var new_post post

		if err := c.BindJSON(&new_post); err != nil {
			return
		}

		var sqlStatement string = fmt.Sprintf("INSERT INTO blogs(post_name , summary, post_text, date) VALUES ('%s','%s','%s','%s')", new_post.PostName, new_post.Summary, new_post.PostContent, new_post.Date)
		_, err = conn.Exec(context.Background(), sqlStatement)
		if err != nil {
			panic(err)
		}
		c.IndentedJSON(201, new_post)

	})
	r.GET("/getPosts", func(c *gin.Context) {
		rows, err := conn.Query(context.Background(), "SELECT * FROM blogs")
		posts := []post{}
		for rows.Next() {
			var id float32
			var post_name string
			var summary string
			var post_content string
			var date string
			if err := rows.Scan(&id, &post_name, &summary, &post_content, &date); err != nil {
				panic(err)
			}
			new_post := post{fmt.Sprintf("%f", id), post_name, summary, post_content, date}

			posts = append(posts, new_post)
			fmt.Printf("%f | %s | %s\n", id, post_name, post_content)
		}
		if err != nil {
			panic(err)
		}
		fmt.Print(posts)
		defer rows.Close()

		c.JSON(200, gin.H{
			"posts": posts,
		})
	})
	r.Run()
}
