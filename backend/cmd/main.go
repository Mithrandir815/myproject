package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	// MySQLデータベースに接続します
	db, err := sql.Open("mysql", "root:password@tcp(db:3306)/my_database")
	if err != nil {
		log.Fatal("Failed to connect to MySQL:", err)
	}
	defer db.Close()

	// Ginのルーターを初期化します
	router := gin.Default()

	router.StaticFS("/static", http.Dir("../frontend/build/static"))
	router.StaticFile("/index", "../frontend/build/index.html")

	// ルートパスに対するハンドラを定義します
	router.GET("/", func(c *gin.Context) {
		// MySQLデータベースへのクエリを実行します
		var result string
		err := db.QueryRow("SELECT 'Hello, MySQL!'").Scan(&result)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to query MySQL"})
			return
		}

		c.JSON(200, gin.H{
			"message": result,
		})
	})

	// ルーターを実行します
	router.Run(":8080")
}
