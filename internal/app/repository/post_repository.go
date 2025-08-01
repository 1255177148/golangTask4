package repository

import (
	"github.com/1255177148/golangTask4/internal/app/model/dto"
	"github.com/1255177148/golangTask4/internal/app/model/po"
	"github.com/1255177148/golangTask4/internal/utils"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type PostRepository struct {
	db     *gorm.DB
	sqlxDB *sqlx.DB
}

func NewPostRepository(db *gorm.DB, sqlxDB *sqlx.DB) *PostRepository {
	return &PostRepository{db: db, sqlxDB: sqlxDB}
}

func (p *PostRepository) CreatePost(postDTO *dto.PostDTO) error {
	var post po.Post
	if err := utils.MapStruct(postDTO, &post); err != nil {
		return err
	}
	return p.db.Create(&post).Error
}

// FindPosts 获取文章列表
func (p *PostRepository) FindPosts() ([]dto.PostDTO, error) {
	var posts []dto.PostDTO
	if err := p.db.Table("posts").Scan(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
