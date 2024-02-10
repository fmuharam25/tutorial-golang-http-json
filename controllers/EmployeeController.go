package controllers

import (
	"log"

	. "github.com/fmuharam25/tutorial-golang-http-json/model"
)

func CreateEmployee(name string, dept *Department) (*Employee, error) {
	emp := &Employee{Name: name, DepartmentID: dept.ID}
	err := db.Create(emp).Error
	return emp, err
}

func GetEmployee(id *int) *Employee {
	res := db.Find(&Employee{})

	if id != nil {
		res = db.First(&Employee{}, id)
	}

	if res.Error != nil {
		log.Fatal(res.Error)
	}
	if res.RowsAffected > 0 {
		return &Employee{}
	}
	return nil

}

func UpdateEmployee(id uint, name string, deptId uint) (*Employee, error) {
	emp := &Employee{ID: id, Name: name, DepartmentID: deptId}
	err := db.Save(emp).Error
	return emp, err
}

func UpdateBatchEmployee(id []uint, deptId uint) []uint {

	db.Model(Employee{}).Where("id IN ?", id).Updates(Employee{DepartmentID: deptId})
	return id
}

func DeleteEmployee(id uint) error {
	emp := &Employee{ID: id}
	return db.Delete(emp).Error
}
