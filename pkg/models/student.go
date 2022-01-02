package models

import (
	"github.com/SatishBalajiP/student-managment/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Student struct {
	gorm.Model
	Name  string `gorm:""json:"name"`
	Grade string `json:"grade"`
	Class string `json:"class"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Student{})
}

func (s *Student) CreateStudent() *Student {
	db.NewRecord(s)
	db.Create(&s)
	return s
}

func GetAllStudents() []Student {
	var Students []Student
	db.Find(&Students)
	return Students
}

func GetStudentById(Id int64) (*Student, *gorm.DB) {
	var getStudent Student
	db := db.Where("ID=?", Id).Find(&getStudent)
	return &getStudent, db
}

func DeleteStudent(Id int64) Student {
	var student Student
	db.Where("Id=?", Id).Delete(student)
	return student
}
