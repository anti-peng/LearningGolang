package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"switchDB/structures"
)

// Kind of template cache
// cache templates at start up
var templates = template.Must(template.ParseFiles(
	"./statics/html/index.html",
	"./statics/html/page-view.html",
	"./statics/html/page-edit.html",
))

// validation
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-z0-9A-Z]+)$")

func renderTemplate(w http.ResponseWriter, tmpl string, p *structures.Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

// View
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := structures.LoadPage(title)
	// page not found handler
	// redirect to editing page
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "page-view", p)
}

// Edit
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := structures.LoadPage(title)
	if err != nil {
		p = &structures.Page{Title: title}
	}
	renderTemplate(w, "page-edit", p)
}

// Save
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	fmt.Println(body)
	p := &structures.Page{Title: title, Body: []byte(body)}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// save and redirect to page view
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func main() {

	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":3000", nil))
}
