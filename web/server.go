package web

import (
	"net/http"

	"github.com/fmuharam25/tutorial-golang-http-json/database"
	. "github.com/fmuharam25/tutorial-golang-http-json/model"
	"github.com/fmuharam25/tutorial-golang-http-json/web/handler"
)

var db = database.Connect()

func Serve() {
	//Migrate & Seed
	db.Migrator().DropTable(&Department{}, &Employee{})
	db.AutoMigrate(&Department{}, &Employee{})
	database.DepartmentSeeder(10)
	database.EmployeeSeeder(100, 1, 2, 3)

	//Router
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Employee and Department API V1"))
	})
	http.HandleFunc("/employee", handler.EmployeesHandler)
	http.HandleFunc("/employee/{id}", handler.EmployeesHandler)
	http.HandleFunc("/department", handler.DepartmentsHandler)
	http.HandleFunc("/department/{id}", handler.DepartmentsHandler)

	//Run Server
	http.ListenAndServe(":8080", nil)
}
