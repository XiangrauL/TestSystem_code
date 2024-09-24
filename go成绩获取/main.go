package main

import (
	"github.com/gin-gonic/gin"
	"go_get_score/core"
	"go_get_score/httpserver"
)

//type Student struct {
//	ID       int    `gorm:"primaryKey"`
//	Name     string `gorm:"unique;not null"` // 姓名字段，唯一且不能为空
//	Password string `gorm:"not null"`        // 密码字段，不能为空
//	Exam1    int
//	Exam2    int
//	Exam3    int
//	Exam4    int
//	Exam5    int
//	Exam6    int
//	Exam7    int
//	Exam8    int
//	Exam9    int
//}
//
//// 用于数据库连接
//var db *gorm.DB
//var err error
//
//func init() {
//	// MySQL 连接字符串: "用户名:密码@tcp(服务器地址:端口)/数据库名?charset=utf8mb4&parseTime=True&loc=Local"
//	dsn := "root:Wcs200369@tcp(127.0.0.1:3306)/student_db?charset=utf8mb4&parseTime=True&loc=Local"
//
//	// 初始化 GORM 并连接 MySQL
//	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		fmt.Println(err.Error())
//		panic("连接数据库失败: " + err.Error())
//	}
//
//	// 自动迁移 (这将创建或更新表结构)
//	db.AutoMigrate(&Student{})
//}
//
//func getGrades(c *gin.Context) {
//	name := c.Query("name") // 从URL中获取查询参数
//
//	var student Student
//	// 查询数据库，根据学生的名字获取成绩
//	if err := db.Where("name = ?", name).First(&student).Error; err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"success": false,
//			"message": "未找到该学生的成绩",
//		})
//		return
//	}
//
//	// 如果找到学生，返回成绩
//	grades := map[string]int{
//		"test1": student.Exam1,
//		"test2": student.Exam2,
//		"test3": student.Exam3,
//		"test4": student.Exam4,
//		"test5": student.Exam5,
//		"test6": student.Exam6,
//		"test7": student.Exam7,
//		"test8": student.Exam8,
//		"test9": student.Exam9,
//	}
//	c.JSON(http.StatusOK, gin.H{
//		"success": true,
//		"grades":  grades,
//	})
//}

func main() {
	// 创建一个默认的gin路由器
	router := gin.Default()

	core.InitGorm()

	// 处理 GET 请求 /grades
	httpserver.RegisterRouter(router)
}
