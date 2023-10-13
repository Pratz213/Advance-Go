package WebApp1

import (
	"fmt"
	"log"
	"net/http"
)

func Handler() {
	fmt.Println(8000)
	http.HandleFunc("/", Homepage)
	http.HandleFunc("/get_employees", GetAllEmployees)
	log.Fatal(http.ListenAndServe(":8000", nil))
	fmt.Print("goti")
}
