package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hollow")
}

func main() {
    http.HandleFunc(pattern "/", homehomePage)

}
