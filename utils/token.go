package utils

import (
	"beego_blog_mvc/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CreateToken 设置token
func CreateToken(user models.User, secret string) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(time.Minute * 60 * 24).Unix(), // 1天
	})
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

// DeleteToken 删除token
func DeleteToken(user models.User, secret string) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(time.Second * 1).Unix(), // 1秒
	})
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

// ParseToken 解析token
func ParseToken(token string, secret string) (jwt.MapClaims, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return claim.Claims.(jwt.MapClaims), nil
}
