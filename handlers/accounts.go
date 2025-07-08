package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/lucasslima/puc-myfinance-o8/db"
)

type Accounts struct{
  Code int
  Description string
  AccountType string
}

type AccountHandler struct{
	DB *db.DBConnector
}

func TransactionList(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("views/list-template.html.tmpl")
	page.Execute(w,nil)
}

func (accHandler AccountHandler) ListAccounts(w http.ResponseWriter, r *http.Request) {
    log.Printf("ListAccounts handling path %s", r.URL)
    params := r.URL.Query()
    if params.Has("type"){
        accountType := params.Get("type")
        accounts, err := accHandler.accountsByType(accountType)
        if err != nil {
            log.Fatalf("Error fetching accounts with type %s: %v",accountType, err)
            http.Error(w,fmt.Sprintf("Error fetching accounts with type %s: %s", accountType, err),http.StatusInternalServerError)
            return
        }
        fmt.Fprintf(w, "Accounts found: %v\n", accounts)
        
    } else {
        page, _ := template.ParseFiles("views/list-template.html.tmpl")
        page.Execute(w,nil)
    }
}



// Fetches the account with the given code.
func (accountHandler AccountHandler) accountsByType(accountType string) ([]Accounts, error) {
    // An account slice to hold data from returned rows.
    var accounts []Accounts

    rows, err := accountHandler.DB.Connection.Query("SELECT * FROM accounts WHERE type = ? ", accountType)
    if err != nil {
        return nil, fmt.Errorf("Account with the type %q: %v", accountType, err)
    }
    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var acc Accounts
        if err := rows.Scan(&acc.Code, &acc.Description, &acc.AccountType); err != nil {
            return nil, fmt.Errorf("accountsByType %q: %v", accountType, err)
        }
        accounts = append(accounts, acc)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("accountsByType %q: %v", accountType, err)
    }
    return accounts, nil
}

// Fetches the account with the given code.
func (accountHandler AccountHandler) listAllAccounts() ([]Accounts, error) {
    // An account slice to hold data from returned rows.
    var accounts []Accounts

    rows, err := accountHandler.DB.Connection.Query("SELECT * FROM accounts ")
    if err != nil {
        return nil, fmt.Errorf("Account with the type %q: %v", err)
    }
    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var acc Accounts
        if err := rows.Scan(&acc.Code, &acc.Description, &acc.AccountType); err != nil {
            return nil, fmt.Errorf("listAllAccounts: %v", err)
        }
        accounts = append(accounts, acc)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("listAllAcccounts: %v", err)
    }
    return accounts, nil
}