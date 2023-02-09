package utils

import (
	"beego_blog_mvc/global"
	"beego_blog_mvc/models"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/jaevor/go-nanoid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	mysqlAdmin, _ := beego.AppConfig.String("mysqlAdmin")
	mysqlPwd, _ := beego.AppConfig.String("mysqlPwd")
	mysqlDb, _ := beego.AppConfig.String("mysqlDb")
	mysqlUrl, _ := beego.AppConfig.String("mysqlUrl")
	mysqlPort, _ := beego.AppConfig.String("mysqlPort")
	dbUrl := mysqlAdmin + ":" + mysqlPwd + "@tcp(" + mysqlUrl + ":" + mysqlPort + ")/" + mysqlDb + "?charset=utf8&parseTime=True"

	var err error
	global.DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dbUrl,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	//判断mysql 链接是否异常
	if err != nil {
		panic("mysql数据库链接失败")
	}
	// 自动注册 模型
	AutoErr := global.DB.AutoMigrate(&models.Article{}, &models.User{}, &models.NavBar{}, models.Link{})
	if AutoErr != nil {
		fmt.Println("mysql迁移表失败")
	} else {
		canonical, _ := nanoid.Standard(36)
		// 迁移成功  向用户表插入一个用户，方便管理文章使用
		var users = models.NewUser()
		if users.GetUserListCount() == 0 {
			uid := canonical()
			global.DB.Create(&models.User{
				Uid:      uid,
				Username: "admin",
				Password: "123456",
			})
		}
		var article = models.NewArticle()
		if article.GetArticleListCount() == 0 {
			uid := canonical()
			article.Uid = uid
			article.Title = "新的一天"
			article.ContentHtml = "乾坤未定，你我皆是黑马"
			var _ = article.Create()
		}
	}
}
