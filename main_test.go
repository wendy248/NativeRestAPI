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

=== RUN   TestPostHandle
--- PASS: TestPostHandle (0.00s)
PASS
ok      Github/NativeRestAPI    0.021s
*/

func TestPostHandle(t *testing.T) {
	var data = []byte(`{"id":1,"name":"budi","age":5}`)
	req, err := http.NewRequest("POST", "/student", bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.PostHandle)
	handler.ServeHTTP(res, req)

	if httpStatus := res.Code; httpStatus != http.StatusOK {
		t.Errorf("http.status code from handler: %v, the code supposed to be %v", httpStatus, http.StatusOK)
	}

	value := `{"Message": "Success input data to table"}`
	if res.Body.String() != value {
		t.Errorf("value from handler: %v, the value supposed to be %v", res.Body.String(), value)
	}
}
