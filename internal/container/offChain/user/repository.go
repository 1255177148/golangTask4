package user

import (
	"github.com/1255177148/golangTask4/internal/app/repository"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

func InitUserRepository(db *gorm.DB, sqlxDB *sqlx.DB) *repository.UserRepository {
	return repository.NewUserRepository(db, sqlxDB)
}
