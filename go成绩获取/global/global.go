package global

import "gorm.io/gorm"

var (
	Host     = "root"
	Password = "root"
	Port     = "3307"
	DataBase = "student_db"
)

var (
	DB gorm.DB
)
