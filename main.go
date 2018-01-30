package main

import (
	"net/http"
	"fmt"
	"html/template"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.tpl")
	t.Execute(w, nil)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
	w.Write([]byte(r.PostFormValue("address_text")))
	_, fh, err := r.FormFile("address_file")
	if err != nil {
		panic(err)
	}
	w.Write([]byte(fh.Filename))

}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("StupidL"))
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path)
}

func main() {
	//http.HandleFunc("/", handler)
	http.HandleFunc("/demo", postHandler)
	http.HandleFunc("/save", saveHandler)
	http.ListenAndServe(":8080", nil)
}
