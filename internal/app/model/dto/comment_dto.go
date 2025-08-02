package dto

import (
	"github.com/1255177148/golangTask4/internal/types"
)

type CommentDTO struct {
	ID        uint           `json:"id"`
	Content   string         `json:"content"`
	CreatedAt types.JSONTime `json:"createdAt"`
}
