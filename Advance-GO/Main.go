package main

import (
	webapp1 "advance-go/WebApp/WebApp1"
	webapp2 "advance-go/WebApp/WebApp2"
	"fmt"
)

func main() {
	fmt.Print("Application Started at port: ")
	webapp1.Handler()
	fmt.Println("jnascui")
	webapp2.Handler()
}
