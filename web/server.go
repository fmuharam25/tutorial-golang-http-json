package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	c "github.com/fmuharam25/tutorial-golang-http-json/controllers"
	"github.com/fmuharam25/tutorial-golang-http-json/database"
	. "github.com/fmuharam25/tutorial-golang-http-json/model"
)

var db = database.Connect()

func main() {
	//Migrate & Seed
	db.AutoMigrate(&Department{}, &Employee{})
	database.DepartmentSeeder(10)
	database.EmployeeSeeder(100)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Employee and Department API V1"))
	})
	http.HandleFunc("/employee/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			id, _ := strconv.Atoi(r.PathValue("id"))
			res := c.GetEmployee(nil)
			if id > 0 {
				res = c.GetEmployee(&id)
			}

			json.NewEncoder(w).Encode(res)

		case "POST":
			w.Write([]byte("POST Employee"))

		case "PUT":
			w.Write([]byte("POST Employee"))

		case "DELETE":
			w.Write([]byte("POST Employee"))
		}

	})

	http.ListenAndServe(":8080", nil)
}
