package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http, ResponseWritter, r http.Request){
	fmt.Fprint(w, "Homepage Endpoint Hit")
}

handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	// router initializer
	handleRequests()
}
