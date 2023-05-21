package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rafikmoreira/go-blog-api/cmd/API/handler"
	"github.com/rafikmoreira/go-blog-api/cmd/API/middleware"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":  "Welcome to the Go Blog API",
			"author":   "Rafik Moreira",
			"github":   "https://github.com/rafikmoreira",
			"linkedin": "https://www.linkedin.com/in/rafikmoreira/",
		})
	})

	// Post routes
	r.POST("/post", middleware.AuthMiddleware(), handler.CreatePostHandler)
	r.GET("/post", handler.ListPostHandler)
	r.GET("/post/:postId", handler.GetPostHandler)
	r.PATCH("/post/:postId", handler.UpdatePostHandler)
	r.DELETE("/post/:postId", handler.DeletePostHandler)
	// Comment routes
	r.POST("/post/:postId/comment", handler.CreateCommentHandler)
	r.DELETE("/post/:postId/comment/:commentId", handler.DeleteCommentHandler)
	// User routes
	r.POST("/user", handler.CreateUserHandler)
	r.GET("/user", handler.ListUserHandler)
	r.GET("/user/:userId", handler.GetUserHandler)
	r.PATCH("/user/:userId", handler.UpdateUserHandler)
	r.DELETE("/user/:userId", handler.DeleteUserHandler)

	err := r.Run(":3333")

	if err != nil {
		panic(err)
	}
}
