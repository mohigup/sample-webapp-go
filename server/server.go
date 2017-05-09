package main

import (
	"net/http"
	"html/template"
	"log"
)

func main(){

	templates := populateTemplates()
	// handle every URL
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		// slice slash
		requestedFile := r.URL.Path[1:]
		// get name of templates
		t := templates.Lookup(requestedFile + ".html")
		// error handling
		if t!=nil{
			err := t.Execute(w,nil)
			if err != nil{
				log.Println(err)
			}
		}
	})

	http.Handle("/img/",http.FileServer(http.Dir("public")))
	http.Handle("/css/",http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8000",nil)
}

func populateTemplates() *template.Template{
	// container for all templates
	result := template.New("templates")
	// path for templates
	const BasePath  = "templates"
	// parse templates for result, handles error
	// any file ending with .html will be parsed
	template.Must(result.ParseGlob(BasePath+"/*.html"))
	return result
}