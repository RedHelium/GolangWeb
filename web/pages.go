package web

import (
	"html/template"
	"net/http"
	"os"
	"web/packages/extensions"
)

type Page struct {
	Title    string
	BodyHTML template.HTML
	Body     []byte
}

// Save page
func (p *Page) save() error {
	filename := p.Title + ".md"

	return os.WriteFile("pages/"+filename, p.Body, 0600)
}

// Remove page
// func (p *Page) remove() error {

// 	return os.Remove(p.Title)
// }

// Load page by title
func loadPage(title string) (*Page, error) {
	filename := title + ".md"
	body, err := os.ReadFile("pages/" + filename)

	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil

}

func loadMarkdownPage(title string) (*Page, error) {

	filename := title + ".md"
	body, err := os.ReadFile("pages/" + filename)

	if err != nil {
		return nil, err
	}
	return &Page{Title: title, BodyHTML: template.HTML(extensions.MDToHTML(body))}, nil
}

func renderPage(w http.ResponseWriter, templateName string, p *Page) {
	t, err := template.ParseFiles("templates/"+templateName+".html", "templates/head.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
