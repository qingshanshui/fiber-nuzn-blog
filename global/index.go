package global

import (
	"gorm.io/gorm"
)

var (
	DB *gorm.DB // mysql连接
	//Config *models.YamlModel // 配置文件
	//Redis  *redis.Client     // redis连接
	//Log    *zap.Logger       // 日志
	//Req    *response.Req     // 响应处理
)
