package service

import "github.com/hello-micro/auth/pkg/jwt"

// RefreshToken 刷新token
func RefreshToken(token string) (string, error) {
	return jwt.RefreshToken(token)
}
