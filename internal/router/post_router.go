package router

import (
	"github.com/1255177148/golangTask4/internal/app/controller"
	"github.com/1255177148/golangTask4/internal/container"
	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(rg *gin.RouterGroup) {
	postService := container.Instance.PostService
	postController := controller.NewPostController(postService)

	posts := rg.Group("/posts")
	{
		posts.POST("/", postController.CreatePost)
		posts.GET("/list", postController.GetPostList)
		posts.GET("/detail", postController.Detail)
		posts.PUT("/modify", postController.ModifyPost)
		posts.DELETE("/:id", postController.DeletePost)
	}
}
