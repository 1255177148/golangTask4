package post

import (
	"github.com/1255177148/golangTask4/internal/app/repository"
	"github.com/1255177148/golangTask4/internal/app/service"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

func InitPostService(db *gorm.DB, sqlxDB *sqlx.DB, postRepo *repository.PostRepository, userRepo *repository.UserRepository) *service.PostService {
	return service.NewPostService(db, sqlxDB, postRepo, userRepo)
}
