package main

import (
	"log"
	"net/http"
	"web/packages/web"
)

//TODO
//Exclude head tag from template in base html and implement other templates by base.html
//Validation path exclude to another golang file and make by urls.py django analog
//Make a handlers system

func main() {

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))

	http.HandleFunc("/", web.MakeHandler(web.MainHandler, web.GeneralPath))
	http.HandleFunc("/view/", web.MakeHandler(web.ViewHandler, web.ValidPath))
	http.HandleFunc("/edit/", web.MakeHandler(web.EditHandler, web.ValidPath))
	http.HandleFunc("/save/", web.MakeHandler(web.SaveHandler, web.ValidPath))

	log.Fatal(http.ListenAndServe(":8080", nil))

}
