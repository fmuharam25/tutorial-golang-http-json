package database

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	. "github.com/fmuharam25/tutorial-golang-http-json/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {

	// Connect to SQLite3 database

	db, err := gorm.Open(sqlite.Open("database/hsi.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func DepartmentSeeder(max int) {
	departments := []Department{}
	for i := 1; i <= max; i++ {
		departments = append(departments, Department{Name: "Department" + strconv.Itoa(i)})
	}
	//Create batch
	Connect().CreateInBatches(departments, 3)

}

func EmployeeSeeder(max int, deptIds ...int) {
	// Create a new random number generator
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random index
	randomIndex := rng.Intn(len(deptIds))

	employees := []Employee{}
	for i := 1; i <= max; i++ {
		employees = append(employees, Employee{Name: "Employee" + strconv.Itoa(i), DepartmentID: uint(deptIds[randomIndex%i])})
	}
	Connect().CreateInBatches(employees, 5)

}
