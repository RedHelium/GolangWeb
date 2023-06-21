package web

import (
	"html/template"
	"net/http"
	"regexp"
)

var ValidPath = regexp.MustCompile("^/(" + "edit|save|view" + ")" + "/([a-zA-Z0-9]+)$")
var GeneralPath = regexp.MustCompile("^/$")

func MakeHandler(fn func(http.ResponseWriter, *http.Request, string), url *regexp.Regexp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := url.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[len(m)-1])
	}
}

func MainHandler(w http.ResponseWriter, r *http.Request, title string) {

	renderPage(w, "index", nil)
}

func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadMarkdownPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderPage(w, "view", p)
}

func EditHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)

	if err != nil {
		p = &Page{Title: title}
	}
	renderPage(w, "edit", p)
}

func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	header := r.FormValue("title")

	p := &Page{Title: header, BodyHTML: template.HTML(body), Body: []byte(body)}

	// if header != title {
	// 	err := p.remove()
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// }
	err := p.save()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+header, http.StatusFound)

}
