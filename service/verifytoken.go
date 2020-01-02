package service

import "github.com/hello-micro/auth/pkg/jwt"

// VerifyToken 验证token
func VerifyToken(token string) bool {
	return jwt.VerifyToken(token)
}
