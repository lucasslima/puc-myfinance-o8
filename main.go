package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
) 

func serveIndex(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("views/index.html")
	page.Execute(w,nil)
}
func transactionList(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("views/list-template.html.tmpl")
	page.Execute(w,nil)
}

func serveBootstrapCSS(w http.ResponseWriter, r *http.Request) {
	jsPath := fmt.Sprintf("views%s", r.RequestURI)
	log.Printf("Serving JS at %s", jsPath)
	page, err := os.ReadFile(jsPath)
	if err != nil{
		http.NotFound(w,r)
		return
	}
	if strings.HasSuffix(jsPath,".js"){
		w.Header().Set("Content-Type", "text/javascript")
	} else {
		w.Header().Set("Content-Type", "text/css")
	}
	w.Write(page)
}


func main()  {
    http.HandleFunc("/", serveIndex)
    http.HandleFunc("/transactions", transactionList)
    http.HandleFunc("/bootstrap/", serveBootstrapCSS)
	http.HandleFunc("/js/", serveBootstrapCSS)
	http.HandleFunc("/css/", serveBootstrapCSS)
	defer http.ListenAndServe(":8080", nil)
    log.Printf("Started the server port at: %d", 8080)
}