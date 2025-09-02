package repository

import (
	"errors"
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

func (p *PostRepository) WithTx(tx *gorm.DB) *PostRepository {
	return &PostRepository{
		db:     tx,
		sqlxDB: p.sqlxDB,
	}
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
	if err := p.db.Table("posts").Select("id", "title", "created_at").Scan(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *PostRepository) GetPostById(id uint) (*dto.PostDTO, error) {
	var post dto.PostDTO
	if err := p.db.Table("posts").Where("id = ?", id).Scan(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (p *PostRepository) UpdatePost(postDTO *dto.PostDTO) error {
	return p.db.Transaction(func(tx *gorm.DB) error {
		// 查询文章作者是否是当前登录人员
		var userId uint
		if err := tx.Table("posts").Select("user_id").Where("id = ?", postDTO.ID).Scan(&userId).Error; err != nil {
			return err
		}
		if userId != postDTO.UserId {
			return errors.New("非文章作者，无权修改此文章")
		}
		var post po.Post
		if err := utils.MapStruct(postDTO, &post); err != nil {
			return err
		}
		return p.db.Model(&post).Updates(post).Error
	})
}

func (p *PostRepository) DeletePost(id uint, userId uint) error {
	return p.db.Transaction(func(tx *gorm.DB) error {
		var _userId uint
		if err := tx.Table("posts").Select("user_id").Where("id = ?", id).Scan(&_userId).Error; err != nil {
			return err
		}
		if _userId != userId {
			return errors.New("非文章作者，无权删除此文章")
		}
		return tx.Delete(&po.Post{}, "id = ?", id).Error
	})
}
