package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"pb4/db"
	"pb4/entity"

	"github.com/sirupsen/logrus"
)

func PostStudent(rw http.ResponseWriter, r *http.Request) {
	reqBody := r.Body

	bodyBytes, err := ioutil.ReadAll(reqBody)
	if hasError(rw, err, "Internal issue") {
		return
	}

	var student = entity.Student{}
	err = json.Unmarshal(bodyBytes, &student)
	if hasError(rw, err, "Internal issue") {
		return
	}

	result := db.GetDB().Create(&student)
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}

	fmt.Println(student)
	rw.Write(bodyBytes)
}

func GetStudent(rw http.ResponseWriter, r *http.Request) {
	idValue := r.URL.Query().Get("id")
	var student entity.Student
	result := db.GetDB().Find(&student, "id=?", idValue)
	if result.RecordNotFound() {
		http.Error(rw, "No Record Found", http.StatusNotFound)
		return
	}

	if result.Error != nil {
		http.Error(rw, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	personData, _ := json.Marshal(student)
	rw.Write(personData)
}

func hasError(rw http.ResponseWriter, err error, message string) bool {
	logger := new(logrus.Entry)
	if err != nil {
		logger.WithError(err).Error(message)
		rw.Write([]byte(message))
		return true
	}
	return false
}

func DeleteStudent(rw http.ResponseWriter, r *http.Request) {
	idValue := r.URL.Query().Get("id")
	result := db.GetDB().Delete(&entity.Student{}, "id=?", idValue)
	if result.Error != nil {
		http.Error(rw, "Internal Error. Please try again after some time", http.StatusInternalServerError)
		return
	}

	rw.Write([]byte("Record successfully deleted"))
}

func PutStudent(rw http.ResponseWriter, r *http.Request) {
	reqBody := r.Body

	bodyBytes, err := ioutil.ReadAll(reqBody)
	if hasError(rw, err, "Internal issue") {
		return
	}

	var inputStudent = entity.Student{}
	err = json.Unmarshal(bodyBytes, &inputStudent)
	if hasError(rw, err, "Internal issue") {
		return
	}

	result := db.GetDB().Model(&entity.Student{}).Updates(inputStudent)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	//fmt.Println(student)
	rw.Write(bodyBytes)
}

func ListOfStudents(rw http.ResponseWriter, r *http.Request) {
	var student []entity.Student
	result := db.GetDB().Find(&student)
	if result.RecordNotFound() {
		http.Error(rw, "No Record Found", http.StatusNotFound)
		return
	}

	if result.Error != nil {
		http.Error(rw, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	personData, _ := json.Marshal(student)
	rw.Write(personData)
}

func Enroll(rw http.ResponseWriter, r *http.Request) {
	body := r.Body
	bodyBytes, err := ioutil.ReadAll(body)
	if hasError(rw, err, "Internal issue") {
		return
	}

	var student entity.Student
	err = json.Unmarshal(bodyBytes, &student)
	if hasError(rw, err, "Internal issue") {
		return
	}

	result := db.GetDB().Model(&student).Association("classes").Append(student.Classes) //replace deletes previos entries in classes field for student
	//it also exists delete associations
	//also clear associations
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}
}
