package main

import (
	"Github/NativeRestAPI/controller"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
Catatan :
 main_test.go hanya melakukan Unit Test pada function POST dari main.go,
 sehingga proses pengujian fungsi lainnya dilakukan dengan menggunakan program Postman.

*/

func TestPostHandle(t *testing.T) {
	var jsonStr = []byte(`{"id":1,"name":"budi","age":5}`)
	req, err := http.NewRequest("POST", "/student", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.PostHandle)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"Message": "Success input data to table"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
