package main

import (
	"fmt"
)

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

type App struct{}

var templates *template.Template

// Get root directory + folder
func projectDir(folder string) string {
	_, filename, _, _ := runtime.Caller(1)
	fmt.Println("FILEPATH: " + filename)
	return path.Join(path.Dir(filename), folder)
}

// Default is the default greeting.
func (a App) Hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello World!")
}

func (a App) Japanese(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Konichiwa")
}

func (a App) Default(res http.ResponseWriter, req *http.Request) {

	templates = append(templates, template.Must(template.ParseFiles(
		projectDir("views")+"default.html",
	)))

	err := templates.ExecuteTemplate(res, "default.html", nil)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	var a App

	// A. Hello World
	http.HandleFunc("/hello", a.Hello)
	http.HandleFunc("/japanese", a.Japanese)

	// 1. Respond with a view
	http.HandleFunc("/", a.Default)

	// 2. Serve static asset
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
