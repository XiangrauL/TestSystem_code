package core

import (
	"fmt"
	"go_get_score/dao"
	"go_get_score/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitGorm() {
	// MySQL 连接字符串: "用户名:密码@tcp(服务器地址:端口)/数据库名?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := global.Host + ":" + global.Password + "@tcp(127.0.0.1:" + global.Port + ")/" + global.DataBase + "?charset=utf8mb4&parseTime=True&loc=Local"

	// 初始化 GORM 并连接 MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败: " + err.Error())
	}

	global.DB = *db

	// 自动迁移 (这将创建或更新表结构)
	err = db.AutoMigrate(&dao.Student{})
	if err != nil {
		fmt.Println(err.Error())
		panic("移表结构失败: " + err.Error())
	}
}
