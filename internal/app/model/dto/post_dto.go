package dto

type PostDTO struct {
	ID        uint   `json:"id" form:"id" uri:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
	UserId    uint   `json:"userId"`
}
