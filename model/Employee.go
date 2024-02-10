package model

type Employee struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	DepartmentID uint   `json:"department_id"`
}
