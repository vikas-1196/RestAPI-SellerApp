package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Employee struct {
	Id   string
	Name string
}

var Emp []Employee
var test Database *gorm.DB1
var MySqlConnection = "root:root@tcp(localhost:3306)/EmployeeDatabase?parseTime=true"
var err error

func Routing() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/test/v1/a/getEmployee", GetEmployees).Methods("GET")
	router.HandleFunc("/test/v1/a/getEmployee/{id}", GetEmployeeById).Methods("GET")
	router.HandleFunc("/test/v1/a/createEmp", CreateEmployee).Methods("POST")
	router.HandleFunc("/test/v1/a/deleteEmp/{id}", DeleteEmployee).Methods("DELETE")
	router.HandleFunc("/test/v1/a/updateEmp/{id}", updateEmp).Methods("PUT")
	log.Fatal(http.ListenAndServe(":9000", router))
}

func updateEmp(w http.ResponseWriter, r *http.Request) {
	val := mux.Vars(r)
	value := val["id"]
	var a Employee
	for i, v := range Emp {
		if v.Id == value {
			resBody, _ := ioutil.ReadAll(r.Body)
			json.Unmarshal(resBody, &a)
			Emp = append(Emp[:i], Emp[i+1:]...)
			Emp = append(Emp, a)
		}
	}
	json.NewEncoder(w).Encode(Emp)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	val := mux.Vars(r)
	value := val["id"]
	for i, v := range Emp {
		if v.Id == value {
			Emp = append(Emp[:i], Emp[i+1:]...)
		}
	}

}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	resBody, _ := ioutil.ReadAll(r.Body)
	var a Employee
	json.Unmarshal(resBody, &a)
	Emp = append(Emp, a)
	json.NewEncoder(w).Encode(a)
}

func GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	val := mux.Vars(r)
	value := val["id"]
	for _, v := range Emp {
		if v.Id == value {
			json.NewEncoder(w).Encode(v)
		}
	}
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Emp)
}

func main() {
	Database, err := gorm.Open("mysql", MySqlConnection)
	if err != nil {
		log.Println("test connection failed1", err)
		return
	}
	log.Println("Connection Established")
	Database.AutoMigrate(&Employee{})
	Routing()
}
