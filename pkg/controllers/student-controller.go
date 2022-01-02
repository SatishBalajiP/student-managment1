package contrllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SatishBalajiP/student-managment/pkg/models"
	"github.com/SatishBalajiP/student-managment/pkg/utils"
	"github.com/gorilla/mux"
)

var NewStudent models.Student

func GetStudent(w http.ResponseWriter, r *http.Request) {
	newStudents := models.GetAllStudents()
	res, _ := json.Marshal(newStudents)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	StudentId := vars["StudentId"]
	ID, err := strconv.ParseInt(StudentId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	studentDetails, _ := models.GetStudentById(ID)
	res, _ := json.Marshal(studentDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	CreateStudent := &models.Student{}
	utils.ParseBody(r, CreateStudent)
	s := CreateStudent.CreateStudent()
	res, _ := json.Marshal(s)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentId := vars["StudentId"]
	ID, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	student := models.DeleteStudent(ID)
	res, _ := json.Marshal(student)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	updateStudent := &models.UpdateStudent{}
	utils.ParseBody(r, updateStudent)
	vars := mux.Vars(r)
	studentId := vars["Studentid"]
	ID, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	studentDetails, db := models.GetStudentById(ID)
	if updateStudent.Name != "" {
		studentDetails.Name = updateStudent.Name
	}
	if updateStudent.Grade != "" {
		studentDetails.Grade = updateStudent.Grade
	}
	if updateStudent.Class != "" {
		studentDetails.Class = updateStudent.Class
	}
	db.Save(&studentDetails)
	res, _ := json.Marshal(studentDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
