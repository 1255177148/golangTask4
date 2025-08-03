package controller

import (
	"github.com/1255177148/golangTask4/internal/app/model/request"
	"github.com/1255177148/golangTask4/internal/app/service"
	"github.com/1255177148/golangTask4/internal/binder"
	"github.com/1255177148/golangTask4/internal/constant"
	"github.com/1255177148/golangTask4/internal/utils/log"
	"github.com/1255177148/golangTask4/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type CommentController struct {
	commentService *service.CommentService
}

func NewCommentController(commentService *service.CommentService) *CommentController {
	return &CommentController{commentService: commentService}
}

// CreateComment 创建评论
// @Summary 创建评论
// @Description 创建评论
// @Tags 评论
// @Produce json
// @Param CommentReq body request.CommentReq true "新增的评论数据，userId不用传"
// @Success 200 {object} response.ResultResponse
// @Security ApiKeyAuth
// @Router /v1/comments [post]
func (cc *CommentController) CreateComment(c *gin.Context) {
	userId, _ := c.Get("user_id")
	var commentReq request.CommentReq
	if err := binder.BindAndValidate(c, &commentReq); err != nil {
		response.Fail(c, http.StatusBadRequest, constant.ParseParamFail)
		return
	}
	commentReq.UserID = userId.(uint)
	if err := cc.commentService.CreateComment(&commentReq); err != nil {
		log.Error("文章评论创建失败", zap.Error(err))
		response.Fail(c, http.StatusInternalServerError, constant.ServiceFail)
		return
	}
	response.Success(c, nil)
}

// GetComments 获取文章评论列表
// @Summary 获取文章评论列表
// @Description 获取文章评论列表
// @Tags 评论
// @Produce json
// @Param postId query uint true "文章id"
// @Success 200 {object} response.ResultResponse
// @Security ApiKeyAuth
// @Router /v1/comments [get]
func (cc *CommentController) GetComments(c *gin.Context) {
	var commentReq request.CommentReq
	if err := binder.BindAndValidate(c, &commentReq); err != nil {
		response.Fail(c, http.StatusBadRequest, constant.ParseParamFail)
		return
	}
	comments, err := cc.commentService.GetComments(commentReq.PostID)
	if err != nil {
		log.Error("获取文章评论列表失败", zap.Error(err))
		response.Fail(c, http.StatusInternalServerError, constant.ServiceFail)
		return
	}
	response.Success(c, comments)
}
