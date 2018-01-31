package main

import (
	"net/http"
	"html/template"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.tpl")
	t.Execute(w, nil)
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
	w.Write([]byte(r.PostFormValue("address_text")))
	f, fh, err := r.FormFile("address_file")
	if f == nil || fh == nil {
		return
	}
	if err != nil {
		panic(err)
	}
	w.Write([]byte(fh.Filename))

}

func handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/post", 301)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/result", resultHandler)
	http.ListenAndServe(":8080", nil)
}
