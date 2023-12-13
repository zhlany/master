package service

import (
	"fmt"
	"zhl/src/gin/app/model"
	gro "zhl/src/gin/databases"
	_default "zhl/src/gin/default"
	"zhl/src/gin/errors"
)

type Te _default.Te

type StudentService interface {
	FindAll() (Students []model.Student, err error)
	ByName() (Student model.Student, err error)
	CreateStudents(Students []*model.Student) bool
}

func (i Te) FindAll() (Students []model.Student, err error) {
	if err := gro.DB.Find(&Students).Error; err != nil {
		return nil, errors.NewErrorf("Find All Student failed， err :%s", err)
	}
	return
}
func (i Te) ByName(name string) (Student model.Student, err error) {
	if err := gro.DB.Find(name, Student).Error; err != nil {
		return Student, errors.NewErrorf("Find Student ByName failed， err :%s", err)
	}
	return
}

func (i Te) CreateStudents(Students []*model.Student) bool {
	fmt.Println("Student:::::", Students)
	result := gro.DB.CreateInBatches(Students, len(Students))
	fmt.Println("result:", result)
	if result.Error != nil {
		fmt.Println("Failed to create batch:", result.Error)
		return false
	}
	return true
}
