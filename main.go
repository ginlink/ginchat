package main

import (
	"ginchat/docs"
	"ginchat/routes"
	_ "ginchat/utils"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := gin.Default()

	docs.SwaggerInfo.BasePath = ""
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	routes.UserRoute(router)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
