package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"pb4/db"
	"pb4/entity"
)

func PostClass(rw http.ResponseWriter, r *http.Request) {
	reqBody := r.Body

	bodyBytes, err := ioutil.ReadAll(reqBody)
	if hasError(rw, err, "Internal issue") {
		return
	}

	var class = entity.Class{}
	err = json.Unmarshal(bodyBytes, &class)
	if hasError(rw, err, "Internal issue") {
		return
	}

	db.GetDB().Create(&class)

	fmt.Println(class)
	rw.Write(bodyBytes)
}

func GetClass(rw http.ResponseWriter, r *http.Request) {
	idValue := r.URL.Query().Get("id")
	var class entity.Class
	result := db.GetDB().Find(&class, "id=?", idValue)
	if result.RecordNotFound() {
		http.Error(rw, "No Record Found", http.StatusNotFound)
		return
	}

	if result.Error != nil {
		http.Error(rw, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	personData, _ := json.Marshal(class)
	rw.Write(personData)
}

func DeleteClass(rw http.ResponseWriter, r *http.Request) {
	idValue := r.URL.Query().Get("id")
	result := db.GetDB().Delete(&entity.Class{}, "id=?", idValue)
	if result.Error != nil {
		http.Error(rw, "Internal Error. Please try again after some time", http.StatusInternalServerError)
		return
	}

	rw.Write([]byte("Record successfully deleted"))
}

func PutClass(rw http.ResponseWriter, r *http.Request) {
	reqBody := r.Body

	bodyBytes, err := ioutil.ReadAll(reqBody)
	if hasError(rw, err, "Internal issue") {
		return
	}

	var class = entity.Class{}
	err = json.Unmarshal(bodyBytes, &class)
	if hasError(rw, err, "Internal issue") {
		return
	}

	db.GetDB().Update(&class)

	fmt.Println(class)
	rw.Write(bodyBytes)
}

func ListOfClasses(rw http.ResponseWriter, r *http.Request) {
	var class []entity.Class
	result := db.GetDB().Find(&class)
	if result.RecordNotFound() {
		http.Error(rw, "No Record Found", http.StatusNotFound)
		return
	}

	if result.Error != nil {
		http.Error(rw, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	personData, _ := json.Marshal(class)
	rw.Write(personData)
}
