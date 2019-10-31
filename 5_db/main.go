package main

import (
	"crawshaw.io/sqlite"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
	"path"
	"runtime"
)

type App struct{}

type Post struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func findPostById(id string) Post {
	for _, p := range posts {
		if p.Id == id {
			return p
		}
	}
	return Post{}
}

var posts = []Post{
	{
		Id:    "1",
		Title: "Hey Gophercon!",
		Body:  "Welcome to Alaska!",
	}, {
		Id:    "2",
		Title: "It's really warm in Australia",
		Body:  "This is the first day I've worn pants",
	},
}

type PageData struct {
	PageTitle string
	Flash     string
	Post
	Posts []Post
}

/////Utility functions ///////////////

// Get root directory + folder
func projectDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

// Get static path to template file
func projectView(view string) string {
	return path.Join(projectDir(), "views", view+".html")
}

// Build templates for us
func buildView(view string) *template.Template {
	tmpl := template.Must(template.ParseFiles(
		projectView("layout"),
		projectView("nav"),
		projectView(view),
	))

	return tmpl
}

//////End Utility Functions ///////////////

// Handler to root path
func (a App) Root(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	http.Redirect(res, req, "/posts", http.StatusFound)
}

// Handler to show all posts
func (a App) Posts(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	// easier to do
	tmpl := buildView("posts")
	pd := PageData{
		PageTitle: "Hello Gophercon!",
		Posts:     posts,
	}

	// easier to understand what's going on??
	err := tmpl.ExecuteTemplate(res, "layout", pd)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

// Handler to new post form
func (a App) NewPost(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	tmpl := buildView("new_post")

	pd := PageData{
		PageTitle: "New Post",
		Post:      Post{},
	}

	err := tmpl.ExecuteTemplate(res, "layout", pd)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

// Handler to show a post
func (a App) ShowPost(res http.ResponseWriter, req *http.Request, params httprouter.Params) {

	tmpl := buildView("show_post")

	// Standard server error, but don't crash
	p := findPostById(params.ByName("id"))
	if p.Id == "" {
		msg := fmt.Sprintf("No post with id %v", params.ByName("id"))
		http.Error(res, msg, http.StatusInternalServerError)
	}

	pd := PageData{
		PageTitle: p.Title,
		Post:      p,
	}

	err := tmpl.ExecuteTemplate(res, "layout", pd)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

// Handler to new post form
func (a App) CreatePost(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	err := req.ParseForm()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	p := Post{
		Id:    "3",
		Title: req.PostFormValue("title"),
		Body:  req.PostFormValue("body"),
	}

	posts = append(posts, p)

	http.Redirect(res, req, "/post/3", http.StatusFound)
}

func main() {
	var a App

	router := httprouter.New()
	router.ServeFiles("/assets/*filepath", http.Dir(path.Join(projectDir(), "assets")))
	router.GET("/", a.Root)

	router.GET("/posts", a.Posts)
	router.GET("/posts/new", a.NewPost)
	router.POST("/posts/create", a.CreatePost)

	router.GET("/post/:id", a.ShowPost)
	//router.GET("/post/:id/edit", a.ShowPost)

	fmt.Println("Listening on localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
