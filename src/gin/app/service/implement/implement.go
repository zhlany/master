package service

import (
	"zhl/src/gin/app/model"
	gro "zhl/src/gin/databases"
	"zhl/src/gin/errors"
)

type Student model.Student

func (s *Student) FindAll() (Students []*Student, err error) {
	if err := gro.Db.Find(&Students).Error; err != nil {
		return nil, errors.NewErrorf("Find All Student failed， err :%s", err)
	}
	return
}
func (s *Student) ByName(name string) (Student *Student, err error) {
	if err := gro.Db.Find(name, &Student).Error; err != nil {
		return nil, errors.NewErrorf("Find Student ByName failed， err :%s", err)
	}
	return
}

type NameSSSS string
