package test_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fmuharam25/tutorial-golang-http-json/database"
	. "github.com/fmuharam25/tutorial-golang-http-json/model"
	"github.com/fmuharam25/tutorial-golang-http-json/web/handler"

	"github.com/stretchr/testify/assert"
)

func TestEmployee(t *testing.T) {
	database.Migrate()
	//Define default
	emp := Employee{ID: 1, Name: "Employee1", DepartmentID: 1}
	reqBodyPost, _ := json.Marshal(Employee{
		Name:         "John",
		DepartmentID: 1,
	})

	reqBodyPut, _ := json.Marshal(Employee{
		Name:         "Doe",
		DepartmentID: 1,
	})

	testCases := []struct {
		name   string
		method string
		url    string
		body   []byte
	}{

		{"GET Employe", "GET", "/employees/1", nil},
		{"POST Employe", "POST", "/employees", reqBodyPost},
		{"PUT Employe", "PUT", "/employees/4", reqBodyPut},
		{"DELETE Employe", "DELETE", "/employees/1", nil},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			// Create a new response recorder
			w := httptest.NewRecorder()
			// Create a new request
			req, _ := http.NewRequest(testCase.method, testCase.url, bytes.NewBuffer(testCase.body))

			// Set the response writer to the recorder
			handler.EmployeesHandler(w, req)

			// Assert the response status code
			assert.Equal(t, http.StatusOK, w.Code)

			// Assert the response body
			var response Employee
			json.Unmarshal(w.Body.Bytes(), &response)

			//Check
			switch testCase.method {
			case "GET":
				assert.Equal(t, emp.ID, response.ID)
				assert.Equal(t, emp.Name, response.Name)
			case "PUT":
				assert.Equal(t, "Doe", response.Name)
			case "POST":
				assert.Equal(t, "John", response.Name)
			case "DELETE":
				assert.Equal(t, http.StatusOK, w.Code)
			}

		})
	}

}
