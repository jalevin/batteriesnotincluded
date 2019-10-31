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
func projectDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

func projectView(view string) string {
	return path.Join(projectDir(), "views", view+".html")
}

// Handler to root path
func (a App) Root(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, "/posts", http.StatusFound)
}

// Handler to show all posts
func (a App) Posts(res http.ResponseWriter, req *http.Request) {

	tmpl := template.Must(template.ParseFiles(
		projectView("layout"),
		projectView("nav"),
		projectView("posts"),
	))

	err := tmpl.ExecuteTemplate(res, "layout", posts)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

// Handler to show a post
func (a App) ShowPost(res http.ResponseWriter, req *http.Request) {

	tmpl := template.Must(template.ParseFiles(
		projectView("layout"),
		projectView("nav"),
		projectView("show_post"),
	))

	err := tmpl.ExecuteTemplate(res, "layout", posts[1])
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

// Handler to new post form
func (a App) NewPost(res http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		projectView("layout"),
		projectView("nav"),
		projectView("new_post"),
	))

	err := tmpl.ExecuteTemplate(res, "layout", posts)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	var a App

	// Serve static asset
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// B. Root path
	http.HandleFunc("/", a.Root)

	// 1. Posts Path
	http.HandleFunc("/posts", a.Posts)

	// 2. View
	http.HandleFunc("/post/1", a.ShowPost)

	// 3. New Post
	http.HandleFunc("/posts/new", a.NewPost)

	fmt.Println("Listening on localhost:8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
