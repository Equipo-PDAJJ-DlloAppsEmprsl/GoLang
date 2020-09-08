package music

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http, ResponseWritter, r http.Request) {
	fmt.Fprint(w, "")
	log.Println("Page has been initialized")
}
