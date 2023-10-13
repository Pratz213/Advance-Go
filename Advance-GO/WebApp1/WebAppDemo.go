package WebApp1

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Homepage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hey Welcome to page!")
}

var employees = []Employee{
	{Id: 1, Name: "John", Role: "CEO", Department: "Accounts"},
	{Id: 2, Name: "David", Role: "CTO", Department: "Software and Quality analysis"},
	{Id: 3, Name: "Mary", Role: "CFO", Department: "Sales"},
}

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Returning all employees")
	json.NewEncoder(w).Encode(employees)
	fmt.Println("fetched all employees")
}
