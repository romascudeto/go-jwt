package main

import (
	"fmt"
	"go-jwt/article"
	"go-jwt/author"
	Config "go-jwt/config"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var err error

func main() {
	godotenv.Load()
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	// defer Config.DB.Close()

	r := gin.Default()
	article.Routes(r)
	author.Routes(r)

	r.Run(":8088") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
