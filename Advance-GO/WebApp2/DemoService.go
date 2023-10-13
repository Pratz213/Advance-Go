package webapp2

import (
	"fmt"
	"net/http"
)

func CallDemo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Welcome to Demo Page")
}
