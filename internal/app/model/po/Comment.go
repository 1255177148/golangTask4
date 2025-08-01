package po

import (
	"github.com/1255177148/golangTask4/internal/types"
)

type Comment struct {
	ID        uint           `json:"id"`
	PostID    uint           `json:"postId"`
	UserID    uint           `json:"userId"`
	Content   string         `json:"content"`
	CreatedAt types.JSONTime `json:"createdAt"`
}
