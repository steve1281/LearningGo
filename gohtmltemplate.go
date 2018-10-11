package main

import ("fmt"
		"net/http"
		"html/template")

type NewsAggPage struct {
	Title string
	News string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Web stuff here")
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p:= NewsAggPage{Title: "Amazing News Aggregator", News: "some news"}
	t, _ := template.ParseFiles("basicTemplating.html")
	t.Execute(w, p)
}


func main() {
	fmt.Println("Server on port 8000")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}

