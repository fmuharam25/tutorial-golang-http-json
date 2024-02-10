package web

import (
	"net/http"

	"github.com/fmuharam25/tutorial-golang-http-json/web/handler"
)

func Serve() {
	//Router
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Employee and Department API V1"))
	})
	http.HandleFunc("/employees", handler.EmployeesHandler)
	http.HandleFunc("/employees/{id}", handler.EmployeesHandler)
	http.HandleFunc("/departments", handler.DepartmentsHandler)
	http.HandleFunc("/departments/{id}", handler.DepartmentsHandler)

	//Run Server
	http.ListenAndServe(":8080", nil)
}
