package admin

import (
	"errors"
	"fiber-nuzn-blog/models"
	"fiber-nuzn-blog/pkg/utils"
	validatorForm "fiber-nuzn-blog/validator/form/admin"
	"github.com/spf13/viper"
)

type Login struct{}

func NewLoginService() *Login {
	return &Login{}
}

// Login 登录
func (t *Login) Login(r validatorForm.LoginRequest) (string, error) {
	// 业务处理
	if r.Username == "" || r.Password == "" {
		return "", errors.New("用户名密码不能为空")
	}
	mu := models.NewUser()
	if err := mu.Login(r.Username, r.Password); err != nil {
		return "", err
	}
	if mu.Uid != "" {
		token, _ := utils.CreateToken(mu, viper.GetString("Jwt.Secret"))
		return token, nil
	}
	return "", errors.New("登录失败")
}
