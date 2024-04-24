package server

import (
	"yugod-backend/app/dao"
	"yugod-backend/app/openapi"
)

func CreateUser(param *openapi.CreateUserParam) error {
	err := dao.UserDao.CreateUser(param)
	return err
}
