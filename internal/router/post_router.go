package router

import (
	"github.com/1255177148/golangTask4/internal/app/controller"
	"github.com/1255177148/golangTask4/internal/app/repository"
	"github.com/1255177148/golangTask4/internal/app/service"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

func RegisterPostRoutes(rg *gin.RouterGroup, db *gorm.DB, sqlxDB *sqlx.DB) {
	postRepository := repository.NewPostRepository(db, sqlxDB)
	userRepo := repository.NewUserRepository(db, sqlxDB)
	postService := service.NewPostService(db, sqlxDB, postRepository, userRepo)
	postController := controller.NewPostController(postService)

	posts := rg.Group("/posts")
	{
		posts.POST("/", postController.CreatePost)
		posts.GET("/list", postController.GetPostList)
		posts.GET("/detail", postController.Detail)
		posts.PUT("/modify", postController.ModifyPost)
	}
}
