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

type Post struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

var posts = []Post{
	{
		Title: "Hey Gophercon!",
		Body:  "Welcome to Alaska!",
	}, {
		Title: "It's really warm in Australia",
		Body:  "This is the first day I've worn pants",
	},
}

// Get root directory + folder
func projectDir(file string) string {
	_, filename, _, _ := runtime.Caller(1)

	// NOTE check me out
	//fmt.Println("FILEPATH: " + filename)
	return path.Join(path.Dir(filename), file)
}

func (a App) Default(res http.ResponseWriter, req *http.Request) {

	tmpl := template.Must(template.ParseFiles(projectDir("views/root.html")))

	err := tmpl.Execute(res, nil)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func (a App) Posts(res http.ResponseWriter, req *http.Request) {

	tmpl := template.Must(template.ParseFiles(projectDir("views/posts.html")))

	err := tmpl.Execute(res, posts)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func (a App) NewPost(res http.ReadResponse, req *http.Request) {
	tmpl := template.Must(template.ParseFiles(projectDir("views/new_post.html")))

	err := tmpl.Execute(res, posts)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

}

func (a App) Root(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, "/posts", http.StatusFound)
}

func main() {
	var a App

	// Root path
	http.HandleFunc("/", a.Root)

	// Serve static asset
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// 1. Posts Path
	http.HandleFunc("/posts", a.Posts)

	// 2. New Post
	http.HandleFunc("/posts/new", a.NewPost)

	fmt.Println("Listening on localhost:8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
