package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	log.Println("Server starting at :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
