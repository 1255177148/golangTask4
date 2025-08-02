package request

type CommentReq struct {
	PostID  uint   `json:"postId" form:"postId"`
	Content string `json:"content"`
	UserID  uint   `json:"userId"`
}
