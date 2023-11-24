package main

import (
	"fmt"
	"time"

	"html/template"
	"log"
	"net/http"
)

type Film struct{
	Title     string
	Director  string 
}


func h1(w http.ResponseWriter , r *http.Request){
	tmpl := template.Must(template.ParseFiles("index.html"));
	films := map[string][]Film{
		"films":{
			{Title: "Blade Runner", Director:  "Ridley Scoot" },
		},

	}
	tmpl.Execute(w, films);
}

func h2(w http.ResponseWriter , r *http.Request){
	time.Sleep( 1 * time.Second )
	title := r.PostFormValue("title");
	director := r.PostFormValue("director");

	tpml := template.Must(template.ParseFiles("index.html"))
	tpml.ExecuteTemplate(w ,"film-list-element", Film{Title: title , Director: director})
} 

func main(){

	fmt.Println("Hello Again");	

	http.HandleFunc("/", h1);
	http.HandleFunc("/add-film/", h2);

	log.Fatal( http.ListenAndServe(":8081", nil ))
}