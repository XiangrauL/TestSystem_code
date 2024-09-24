package dao

import (
	"go_get_score/global"
)

type Student struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"unique;not null"` // 姓名字段，唯一且不能为空
	Password string
	Exam1    int
	Exam2    int
	Exam3    int
	Exam4    int
	Exam5    int
	Exam6    int
	Exam7    int
	Exam8    int
	Exam9    int
}

func SearchStudentDB(name string) (Student, error) {
	var student Student
	db := global.DB
	err := db.Where("name = ?", name).First(&student).Error
	if err != nil {
		return student, err
	}
	return student, nil
}

func RegisterStudentDB(name, password string) error {
	var student Student
	db := global.DB
	err := db.Where("name = ?", name).First(&student).Error
	if err != nil {
		return err
	}

	student.Password = password
	//fmt.Println(student)
	err = db.Save(&student).Error
	if err != nil {
		return err
	}

	return nil
}
