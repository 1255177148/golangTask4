package comment

import (
	"github.com/1255177148/golangTask4/internal/app/repository"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

func InitCommentRepository(db *gorm.DB, sqlxDB *sqlx.DB) *repository.CommentRepository {
	return repository.NewCommentRepository(db, sqlxDB)
}
