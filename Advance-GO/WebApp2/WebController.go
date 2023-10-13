package webapp2

import (
	"log"
	"net/http"
)

func Handler() {
	http.HandleFunc("/my_demo", CallDemo)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
