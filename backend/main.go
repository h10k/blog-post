package main

import (
    "github.com/gin-gonic/gin"
    "github.com/h10k/blog-post/handlers"
)

func main() {
    r := gin.Default()

    r.POST("/posts", handlers.CreatePost)
    r.GET("/posts", handlers.GetPosts)
    r.GET("/posts/:id", handlers.GetPost)
    r.PUT("/posts/:id", handlers.UpdatePost)
    r.DELETE("/posts/:id", handlers.DeletePost)

    r.Run() 
}
