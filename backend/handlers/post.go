package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
	"github.com/google/uuid"
    "github.com/h10k/blog-post/models" 
)

var posts = make(map[string]*models.Post)

// CreatePost creates a new blog post
func CreatePost(c *gin.Context) {
    var post models.Post
    if err := c.ShouldBindJSON(&post); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    post.ID = uuid.New().String()

    posts[post.ID] = &post
	c.JSON(http.StatusCreated, gin.H{
        "message": "Article créé avec succès",
        "post":    post,
    })
}

// GetPosts retrieves all blog posts
func GetPosts(c *gin.Context) {
    values := make([]*models.Post, 0, len(posts))
    for _, v := range posts {
        values = append(values, v)
    }
    c.JSON(http.StatusOK, values)
}

// GetPost retrieves a specific blog post by ID
func GetPost(c *gin.Context) {
    id := c.Param("id")
    post, ok := posts[id]
    if !ok {
        c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        return
    }
    c.JSON(http.StatusOK, post)
}

// UpdatePost updates a blog post by ID
func UpdatePost(c *gin.Context) {
    id := c.Param("id")
    post, ok := posts[id]
    if !ok {
        c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        return
    }

    var updatedPost models.Post
    if err := c.ShouldBindJSON(&updatedPost); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    post.Title = updatedPost.Title
    post.Author = updatedPost.Author
    post.Description = updatedPost.Description

	c.JSON(http.StatusOK, gin.H{
        "message": "Article mis à jour avec succès",
        "post":    post,
    })
}

// DeletePost deletes a blog post by ID
func DeletePost(c *gin.Context) {
    id := c.Param("id")
    _, ok := posts[id]
    if !ok {
        c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        return
    }

    delete(posts, id)
	c.JSON(http.StatusOK, gin.H{
        "message": "Article supprimé avec succès",
    })
}
