package request

import (
	"fmt"
	"log"
	"net/http"
)

func homepage(m http.ResponseWriter, r *http.Request) {
	fmt.Fprint(m, "WelCome to My API")
}

// func input(
// 	handler.inputHandler()
// )

func Request() {

	http.HandleFunc("/", homepage)
	http.HandleFunc("/input", handler.hii) // user input page

	err := http.ListenAndServe(":8000", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
