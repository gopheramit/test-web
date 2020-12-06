package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

func home(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
	//w.Write([]byte("HEllo from test"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
	ctx := context.Background()

	opt := option.WithCredentialsFile("C:/Users/admin/Documents/test-web/test-web/test-web-4afe9-firebase-adminsdk-fpwiv-98bebadf15.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
}
