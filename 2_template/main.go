package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"runtime"
)

type App struct{}

//type Post struct {
//Title string `json:"title"`
//Body  string `json:"body"`
//}

// Get root directory + folder
func projectDir(file string) string {
	_, filename, _, _ := runtime.Caller(1)

	// NOTE check me out
	//fmt.Println("FILEPATH: " + filename)
	return path.Join(path.Dir(filename), file)
}

func (a App) Hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello World!")
}

func (a App) Japanese(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Konichiwa")
}

func (a App) Default(res http.ResponseWriter, req *http.Request) {

	tmpl := template.Must(template.ParseFiles(projectDir("views/default.html")))

	err := tmpl.Execute(res, nil)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

//func (a App) Post(res http.ResponseWriter, req *http.Request) {

//tmpl := template.Must(template.ParseFiles(projectDir("views/post.html")))

//p := Post{
//Title: "Hey Gophercon!",
//Body:  "Alaska says hello!",
//}

//err := tmpl.Execute(res, p)
//if err != nil {
//http.Error(res, err.Error(), http.StatusInternalServerError)
//}
//}

func main() {
	var a App

	// A. Hello World
	http.HandleFunc("/hello", a.Hello)
	http.HandleFunc("/japanese", a.Japanese)

	// 1. Respond with a view
	//http.HandleFunc("/", a.Default)

	// 2. Serve static asset
	//http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// 3. Fill out a template
	//http.HandleFunc("/", a.Post)

	fmt.Println("Listening on localhost:8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
