package main

// DONE
// A. Hello World
// 1. respond to request

// TODO
// NOTE add logging to this
// 2. respond with view (templating)
// 3. Basic server
// 4.	Routing Mux (REST)
// 5. Sessions and Cookies
// 6. Database connections SQLite
// 7. SSL - https://blog.cloudflare.com/exposing-go-on-the-internet/
// 8.	Clean shutdown (Timeouts / Context package)

// Go Ricebox for embedding/packaging - https://github.com/GeertJohan/go.rice

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

var runDir string
var templates *template.Template

func funcName() string {
	pc, _, _, _ := runtime.Caller(1)
	nameFull := runtime.FuncForPC(pc).Name() // main.foo
	nameEnd := filepath.Ext(nameFull)        // .foo
	name := strings.TrimPrefix(nameEnd, ".") // foo
	return name
}

// Get full path of directory in project folder
func projectDir(folder string) string {
	// Locate from the runtime the location of the apps static files.
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

	// load template
	//t := loadTemplate("default")

	// FIXME - cache template. do a lookup in slice
	//templates = append(templates, t)

	// load templates
	templates = append(templates, template.Must(template.ParseFiles(
		"/Users/jefe/projects/batteriesnotincluded/go/views/default.html",
	)))

	err := templates.ExecuteTemplate(res, "default.html", nil)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

//func loadTemplate(name string) template.Template {
//templateName := template.Parse("views/" + name + ".html")
//t := template.Must(templateName)
//return t
//}

func main() {
	//// setup

	var a App
	// load runtime directory

	//// execution

	// A. Hello World
	http.HandleFunc("/hello", a.Hello)
	http.HandleFunc("/japanese", a.Japanese)

	// 1. Respond with a view
	http.HandleFunc("/", a.Default)

	// 2. Serve static asset
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	//http.Handle("/assets/img", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// 3. Add JS to page

	// Handle a route

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
