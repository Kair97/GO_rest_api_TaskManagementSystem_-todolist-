package main

import (
	"log"
	"try/database"
	"try/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.Static("/static", "./public")
	r.LoadHTMLFiles("public/index.html")
	routes.RegisterRoutes(r)

	r.Run(":8080")

}
