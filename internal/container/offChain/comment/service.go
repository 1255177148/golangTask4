package comment

import (
	"github.com/1255177148/golangTask4/internal/app/repository"
	"github.com/1255177148/golangTask4/internal/app/service"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

func InitCommentService(db *gorm.DB, sqlxDB *sqlx.DB, commentRepo *repository.CommentRepository, userRepo *repository.UserRepository) *service.CommentService {
	return service.NewCommentService(db, sqlxDB, commentRepo, userRepo)
}
