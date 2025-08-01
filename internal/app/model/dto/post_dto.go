package dto

import "github.com/1255177148/golangTask4/internal/types"

type PostDTO struct {
	ID        uint           `json:"id" form:"id" uri:"id"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	CreatedAt types.JSONTime `json:"createdAt"`
	UserId    uint           `json:"userId"`
}
