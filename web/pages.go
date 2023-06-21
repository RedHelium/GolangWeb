package web

import (
	"html/template"
	"net/http"
	"os"
	"web/packages/extensions"
)

// Wiki markdown page
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

// Load page by title
func loadPage(title string) (*Page, error) {
	filename := title + PAGE_EXTENSION
	body, err := os.ReadFile(PAGE_FOLDER + filename)

	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil

}

// Convert array bytes file with markdown in HTML
func loadMarkdownPage(title string) (*Page, error) {

	filename := title + ".md"
	body, err := os.ReadFile("pages/" + filename)

	if err != nil {
		return nil, err
	}
	return &Page{Title: title, BodyHTML: template.HTML(extensions.MDToHTML(body))}, nil
}

// Render page by template name
func renderPage(w http.ResponseWriter, templateName string, p *Page) {
	t, err := template.ParseFiles(TEMPLATE_FOLDER+templateName+".html", TEMPLATE_FOLDER+"head.html", TEMPLATE_FOLDER+"navigation.html", TEMPLATE_FOLDER+"footer.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
