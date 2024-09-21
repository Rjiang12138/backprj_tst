package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "path error", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method wrong", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "hello!")
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parse error:%v", err)
	}
	fmt.Fprintf(w, "parse successfully")
	name := r.FormValue("name")
	email := r.FormValue("email")
	psw := r.FormValue("password")
	fmt.Fprintf(w, "name is %s,email is %s,password is %s", name, email, psw)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/file", fileHandler)
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
