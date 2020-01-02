package handler

import (
	"context"
	"errors"

	auth "github.com/hello-micro/auth/proto/auth"
	"github.com/hello-micro/auth/service"
)

type Auth struct{}

// VerifyToken 验证token
func (e *Auth) VerifyToken(ctx context.Context, req *auth.VerifyTokenRequest, rsp *auth.VerifyTokenResponse) error {
	rsp.Valid = service.VerifyToken(req.Token)
	return nil
}

// RefreshToken 刷新token
func (e *Auth) RefreshToken(ctx context.Context, req *auth.RefreshTokenRequest, rsp *auth.RefreshTokenResponse) error {
	token, err := service.RefreshToken(req.Token)
	if err != nil {
		return err
	}
	rsp.Token = token
	return nil
}

// MiniProgramLogin 小程序登录
func (e *Auth) MiniProgramLogin(ctx context.Context, req *auth.MiniProgramLoginRequest, rsp *auth.LoginResponse) error {
	miniProgramLoginService := service.MiniProgramLoginService{
		Channel:    req.Channel,
		Code:       req.Code,
		AutoCreate: true,
	}
	request, err := miniProgramLoginService.Login()
	if err != nil {
		return err
	}
	rsp.Token = request.Token
	return nil
}

// PhoneCodeLogin 手机登录
func (e *Auth) PhoneCodeLogin(context.Context, *auth.PhoneCodeLoginRequest, *auth.LoginResponse) error {
	return errors.New("暂不支持")
}

// EmailCodeLogin 邮箱登录
func (e *Auth) EmailCodeLogin(context.Context, *auth.EmailCodeLoginRequest, *auth.LoginResponse) error {
	return errors.New("暂不支持")
}

// PasswordLogin 密码登录
func (e *Auth) PasswordLogin(ctx context.Context, req *auth.PasswordLoginRequest, rsp *auth.LoginResponse) error {
	passwordLoginService := service.PasswordLoginService{
		Username: req.LoginName,
		Password: req.Password,
	}
	request, err := passwordLoginService.Login()
	if err != nil {
		return err
	}
	rsp.Token = request.Token
	return nil
}
