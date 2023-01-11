package models

import (
	"time"

)

//create post model struct

type Post struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"size:200" json:"title"`
	Body      string    `gorm:"size:200" json:"body"`
	Description string `gorm:"size:200" json:"description"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

//tableNmae methods set table name for post model

func (post *Post) tableNamelist() string {
	return "post"
}

//response map -> response method for post

func (post *Post) ResponseMap() map[string]interface{} {
	response := make(map[string]interface{})
	response["id"] = post.ID
	response["title"] = post.Title
	response["body"] = post.Body
	response["description"] = post.Description
	response["created_at"] = post.CreatedAt
	response["updated_at"] = post.UpdatedAt
    return response
}
