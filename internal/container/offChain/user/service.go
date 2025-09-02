package user

import (
	"github.com/1255177148/golangTask4/internal/app/repository"
	"github.com/1255177148/golangTask4/internal/app/service"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

func InitUserService(db *gorm.DB, sqlxDB *sqlx.DB, repo *repository.UserRepository) *service.UserService {
	return service.NewUserService(db, sqlxDB, repo)
}
