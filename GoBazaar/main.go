package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"GoBazaar/handler"
)

func main() {
	router := gin.Default()
	router.GET("/welcome", handler.HomeHandler)
	router.POST("/user/register", handler.RegisterUser)
	err := router.Run("127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
}
