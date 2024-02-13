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

func DepartmentsHandler(w http.ResponseWriter, r *http.Request) {
	pathSegments := strings.Split(r.URL.Path, "/")
	id, _ := strconv.Atoi(pathSegments[len(pathSegments)-1])
	jsonData := []byte{}
	var v map[string]interface{}
	err := json.Unmarshal(jsonData, &v)
	dept := &Department{}
	depts := []Department{}
	bodyBytes, err := io.ReadAll(r.Body)

	switch r.Method {
	case "GET":
		if id > 0 {
			dept, err = c.GetDepartment(id)
			jsonData, err = json.MarshalIndent(dept, "", "  ")
		} else {
			depts, err = c.GetDepartments()
			jsonData, err = json.MarshalIndent(depts, "", "  ")
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something wrong"))
		}
		w.Write(jsonData)

	case "POST":

		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		err = json.Unmarshal(bodyBytes, dept)
		if err != nil {
			http.Error(w, "Error parsing request body", http.StatusBadRequest)
			return
		}

		dept, _ := c.CreateDepartment(dept.Name)
		jsonData, _ := json.MarshalIndent(dept, "", "  ")
		w.Write(jsonData)

	case "PUT":

		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		err = json.Unmarshal(bodyBytes, &dept)
		if err != nil {
			http.Error(w, "Error parsing request body", http.StatusBadRequest)
			return
		}

		dept, _ := c.UpdateDepartment(uint(id), dept.Name)
		jsonData, _ := json.MarshalIndent(dept, "", "  ")
		w.Write(jsonData)
	case "DELETE":

		err = c.DeleteDepartment(uint(id))
		jsonData, err = json.MarshalIndent(dept, "", "  ")

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something wrong"))
		}

		w.Write([]byte("Data Deleted"))
	}
}
