package model

type Department struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"` //add unique if u want
	Employees []Employee
}
