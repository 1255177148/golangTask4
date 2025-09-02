package router

import (
	"github.com/1255177148/golangTask4/internal/app/controller"
	"github.com/1255177148/golangTask4/internal/container"
	"github.com/gin-gonic/gin"
)

func RegisterCommentRouters(rg *gin.RouterGroup) {
	commentService := container.Instance.CommentService
	commentController := controller.NewCommentController(commentService)

	comments := rg.Group("/comments")
	{
		comments.POST("/", commentController.CreateComment)
		comments.GET("/", commentController.GetComments)
	}
}
