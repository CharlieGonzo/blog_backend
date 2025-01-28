package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("en.env")
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
	// insert
	// _, err = conn.Exec(context.Background(), "INSERT INTO playing_with_neon(name, value) SELECT LEFT(md5(i::TEXT), 10), random() FROM generate_series(1, 10) s(i);")
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
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
