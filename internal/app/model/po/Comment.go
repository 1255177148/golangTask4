package po

import "time"

type Comment struct {
	ID        uint      `json:"id"`
	PostID    uint      `json:"postId"`
	UserID    uint      `json:"userId"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}
