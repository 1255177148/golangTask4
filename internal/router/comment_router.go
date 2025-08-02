package router

import (
	"github.com/1255177148/golangTask4/internal/app/controller"
	"github.com/1255177148/golangTask4/internal/app/repository"
	"github.com/1255177148/golangTask4/internal/app/service"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

func RegisterCommentRouters(rg *gin.RouterGroup, db *gorm.DB, sqlxDB *sqlx.DB) {
	userRepo := repository.NewUserRepository(db, sqlxDB)
	commentRepo := repository.NewCommentRepository(db, sqlxDB)
	commentService := service.NewCommentService(db, sqlxDB, commentRepo, userRepo)
	commentController := controller.NewCommentController(commentService)

	comments := rg.Group("/comments")
	{
		comments.POST("/", commentController.CreateComment)
		comments.GET("/", commentController.GetComments)
	}
}
