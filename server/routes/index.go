package routes

import (
	"html/template"
	"net/http"
	"path"
)

type ViewData struct{
    Title string
    Message string
}


func Index(w http.ResponseWriter, r *http.Request) {
    fp := path.Join("html", "index.html")
    tmpl, err := template.ParseFiles(fp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }


	data := ViewData{
		Title: "World Cup",
		Message: "FIFA will never regret it",
	} 

    if err := tmpl.Execute(w, data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}