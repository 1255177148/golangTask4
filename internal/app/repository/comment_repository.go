package repository

import (
	"github.com/1255177148/golangTask4/internal/app/model/dto"
	"github.com/1255177148/golangTask4/internal/app/model/po"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db     *gorm.DB
	sqlxDB *sqlx.DB
}

func NewCommentRepository(db *gorm.DB, sqlxDB *sqlx.DB) *CommentRepository {
	return &CommentRepository{db: db, sqlxDB: sqlxDB}
}

func (cr *CommentRepository) WithTx(tx *gorm.DB) *CommentRepository {
	return &CommentRepository{
		db:     tx,
		sqlxDB: cr.sqlxDB, // 保留 sqlxDB
	}
}

func (cr *CommentRepository) CreateComment(comment *po.Comment) error {
	return cr.db.Create(comment).Error
}

func (cr *CommentRepository) GetCommentsByPost(postId uint) ([]dto.CommentDTO, error) {
	var comments []dto.CommentDTO
	if err := cr.db.Table("comments").Where("post_id = ?", postId).Scan(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
