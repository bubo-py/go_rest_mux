package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// Employee struct with pointer to Person struct
type Employee struct {
	ID     string  `json:"id"`
	Person *Person `json:"person"`
	Salary string  `json:"salary"`
}

// Person struct
type Person struct {
	Firstname string `json:"first"`
	Lastname  string `json:"last"`
	Age       string `json:"age"`
}

// Mock database that actually is a slice
var employees []Employee

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for _, item := range employees {
		if item.ID == param["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Employee{})
}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee Employee
	_ = json.NewDecoder(r.Body).Decode(&employee)
	employee.ID = strconv.Itoa(len(employees) + 1)
	employees = append(employees, employee)

	json.NewEncoder(w).Encode(employee)
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for i, item := range employees {
		if item.ID == param["id"] {
			employees = append(employees[:i], employees[i+1:]...)
			var employee Employee
			_ = json.NewDecoder(r.Body).Decode(&employee)
			employee.ID = param["id"]
			employees = append(employees, employee)
			return
		}
	}
	json.NewEncoder(w).Encode(employees)

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for i, item := range employees {
		if item.ID == param["id"] {
			employees = append(employees[:i], employees[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(employees)
}

func main() {
	// Initialize router
	r := mux.NewRouter()

	// Mock database initialization
	employees = append(employees, Employee{ID: "1", Person: &Person{Firstname: "Vince", Lastname: "Reuter", Age: "23"}, Salary:          "9090"})
	employees = append(employees, Employee{ID: "2", Person: &Person{Firstname: "Karel", Lastname: "Bullivant", Age: "91"}, Salary:       "2137"})
	employees = append(employees, Employee{ID: "3", Person: &Person{Firstname: "Kain", Lastname: "McGeever", Age: "42"}, Salary:         "9332"})
	employees = append(employees, Employee{ID: "4", Person: &Person{Firstname: "Crawford", Lastname: "Van den Dael", Age: "34"}, Salary: "5322"})

	// URL paths and handlers, some with additional parameter
	r.HandleFunc("/api/employees", getEmployees).Methods("GET")
	r.HandleFunc("/api/employees/{id}", getEmployee).Methods("GET")
	r.HandleFunc("/api/employees/add", addEmployee).Methods("POST")
	r.HandleFunc("/api/employees/update/{id}", updateEmployee).Methods("PUT")
	r.HandleFunc("/api/employees/delete/{id}", deleteEmployee).Methods("DELETE")

	// Work on port 8000 of localhost
	log.Fatal(http.ListenAndServe(":8000", r))
}
