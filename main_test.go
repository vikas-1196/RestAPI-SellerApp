package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEmployee(t *testing.T) {
	Emp = []Employee{
		Employee{Id: "1", Name: "Test"},
	}
	req, err := http.NewRequest("GET", "/test/v1/a/getEmployee", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetEmployees)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"Id":"1","Name":"Test"}]`
	if rr.Body.String() != expected {
		t.Errorf("wrong response: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCreateEntry(t *testing.T) {

	var jsonStr = []byte(`{"Id":"2","Name":"Test2"}`)

	req, err := http.NewRequest("POST", "/test/v1/a/createEmp", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateEmployee)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"Id":"2","Name":"Test2"}`
	if rr.Body.String() != expected {
		t.Errorf("wrong response: got %v want %v",
			rr.Body.String(), expected)
	}
}
