package main

import (
	"latihan-assessment/backend/helper"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/users", helper.HandleUsers)
	r.POST("/users/", helper.HandleUsersPOST)
	r.POST("/users/register")
	r.POST("/users/login")
	r.GET("/users/:id", helper.HandleUserID)
	r.PUT("/users/:id")
	r.DELETE("/users/:id")

	r.GET("/books", helper.HandleBook)
	r.POST("/books/")
	r.GET("/books/:id")
	r.PUT("/books/:id")
	r.DELETE("/books/:id")

	r.Run(":8080")
}
