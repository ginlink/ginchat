package main

import (
	"ginchat/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	router := gin.Default()
	routes.UserRoute(router)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
