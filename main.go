package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Println("start listening on port 3000...")
	http.ListenAndServe(":3000", nil)
}
