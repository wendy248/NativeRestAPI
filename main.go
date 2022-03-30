package main

import (
	"Github/NativeRestAPI/controller"
	"fmt"
	"net/http"
)

/*
Catatan :
Untuk melakukan pengujian masing-masing function, diperlukan aplikasi postman karena main_test.go
hanya berisi pengetesan pada function POST saja.
*/

func main() {

	controller.Database[1] = controller.Student{ID: 1, Name: "budi", Age: 5}

	// localhost:8080/student-get/{id}
	http.HandleFunc("/student-get/", controller.GetHandle)

	// localhost:8080/student
	http.HandleFunc("/student", controller.PostHandle)

	// localhost:8080/student-delete/{id}
	http.HandleFunc("/student-delete/", controller.DeleteHandle)

	// localhost:8080/student-update/{id}
	http.HandleFunc("/student-update/", controller.UpdateHandle)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
	}
}
