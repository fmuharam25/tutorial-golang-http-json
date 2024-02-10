package controllers

import (
	. "github.com/fmuharam25/tutorial-golang-http-json/model"
)

func CreateDepartment(name string) (*Department, error) {
	dept := &Department{Name: name}
	err := db.Create(dept).Error
	return dept, err
}

func GetDepartments() ([]Department, error) {
	dept := []Department{}
	res := db.Find(&dept)
	if res.Error != nil {
		return nil, res.Error
	}
	return dept, nil
}

func GetDepartment(id int) (*Department, error) {
	dept := &Department{}
	res := db.First(dept, id)
	if res.Error != nil {
		return nil, res.Error
	}

	return dept, nil
}

func UpdateDepartment(id uint, name string) (*Department, error) {
	emp := &Department{ID: id, Name: name}
	res := db.Save(emp)
	if res.Error != nil {
		return nil, res.Error
	}
	return emp, res.Error
}

func DeleteDepartment(id uint) error {
	dept := &Department{ID: id}
	return db.Delete(dept).Error
}
