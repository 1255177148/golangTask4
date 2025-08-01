package controller

import (
	"github.com/1255177148/golangTask4/internal/app/model/dto"
	"github.com/1255177148/golangTask4/internal/app/service"
	"github.com/1255177148/golangTask4/internal/binder"
	"github.com/1255177148/golangTask4/internal/constant"
	"github.com/1255177148/golangTask4/internal/utils/log"
	"github.com/1255177148/golangTask4/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type PostController struct {
	postService *service.PostService
}

func NewPostController(postService *service.PostService) *PostController {
	return &PostController{postService: postService}
}

// CreatePost 创建文章接口
func (ctr *PostController) CreatePost(c *gin.Context) {
	userId, _ := c.Get("user_id")
	var postDTO dto.PostDTO
	if err := binder.BindAndValidate(c, &postDTO); err != nil {
		response.Fail(c, http.StatusBadRequest, constant.ParseParamFail)
		return
	}
	postDTO.ID = userId.(uint)
	if err := ctr.postService.CreatePost(&postDTO); err != nil {
		log.Error(constant.PostCreateFail, zap.Error(err))
		response.Fail(c, http.StatusInternalServerError, constant.PostCreateFail)
		return
	}
	response.Success(c, nil)
}

// GetPostList 获取文章列表
func (ctr *PostController) GetPostList(c *gin.Context) {
	var postDTOList []dto.PostDTO
	postDTOList, err := ctr.postService.FindPosts()
	if err != nil {
		log.Error("获取文章列表报错", zap.Error(err))
		response.Fail(c, http.StatusBadRequest, constant.ServiceFail)
		return
	}
	response.Success(c, postDTOList)
}

// Detail 获取文章详情
func (ctr *PostController) Detail(c *gin.Context) {
	var postDTO dto.PostDTO
	if err := binder.BindAndValidate(c, &postDTO); err != nil {
		response.Fail(c, http.StatusBadRequest, constant.ParseParamFail)
		return
	}
	post, err := ctr.postService.FindPostById(postDTO.ID)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, constant.ParseParamFail)
		return
	}
	response.Success(c, post)
}

// ModifyPost 修改文章
func (ctr *PostController) ModifyPost(c *gin.Context) {
	userId, _ := c.Get("user_id")
	var postDTO dto.PostDTO
	if err := binder.BindAndValidate(c, &postDTO); err != nil {
		response.Fail(c, http.StatusBadRequest, constant.ParseParamFail)
		return
	}
	postDTO.UserId = userId.(uint)
	if err := ctr.postService.UpdatePost(&postDTO); err != nil {
		log.Error("修改文章报错", zap.Error(err))
		response.Fail(c, http.StatusInternalServerError, constant.ServiceFail)
	}
	response.Success(c, nil)
}

// DeletePost 删除文章
func (ctr *PostController) DeletePost(c *gin.Context) {
	userId, _ := c.Get("user_id")
	var postDTO dto.PostDTO
	if err := binder.BindAndValidate(c, &postDTO); err != nil {
		response.Fail(c, http.StatusBadRequest, constant.ParseParamFail)
		return
	}
	if err := ctr.postService.DeletePost(postDTO.ID, userId.(uint)); err != nil {
		log.Error("删除文章报错", zap.Error(err))
		response.Fail(c, http.StatusInternalServerError, constant.ServiceFail)
		return
	}
	response.Success(c, nil)
}
