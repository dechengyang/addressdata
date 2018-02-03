package main

import (
	"net/http"
	"html/template"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("statics/templates/index.html")
		t.Execute(w, nil)
		return
	}
	http.Redirect(w, r, "/404", 301)
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Hello"))
	//w.Write([]byte(r.PostFormValue("address_text")))
	//f, fh, err := r.FormFile("address_file")
	//if f == nil || fh == nil {
	//	return
	//}
	//if err != nil {
	//	panic(err)
	//}
	//w.Write([]byte(fh.Filename))
	if r.Method == "POST" {
		t, _ := template.ParseFiles("statics/templates/result.html")
		t.Execute(w, nil)
		return
	}
	http.Redirect(w, r, "/404", 301)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/post", 301)
		return
	}
	http.Redirect(w, r, "/404", 301)
}

func handler404(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("statics/templates/404.html")
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/result", resultHandler)
	http.HandleFunc("/404", handler404)
	http.ListenAndServe(":8080", nil)
}
