package post

import (
	"github.com/1255177148/golangTask4/internal/app/repository"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

func InitPostRepository(db *gorm.DB, sqlxDB *sqlx.DB) *repository.PostRepository {
	return repository.NewPostRepository(db, sqlxDB)
}
