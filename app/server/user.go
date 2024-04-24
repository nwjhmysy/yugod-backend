package server

import (
	"yugod-backend/app/dao"
	"yugod-backend/app/openapi"
)

func CreateUser(param *openapi.CreateUserParam) error {
	err := dao.UserDao.CreateUser(param)
	return err
}

func GetUserInfo(userId uint) (*openapi.GetUserInfoData, error) {
	data, err := dao.UserDao.GetUserInfo(userId)
	return data, err
}
