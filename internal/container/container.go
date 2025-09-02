package container

import (
	"github.com/1255177148/golangTask4/internal/app/repository"
	"github.com/1255177148/golangTask4/internal/app/service"
	"github.com/1255177148/golangTask4/internal/app/service/onchain/erc20demo"
	"github.com/1255177148/golangTask4/internal/container/offChain/comment"
	"github.com/1255177148/golangTask4/internal/container/offChain/post"
	"github.com/1255177148/golangTask4/internal/container/offChain/user"
	containererc20 "github.com/1255177148/golangTask4/internal/container/onchain/erc20demo"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

// Container 统一管理repository、service和listener
type Container struct {
	UserRepository    *repository.UserRepository
	PostRepository    *repository.PostRepository
	CommentRepository *repository.CommentRepository
	UserService       *service.UserService
	PostService       *service.PostService
	CommentService    *service.CommentService
	Erc20Service      *erc20demo.Erc20Service
	ERC20Listener     *erc20demo.Listener
}

var Instance Container

// Init 初始化并管理所有的repository、service和contract实例
func Init(db *gorm.DB, sqlxDB *sqlx.DB) *Container {
	userRepository := user.InitUserRepository(db, sqlxDB)
	postRepository := post.InitPostRepository(db, sqlxDB)
	commentRepository := comment.InitCommentRepository(db, sqlxDB)
	userService := user.InitUserService(db, sqlxDB, userRepository)
	postService := post.InitPostService(db, sqlxDB, postRepository, userRepository)
	commentService := comment.InitCommentService(db, sqlxDB, commentRepository, userRepository)
	erc20Service := containererc20.InitERC20DemoService()
	erc20Listener := containererc20.InitERC20Listener(erc20Service)
	Instance = Container{
		UserRepository:    userRepository,
		PostRepository:    postRepository,
		CommentRepository: commentRepository,
		UserService:       userService,
		PostService:       postService,
		CommentService:    commentService,
		Erc20Service:      erc20Service,
		ERC20Listener:     erc20Listener,
	}
	return &Instance
}
