package models

import "github.com/google/uuid"

type Post struct {
    ID          string `json:"id"`
    Title       string `json:"title"`
    Author      string `json:"author"`
    Description string `json:"description"`
}

func NewPost(title, author, description string) *Post {
    return &Post{
        ID:          uuid.New().String(),
        Title:       title,
        Author:      author,
        Description: description,
    }
}
