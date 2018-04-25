package main

import (
	"net/http"
	"log"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("/root/files")))
	log.Fatal(http.ListenAndServe(":8088", nil))
}
