package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Salary struct {
	Basic, HRA float64
}

type Employee struct {
	FirstName, LastName string
	Age                 int
	MonthlySalary       []Salary
}

// Handler to show how to serialize a struct using JSON
func johnEmployeeHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	johnEmployee := Employee{
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
		MonthlySalary: []Salary{

			{
				Basic: 16000.00,
				HRA:   5000.00,
			}, // Note the comma here and in the next line, when the array for Salary is closed

		},
	}

	// Note that the method takes a pointer, not the actual object we created above
	employee, err := json.Marshal(&johnEmployee)

	if err != nil {
		fmt.Println(err)
		return
	}

	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, string(employee))
	}
}

// This handler shows a simple case that outputs "Hello!"
func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func main() {
	fmt.Println("Starting server at port 8080\n")

	http.HandleFunc("/johnEmployee", johnEmployeeHandler)
	http.HandleFunc("/hello", helloHandler)

	// Second argument to http.ListenAndServe is a handler to configure the server for HTTP/2, but it's not important in the tutorial
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// Code I didn't use but thought the snippet would be useful

// if req.URL.Path != "/johnEmployee" {
// 	http.Error(w, "404 not found.", http.StatusNotFound)
// 	return
// }
