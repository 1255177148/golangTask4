package service

import (
	"errors"
	"github.com/1255177148/golangTask4/internal/app/model/dto"
	"github.com/1255177148/golangTask4/internal/app/model/po"
	"github.com/1255177148/golangTask4/internal/app/repository"
	"github.com/1255177148/golangTask4/internal/utils"
	"github.com/1255177148/golangTask4/internal/utils/log"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db     *gorm.DB
	sqlxDB *sqlx.DB
	repo   *repository.UserRepository
}

func NewUserService(db *gorm.DB, sqlxDB *sqlx.DB, repo *repository.UserRepository) *UserService {
	return &UserService{db: db, sqlxDB: sqlxDB, repo: repo}
}

// RegisterUser 注册用户
func (us *UserService) RegisterUser(user *po.User) error {
	return us.repo.CreateUser(user)
}

// CheckLogin 用户登录
func (us *UserService) CheckLogin(user *dto.UserDTO) (string, error) {
	// 根据用户名从数据库里获取数据
	originUser, err := us.repo.FindUserByUsername(user.Username)
	if err != nil {
		log.Error("用户登录报错", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("未查询到该用户")
		} else {
			return "", errors.New("查询报错")
		}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(originUser.Password), []byte(user.Password)); err != nil {
		log.Error("密码校验报错", zap.Error(err))
		return "", errors.New("用户名或密码不正确")
	}
	token, err := utils.GenerateAccessToken(originUser.ID)
	if err != nil {
		log.Error("生成token出错", zap.Error(err))
		return "", errors.New("生成token出错")
	}
	return token, nil
}

// AuthUser 用户认证
func (us *UserService) AuthUser(userAuth *dto.UserAuth) error {
	return us.db.Transaction(func(tx *gorm.DB) error {
		repo := repository.NewUserRepository(tx, us.sqlxDB)
		// 先获取用户数据
		user, err := repo.FindUserByID(userAuth.ID)
		if err != nil {
			return err
		}
		user.Email = &userAuth.Email
		var AuthenticationFlag = "1"
		user.AuthenticationFlag = &AuthenticationFlag
		if err = us.repo.UpdateUser(user); err != nil {
			return err
		}
		return nil
	})
}
