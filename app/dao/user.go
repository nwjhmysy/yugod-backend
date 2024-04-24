package dao

import (
	"errors"
	"yugod-backend/app/model"
	"yugod-backend/app/openapi"
)

type iUserDao interface {
	CreateUser(param *openapi.CreateUserParam) error
	LoginByUserName(loginVals model.LoginParam) (*model.LoginClaims, error)
	GetUserInfo(userId uint) (*openapi.GetUserInfoData, error)
}
type userDao struct{}

var UserDao iUserDao = userDao{}

// 创建新用户
func (userDao) CreateUser(param *openapi.CreateUserParam) error {
	user := &model.User{
		UserName: param.UserName,
		Password: param.Password,
		Email:    param.Email,
		Name:     param.Name,
		Auth:     uint(param.Auth),
		Age:      uint(param.Age),
	}
	err := DB.Model(user).Create(user).Error

	return err
}

// 登录
func (userDao) LoginByUserName(loginVals model.LoginParam) (*model.LoginClaims, error) {
	user := &model.User{}
	// 数据库获取
	claims := &model.LoginClaims{}

	err := DB.Model(user).First(claims, "user_name = ?", loginVals.UserName).Error

	if claims.Id > 0 && claims.Password == loginVals.Password {
		return claims, nil
	}
	if claims.Id <= 0 {
		err = errors.New("")
	}
	if claims.Password != loginVals.Password {
		err = errors.New("")
	}

	return nil, err
}

func (userDao) GetUserInfo(userId uint) (*openapi.GetUserInfoData, error) {
	user := &model.User{}
	err := DB.Model(user).First(user, "id = ?", userId).Error

	if err != nil {
		return nil, err
	}

	userInfo := &openapi.GetUserInfoData{
		UserId:   int32(userId),
		UserName: user.UserName,
		Auth:     int32(user.Auth),
		Name:     user.Name,
		Email:    user.Email,
		Age:      int32(user.Age),
	}
	return userInfo, nil
}
