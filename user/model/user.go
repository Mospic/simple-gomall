package model

import (
	"sync"
	"time"
)

type User struct {
	UserId   int32     `gorm:"primary_key;auto_increment"`
	Email    string    `gorm:"type:varchar(255);unique;not null"`
	Name     string    `gorm:"default:(-)"`
	Password string    `gorm:"default:(-)"`
	CreateAt time.Time `gorm:"not null"`
	DeleteAt time.Time `gorm:"default:NULL"`
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

func (*UserDao) CreateUser(user *User) (int32, error) {
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

func (d *UserDao) FindUserByID(id int32) (*User, error) {
	user := User{UserId: id}

	result := DB.Where("user_id = ?", id).First(&user)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (*UserDao) FindUserByName(username string) (*User, error) {
	user := User{Name: username}

	result := DB.Where("name = ?", username).First(&user)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (*UserDao) FindUserByEmail(email string) (*User, error) {
	user := User{Email: email}

	result := DB.Where("email = ?", email).First(&user)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &user, err
}
