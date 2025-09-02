package repository

import (
	"github.com/1255177148/golangTask4/internal/app/model/po"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	db     *gorm.DB
	sqlxDB *sqlx.DB
}

func NewUserRepository(db *gorm.DB, sqlxDB *sqlx.DB) *UserRepository {
	return &UserRepository{db: db, sqlxDB: sqlxDB}
}

func (u *UserRepository) WithTx(tx *gorm.DB) *UserRepository {
	return &UserRepository{
		db:     tx,
		sqlxDB: u.sqlxDB,
	}
}

// CreateUser 注册用户
func (u *UserRepository) CreateUser(user *po.User) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashPassword)
	// 插入时忽略authentication_flag，让该字段用数据库的默认值
	return u.db.Omit("authentication_flag").Create(user).Error
}

// FindUserByUsername 根据用户名获取用户信息
func (u *UserRepository) FindUserByUsername(username string) (*po.User, error) {
	var user po.User
	if err := u.db.Where("user_name = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) FindUserByID(id uint) (*po.User, error) {
	var user po.User
	if err := u.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) UpdateUser(user *po.User) error {
	if err := u.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

// CheckUserAuth 查询用户是否认证
func (u *UserRepository) CheckUserAuth(id uint) (string, error) {
	var auth string
	if err := u.db.Table("users").Select("authentication_flag").Where("id = ?", id).Scan(&auth).Error; err != nil {
		return "", err
	}
	return auth, nil
}
