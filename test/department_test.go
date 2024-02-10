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

func TestDeparment(t *testing.T) {
	database.Migrate()
	//Define default
	dept := Department{ID: 1, Name: "Department1"}
	reqBodyPost, _ := json.Marshal(Department{
		Name: "IT",
	})

	reqBodyPut, _ := json.Marshal(Department{
		Name: "HR",
	})

	testCases := []struct {
		name   string
		method string
		url    string
		body   []byte
	}{

		{"GET Department", "GET", "/departments/1", nil},
		{"POST Department", "POST", "/departments", reqBodyPost},
		{"PUT Department", "PUT", "/departments/3", reqBodyPut},
		{"DELETE Department", "DELETE", "/departments/5", nil},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			// Create a new response recorder
			w := httptest.NewRecorder()
			// Create a new request
			req, _ := http.NewRequest(testCase.method, testCase.url, bytes.NewBuffer(testCase.body))

			// Set the response writer to the recorder
			handler.DepartmentsHandler(w, req)

			// Assert the response status code
			assert.Equal(t, http.StatusOK, w.Code)

			// Assert the response body
			var response Department
			json.Unmarshal(w.Body.Bytes(), &response)

			//Check
			switch testCase.method {
			case "GET":
				assert.Equal(t, dept.ID, response.ID)
				assert.Equal(t, dept.Name, response.Name)
			case "PUT":
				assert.Equal(t, "HR", response.Name)
			case "POST":
				assert.Equal(t, "IT", response.Name)
			case "DELETE":
				assert.Equal(t, http.StatusOK, w.Code)
			}

		})
	}

}
