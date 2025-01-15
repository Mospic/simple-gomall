package model

import (
	"github.com/jinzhu/gorm"
	"sync"
)

type User struct {
	gorm.Model
	UserId   int64  `gorm:"primary_key;auto_increment"`
	Name     string `gorm:"default:(-)"`
	Password string `gorm:"default:(-)"`
}

func (User) TableName() string {
	return "user"
}

type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once

func NewUserDao() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

/**
根据用户名和密码，创建一个新的User，返回UserId
*/

func (*UserDao) CreateUser(user *User) (int64, error) {
	/*user := User{Name: username, Password: password, FollowingCount: 0, FollowerCount: 0, CreateAt: time.Now()}*/

	result := DB.Create(&user)

	if result.Error != nil {
		return -1, result.Error
	}

	return user.UserId, nil
}

/*
根据用户ID 查找用户实体
*/

func (d *UserDao) FindUserByID(id int64) (*User, error) {
	user := User{UserId: id}

	result := DB.Where("user_id = ?", id).First(&user)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &user, err
}
