package controllers

import (
	"log"

	. "github.com/fmuharam25/tutorial-golang-http-json/model"
)

func CreateDepartment(name string) (*Department, error) {
	dept := &Department{Name: name}
	err := db.Create(dept).Error
	return dept, err
}

func GetDepartment(id *int) *Department {
	dept := &Department{}
	res := db.Find(dept)

	if id != nil {
		res = db.First(dept, id)
	}

	if res.Error != nil {
		log.Fatal(res.Error)
	}
	if res.RowsAffected > 0 {
		return dept
	}
	return nil
}

func UpdateDepartment(id uint, name string) error {
	dept := &Department{ID: id, Name: name}
	return db.Save(dept).Error
}

func DeleteDepartment(id uint) error {
	dept := &Department{ID: id}
	return db.Delete(dept).Error
}
