package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Student struct {
	ID   int    `json:"id" binding:"primary_key"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var Database = make(map[int]Student)

func runJSON(res http.ResponseWriter, message []byte, httpCode int) {
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(httpCode)
	res.Write(message)
}

// Function GET
func GetHandle(res http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		message := []byte(`{"message": "HTTP method is not compatible"}`)
		runJSON(res, message, http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 3 {
		message := []byte(`{"message": "URL is not correct"}`)
		runJSON(res, message, http.StatusNotFound)
		return
	}

	value, err := strconv.Atoi(string(parts[2]))
	if err != nil {
		message := []byte(`{"message": "Error while parsing ID URL to int"}`)
		runJSON(res, message, http.StatusInternalServerError)
		return
	}

	student, ok := Database[value]
	if !ok {
		message := []byte(`{"message": "Data not found"}`)
		runJSON(res, message, http.StatusOK)
		return
	}

	// Parse data to JSON (output)
	studentJSON, err := json.Marshal(&student)
	if err != nil {
		message := []byte(`{"message": "Parsing data to JSON output error"}`)
		runJSON(res, message, http.StatusInternalServerError)
		return
	}

	runJSON(res, studentJSON, http.StatusOK)
}

// Function POST
func PostHandle(res http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		message := []byte(`{"message": "HTTP method is not compatible"}`)
		runJSON(res, message, http.StatusMethodNotAllowed)
		return
	}

	// Decode JSON input
	var student Student
	payload := r.Body
	defer r.Body.Close()

	err := json.NewDecoder(payload).Decode(&student)
	if err != nil {
		message := []byte(`{"Message": "Parsing JSON as input error"}`)
		runJSON(res, message, http.StatusInternalServerError)
		return
	}

	Database[student.ID] = student
	message := []byte(`{"Message": "Success input data to table"}`)
	runJSON(res, message, http.StatusOK)
}

// Function DELETE
func DeleteHandle(res http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		message := []byte(`{"message": "HTTP method is not compatible"}`)
		runJSON(res, message, http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 3 {
		message := []byte(`{"message": "URL is not correct"}`)
		runJSON(res, message, http.StatusNotFound)
		return
	}

	value, err := strconv.Atoi(string(parts[2]))
	if err != nil {
		message := []byte(`{"message": "Error while parsing ID URL to int"}`)
		runJSON(res, message, http.StatusInternalServerError)
		return
	}

	_, ok := Database[value]
	if !ok {
		message := []byte(`{"message": "Data not found"}`)
		runJSON(res, message, http.StatusOK)
		return
	}

	delete(Database, value)
	message := []byte(`{"message": "Data is deleted from table"}`)
	runJSON(res, message, http.StatusOK)
}

// Function UPDATE / PUT
func UpdateHandle(res http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		message := []byte(`{"message": "HTTP method is not compatible"}`)
		runJSON(res, message, http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 3 {
		message := []byte(`{"message": "URL is not correct"}`)
		runJSON(res, message, http.StatusNotFound)
		return
	}

	value, err := strconv.Atoi(string(parts[2]))
	if err != nil {
		message := []byte(`{"message": "Error while parsing ID URL to int"}`)
		runJSON(res, message, http.StatusInternalServerError)
		return
	}

	student, ok := Database[value]
	if !ok {
		message := []byte(`{"message": "Data not found"}`)
		runJSON(res, message, http.StatusOK)
		return
	}

	var updateData Student
	payload := r.Body
	defer r.Body.Close()

	err2 := json.NewDecoder(payload).Decode(&updateData)
	if err2 != nil {
		message := []byte(`{"Message": "Parsing JSON as input error"}`)
		runJSON(res, message, http.StatusInternalServerError)
		return
	}

	student.Name = updateData.Name
	Database[value] = student
	studentJSON, err := json.Marshal(&student)

	if err != nil {
		message := []byte(`{"Message": "Error while parsing JSON output"}`)
		runJSON(res, message, http.StatusInternalServerError)
		return
	}

	runJSON(res, studentJSON, http.StatusOK)
}
