package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	c "github.com/fmuharam25/tutorial-golang-http-json/controllers"
	. "github.com/fmuharam25/tutorial-golang-http-json/model"
)

func EmployeesHandler(w http.ResponseWriter, r *http.Request) {
	pathSegments := strings.Split(r.URL.Path, "/")
	id, _ := strconv.Atoi(pathSegments[len(pathSegments)-1])
	jsonData := []byte{}
	var v map[string]interface{}
	err := json.Unmarshal(jsonData, &v)
	emp := &Employee{}
	emps := []Employee{}

	switch r.Method {
	case "GET":

		if id > 0 {
			emp, err = c.GetEmployee(id)
			jsonData, err = json.MarshalIndent(emp, "", "  ")
		} else {
			emps, err = c.GetEmployees()
			jsonData, err = json.MarshalIndent(emps, "", "  ")
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something wrong"))
		}
		w.Write(jsonData)

	case "POST":

		bodyBytes, err := io.ReadAll(r.Body)
		err = json.Unmarshal(bodyBytes, emp)
		if err != nil {
			http.Error(w, "Error parsing request body", http.StatusBadRequest)
			return
		}

		emp, _ := c.CreateEmployee(emp.Name, emp.DepartmentID)
		jsonData, _ := json.MarshalIndent(emp, "", "  ")
		w.Write(jsonData)

	case "PUT":

		bodyBytes, err := io.ReadAll(r.Body)
		err = json.Unmarshal(bodyBytes, &emp)
		if err != nil {
			http.Error(w, "Error parsing request body", http.StatusBadRequest)
			return
		}

		emp, _ := c.UpdateEmployee(uint(id), emp.Name, emp.DepartmentID)
		jsonData, _ := json.MarshalIndent(emp, "", "  ")

		w.Write(jsonData)
	case "DELETE":

		err = c.DeleteEmployee(uint(id))
		jsonData, err = json.MarshalIndent(emp, "", "  ")

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something wrong"))
		}

		w.Write([]byte("Data Deleted"))
	}
}
