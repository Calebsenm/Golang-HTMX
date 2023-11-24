
package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type Film struct{
	Title     string
	Director  string 
}


func OpenCvs() [][] string{
	file , err := os.Open("./Films.cvs");
	if err != nil{
		fmt.Println(err);
	}

	reader := csv.NewReader(file);
	data , _ := reader.ReadAll();
	defer file.Close();
	
	return data
}


func home(w http.ResponseWriter , r *http.Request){
	tmpl := template.Must(template.ParseFiles("index.html"));
	
	data := OpenCvs()
	var films = map[string][] Film {}	
	
	for i:= 0 ; i < len(data); i++ {
		films["films"] = append(films["films"],	Film{Title:  data[i][0] , Director:  data[i][1] })
	} 

	tmpl.Execute(w, films);
}

func add(w http.ResponseWriter , r *http.Request){
	time.Sleep( 1 * time.Second )
	title := r.PostFormValue("title");
	director := r.PostFormValue("director");

	theData := OpenCvs();

	file , err := os.Create("./Films.cvs");
	defer file.Close();

    if err != nil {
        log.Fatal(err)
    }

	theData = append(theData, []string{title, director})
	
	writer := csv.NewWriter(file);
	err = writer.WriteAll(theData);

	if err != nil{
		log.Fatal("Error " , err)
	}

	tpml := template.Must(template.ParseFiles("index.html"))
	tpml.ExecuteTemplate(w ,"film-list-element", Film{Title: title , Director: director})
} 


func main(){

	fmt.Println("Hello Again");	

	http.HandleFunc("/", home );
	http.HandleFunc("/add-film/", add);

	log.Fatal( http.ListenAndServe(":8084", nil ))
}
