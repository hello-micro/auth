// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/auth/auth.proto

package hello_micro_srv_auth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Auth service

type AuthService interface {
	// 小程序登录
	MiniProgramLogin(ctx context.Context, in *MiniProgramLoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	// 手机验证码登录
	PhoneCodeLogin(ctx context.Context, in *PhoneCodeLoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	// 邮箱验证码登录
	EmailCodeLogin(ctx context.Context, in *EmailCodeLoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	// 密码登录 (手机密码.用户名密码.邮箱密码)
	PasswordLogin(ctx context.Context, in *PasswordLoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	// 验证token是否有效
	VerifyToken(ctx context.Context, in *VerifyTokenRequest, opts ...client.CallOption) (*VerifyTokenResponse, error)
	// 刷新token
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...client.CallOption) (*RefreshTokenResponse, error)
}

type authService struct {
	c    client.Client
	name string
}

func NewAuthService(name string, c client.Client) AuthService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "hello.micro.srv.auth"
	}
	return &authService{
		c:    c,
		name: name,
	}
}

func (c *authService) MiniProgramLogin(ctx context.Context, in *MiniProgramLoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.MiniProgramLogin", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) PhoneCodeLogin(ctx context.Context, in *PhoneCodeLoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.PhoneCodeLogin", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) EmailCodeLogin(ctx context.Context, in *EmailCodeLoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.EmailCodeLogin", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) PasswordLogin(ctx context.Context, in *PasswordLoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.PasswordLogin", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) VerifyToken(ctx context.Context, in *VerifyTokenRequest, opts ...client.CallOption) (*VerifyTokenResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.VerifyToken", in)
	out := new(VerifyTokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...client.CallOption) (*RefreshTokenResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.RefreshToken", in)
	out := new(RefreshTokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthHandler interface {
	// 小程序登录
	MiniProgramLogin(context.Context, *MiniProgramLoginRequest, *LoginResponse) error
	// 手机验证码登录
	PhoneCodeLogin(context.Context, *PhoneCodeLoginRequest, *LoginResponse) error
	// 邮箱验证码登录
	EmailCodeLogin(context.Context, *EmailCodeLoginRequest, *LoginResponse) error
	// 密码登录 (手机密码.用户名密码.邮箱密码)
	PasswordLogin(context.Context, *PasswordLoginRequest, *LoginResponse) error
	// 验证token是否有效
	VerifyToken(context.Context, *VerifyTokenRequest, *VerifyTokenResponse) error
	// 刷新token
	RefreshToken(context.Context, *RefreshTokenRequest, *RefreshTokenResponse) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) error {
	type auth interface {
		MiniProgramLogin(ctx context.Context, in *MiniProgramLoginRequest, out *LoginResponse) error
		PhoneCodeLogin(ctx context.Context, in *PhoneCodeLoginRequest, out *LoginResponse) error
		EmailCodeLogin(ctx context.Context, in *EmailCodeLoginRequest, out *LoginResponse) error
		PasswordLogin(ctx context.Context, in *PasswordLoginRequest, out *LoginResponse) error
		VerifyToken(ctx context.Context, in *VerifyTokenRequest, out *VerifyTokenResponse) error
		RefreshToken(ctx context.Context, in *RefreshTokenRequest, out *RefreshTokenResponse) error
	}
	type Auth struct {
		auth
	}
	h := &authHandler{hdlr}
	return s.Handle(s.NewHandler(&Auth{h}, opts...))
}

type authHandler struct {
	AuthHandler
}

func (h *authHandler) MiniProgramLogin(ctx context.Context, in *MiniProgramLoginRequest, out *LoginResponse) error {
	return h.AuthHandler.MiniProgramLogin(ctx, in, out)
}

func (h *authHandler) PhoneCodeLogin(ctx context.Context, in *PhoneCodeLoginRequest, out *LoginResponse) error {
	return h.AuthHandler.PhoneCodeLogin(ctx, in, out)
}

func (h *authHandler) EmailCodeLogin(ctx context.Context, in *EmailCodeLoginRequest, out *LoginResponse) error {
	return h.AuthHandler.EmailCodeLogin(ctx, in, out)
}

func (h *authHandler) PasswordLogin(ctx context.Context, in *PasswordLoginRequest, out *LoginResponse) error {
	return h.AuthHandler.PasswordLogin(ctx, in, out)
}

func (h *authHandler) VerifyToken(ctx context.Context, in *VerifyTokenRequest, out *VerifyTokenResponse) error {
	return h.AuthHandler.VerifyToken(ctx, in, out)
}

func (h *authHandler) RefreshToken(ctx context.Context, in *RefreshTokenRequest, out *RefreshTokenResponse) error {
	return h.AuthHandler.RefreshToken(ctx, in, out)
}
