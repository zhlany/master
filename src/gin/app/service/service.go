package service

import (
	"zhl/src/gin/app/model"
)

type StudentService interface {
	FindAll() (Students []model.Student, err error)
	ByName() (Student model.Student, err error)
}
