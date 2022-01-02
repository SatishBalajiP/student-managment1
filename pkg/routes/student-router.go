package routes

import (
	"github.com/SatishBalajiP/student-managment/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterStudentRoutes = func(router *mux.Router) {
	router.HandleFunc("/student/", controllers.CreateStudent).Methods("POST")
	router.HandleFunc("/student/", controllers.GetStudent).Methods("GET")
	router.HandleFunc("/student/{studentID}", controllers.GetStudentByID).Methods("GET")
	router.HandleFunc("/student/{studentID}", controllers.UpdateStudent).Methods("PUT")
	router.HandleFunc("/student/{studentID}", controllers.DeleteStudent).Methods("DELETE")
}
