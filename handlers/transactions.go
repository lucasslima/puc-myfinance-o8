package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/lucasslima/puc-myfinance-o8/db"
)

type Transaction struct {
	Code        int
	Description string
	Date sql.NullTime
	Account int
	Value float32

}

type TransactionHandler struct {
	DB *db.DBConnector
}


func (entryHandler TransactionHandler) ListTransactions(w http.ResponseWriter, r *http.Request) {
	log.Printf("ListAccounts handling path %s", r.URL)
	var transactions []Transaction
	var err error
	transactions, err = entryHandler.listEntries()
	if err != nil {
		log.Fatalf("Error fetching transactions: %v", err)
		http.Error(w, fmt.Sprintf("Error fetching accounts with: %s",  err), http.StatusInternalServerError)
		return
	}
    // templateContext := &PageBody{
    //     Title: "Account List",
    //     TemplatePath: "views/list-accounts.html.tmpl",
    //     TemplateData: accounts,
    //     }
	page, err := template.ParseFiles("views/list-transactions.html.tmpl")
    if err != nil {
        log.Fatalf("Failed loading template: %s", err)
    }
	page.Execute(w, transactions)
}

func (entryHandler TransactionHandler) listEntries() ([]Transaction, error) {
	var entries []Transaction

	rows, err := entryHandler.DB.Connection.Query("SELECT * FROM transactions ")
	if err != nil {
		return nil, fmt.Errorf("Failed fetching entries: %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var ent Transaction 
		if err := rows.Scan(&ent.Code, &ent.Description, &ent.Date, &ent.Account, &ent.Value); err != nil {
			return nil, fmt.Errorf("listEntries: %v", err)
		}
		entries = append(entries, ent)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("listEntries: %v", err)
	}
	return entries, nil
}
