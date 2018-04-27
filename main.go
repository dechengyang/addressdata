package main

import (
	"addressdata/data"
	"net/http"
	"log"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("/root/files")))
	log.Fatal(http.ListenAndServe(":8088", nil))
	//genFastTextClassifyData()
}

func genFastTextClassifyData() {
	ft := data.FastText{}
	ft.GenClassifyData()
}
