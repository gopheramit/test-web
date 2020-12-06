package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/oklog/ulid"
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
	id := genUlid()
	fmt.Printf("github.com/oklog/ulid %s\n", id.String())
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

func genUlid() ulid.ULID {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id
}
