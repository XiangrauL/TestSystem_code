package controller

import (
	"github.com/gin-gonic/gin"
	"go_get_score/dao"
	"go_get_score/model"
	"net/http"
)

type Input struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type InputChange struct {
	Name        string `json:"name" binding:"required"`
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

func GetGrades(c *gin.Context) {
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		model.ReturnFailJson(c, err.Error())
		return
	}
	name := input.Name // 从URL中获取查询参数
	if name == "" {
		model.ReturnFailJson(c, "姓名没有")
		return
	}
	password := input.Password
	if password == "" {
		model.ReturnFailJson(c, "密码不能为空")
		return
	}

	var student dao.Student
	// 查询数据库，根据学生的名字获取成绩
	student, err := dao.SearchStudentDB(name)
	if err != nil {
		model.ReturnFailJson(c, "学生信息不在")
		return
	}
	if student.Password == "" {
		model.ReturnFailJson(c, "未注册")
		return
	}
	if student.Password != password {
		model.ReturnFailJson(c, "密码不正确")
		return
	}

	// 如果找到学生，返回成绩
	grades := map[string]int{
		"test1": student.Exam1,
		"test2": student.Exam2,
		"test3": student.Exam3,
		"test4": student.Exam4,
		"test5": student.Exam5,
		"test6": student.Exam6,
		"test7": student.Exam7,
		"test8": student.Exam8,
		"test9": student.Exam9,
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    grades,
		"msg":     "success",
	})
}

func ChangePassword(c *gin.Context) {
	var input InputChange
	if err := c.ShouldBindJSON(&input); err != nil {
		model.ReturnFailJson(c, err.Error())
		return
	}
	name := input.Name
	if name == "" {
		model.ReturnFailJson(c, "姓名不能为空")
		return
	}
	oldPassword := input.OldPassword
	if oldPassword == "" {
		model.ReturnFailJson(c, "旧密码不能为空")
		return
	}
	newPassword := input.NewPassword
	if newPassword == "" {
		model.ReturnFailJson(c, "新密码不能为空")
		return
	}
	var student dao.Student
	// 查询数据库，根据学生的名字获取成绩
	student, err := dao.SearchStudentDB(name)
	if err != nil {
		model.ReturnFailJson(c, "学生信息不在")
		return
	}
	if student.Password != oldPassword {
		model.ReturnFailJson(c, "旧密码不正确")
		return
	}

	// 注册新密码
	err = dao.RegisterStudentDB(name, newPassword)
	if err != nil {
		model.ReturnFailJson(c, err.Error())
		return
	}

	model.ReturnSuccessJson(c, nil, "success")
}
