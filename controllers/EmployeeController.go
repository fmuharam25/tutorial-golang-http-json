package controllers

import (
	. "github.com/fmuharam25/tutorial-golang-http-json/model"
)

func CreateEmployee(name string, deptId uint) (*Employee, error) {
	emp := &Employee{Name: name, DepartmentID: deptId}
	err := db.Create(emp).Error
	return emp, err
}

func GetEmployees() ([]Employee, error) {
	emp := []Employee{}
	res := db.Find(&emp)
	if res.Error != nil {
		return nil, res.Error
	}
	return emp, nil
}

func GetEmployee(id int) (*Employee, error) {
	emp := &Employee{}
	res := db.First(emp, id)
	if res.Error != nil {
		return nil, res.Error
	}

	return emp, nil
}

func UpdateEmployee(id uint, name string, deptId uint) (*Employee, error) {
	emp := &Employee{ID: id, Name: name, DepartmentID: deptId}
	res := db.Save(emp)
	if res.Error != nil {
		return nil, res.Error
	}
	return emp, res.Error
}

func UpdateBatchEmployee(id []uint, deptId uint) []uint {

	db.Model(Employee{}).Where("id IN ?", id).Updates(Employee{DepartmentID: deptId})
	return id
}

func DeleteEmployee(id uint) error {
	emp := &Employee{ID: id}
	return db.Delete(emp).Error
}
