package dto

type CommentDTO struct {
	ID        uint   `json:"id"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
}
