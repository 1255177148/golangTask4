package service

import (
	"errors"
	"github.com/1255177148/golangTask4/internal/app/model/dto"
	"github.com/1255177148/golangTask4/internal/app/repository"
	"github.com/1255177148/golangTask4/internal/constant"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type PostService struct {
	db       *gorm.DB
	sqlxDB   *sqlx.DB
	postRepo *repository.PostRepository
	userRepo *repository.UserRepository
}

func NewPostService(db *gorm.DB, sqlxDB *sqlx.DB, postRepo *repository.PostRepository, userRepo *repository.UserRepository) *PostService {
	return &PostService{db: db, sqlxDB: sqlxDB, postRepo: postRepo, userRepo: userRepo}
}

// CreatePost 创建文章
func (ps *PostService) CreatePost(postDTO *dto.PostDTO) error {
	return ps.db.Transaction(func(tx *gorm.DB) error {
		userRepo := repository.NewUserRepository(tx, ps.sqlxDB)
		// 先查看用户有没有认证
		var authFlag string
		authFlag, err := userRepo.CheckUserAuth(postDTO.ID)
		if err != nil {
			return err
		}
		if authFlag == "0" {
			return errors.New(constant.NotAuth)
		}
		postRepo := repository.NewPostRepository(tx, ps.sqlxDB)
		// 将文章写入数据库
		if err := postRepo.CreatePost(postDTO); err != nil {
			return err
		}
		return nil
	})
}

// FindPosts 获取文章列表
func (ps *PostService) FindPosts() ([]dto.PostDTO, error) {
	return ps.postRepo.FindPosts()
}
