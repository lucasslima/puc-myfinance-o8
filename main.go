package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/lucasslima/puc-myfinance-o8/db"
	"github.com/lucasslima/puc-myfinance-o8/handlers"
) 

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
	db := db.CreateConnection()
	accountHandler := handlers.AccountHandler{DB: db}
	transactionHandler := handlers.TransactionHandler{DB: db}
    http.HandleFunc("/bootstrap/", serveBootstrapCSS)
	http.HandleFunc("/js/", serveBootstrapCSS)
	http.HandleFunc("/css/", serveBootstrapCSS)
	http.HandleFunc("/accounts", accountHandler.ListAccounts)
	http.HandleFunc("/transactions", transactionHandler.ListTransactions)
	http.HandleFunc("/", transactionHandler.ListTransactions)
	defer http.ListenAndServe(":8080", nil)
    log.Printf("Started the server port at: %d", 8080)
}