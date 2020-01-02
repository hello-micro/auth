package service

import (
	"errors"

	"github.com/hello-micro/auth/pkg/database"
	"github.com/hello-micro/auth/pkg/jwt"
	"github.com/hello-micro/auth/pkg/models"
	"github.com/hello-micro/auth/pkg/utils"
)

// PasswordLoginService PasswordLoginService
type PasswordLoginService struct {
	Username   string
	Password   string
	AutoCreate bool
}

// Login Login
func (m *PasswordLoginService) Login() (LoginRequest, error) {
	if m.Password == "" || m.Username == "" {
		return LoginRequest{}, errors.New("code is empty")
	}
	// 用户登录
	userAuth, err := loginOrCreate(m.Username, utils.MD5Sum(m.Password), m.AutoCreate)
	if err != nil {
		return LoginRequest{}, err
	}
	// 判断是否有基础用户信息
	var userBase models.UserBase
	err = database.Conn.Where("uid=?", userAuth.UID).Find(&userBase).Error
	if err != nil {
		return LoginRequest{}, err
	}
	tokenStr, err := jwt.GenerateToken(&userBase)
	if err != nil {
		return LoginRequest{}, err
	}
	return LoginRequest{Token: tokenStr}, err
}
