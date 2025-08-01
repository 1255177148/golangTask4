package po

import (
	"github.com/1255177148/golangTask4/internal/types"
)

type Post struct {
	ID        uint           `json:"id"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	UserId    uint           `json:"userId"`
	CreatedAt types.JSONTime `json:"createdAt"`
	UpdatedAt types.JSONTime `json:"updatedAt"`
}
