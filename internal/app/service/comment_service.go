package service

import (
	"errors"
	"github.com/1255177148/golangTask4/internal/app/model/dto"
	"github.com/1255177148/golangTask4/internal/app/model/po"
	"github.com/1255177148/golangTask4/internal/app/model/request"
	"github.com/1255177148/golangTask4/internal/app/repository"
	"github.com/1255177148/golangTask4/internal/constant"
	"github.com/1255177148/golangTask4/internal/utils"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type CommentService struct {
	db          *gorm.DB
	sqlxDB      *sqlx.DB
	commentRepo *repository.CommentRepository
	userRepo    *repository.UserRepository
}

func NewCommentService(db *gorm.DB, sqlxDB *sqlx.DB, commentRepo *repository.CommentRepository, userRepo *repository.UserRepository) *CommentService {
	return &CommentService{db: db, sqlxDB: sqlxDB, commentRepo: commentRepo}
}

// CreateComment 创建评论
func (cs *CommentService) CreateComment(commentReq *request.CommentReq) error {
	return cs.db.Transaction(func(tx *gorm.DB) error {
		// 校验是否认证过
		userRepo := repository.NewUserRepository(tx, cs.sqlxDB)
		var authFlag string
		authFlag, err := userRepo.CheckUserAuth(commentReq.UserID)
		if err != nil {
			return err
		}
		if authFlag == "0" {
			return errors.New(constant.NotAuth)
		}
		commentRepo := repository.NewCommentRepository(tx, cs.sqlxDB)
		var comment po.Comment
		err = utils.MapStruct(commentReq, &comment)
		if err != nil {
			return err
		}
		return commentRepo.CreateComment(&comment)
	})
}

// GetComments 获取指定文章的评论列表
func (cs *CommentService) GetComments(postId uint) ([]dto.CommentDTO, error) {
	return cs.commentRepo.GetCommentsByPost(postId)
}
