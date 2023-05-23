package main

import
(
 _"github.com/go-sql-driver/mysql"
 "database/sql"
 "testDatabase"
 "log"
 "net/http"
 "fmt"
 "encoding/json"
 "types"
)

var connectionString = "cushon_test_user:12testPWD@tcp(127.0.0.1:3306)/cushon_test"
var db *sql.DB
var err error



func main() {
	// open a database connection once and pass its handle to avoid repeated opening and closing
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic("Database error")
	}
	
	// start webserver
	http.HandleFunc("/funds", getFunds)
	http.HandleFunc("/deposit", makeDeposit)
	log.Fatal(http.ListenAndServe(":8080", nil))	
}

func getFunds (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	jsonResponse, err := json.Marshal(testDatabase.GetFunds(db))
	if err == nil && r.Method == "GET" {
		fmt.Fprint(w, string(jsonResponse))
	} else {
		fmt.Fprint(w, `{"success" : false}`)
	}
}

func makeDeposit (w http.ResponseWriter, r *http.Request) {
	success := true
	// to handle OPTIONS "preflighted" request being sent - only due to cross-domain calls here
	if r.Method != "POST" {
		success = false
	}
	var investment types.Investment
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&investment)
	if err != nil {
		success = false
	}
	if success {
		// check authorisation header key against stored value for customer
		authorisation := r.Header.Get("Authorisation")
		if authorisation != testDatabase.GetAuthorisation(db, investment.CustomerID) || authorisation == ""{
			success = false
		}
	}
	if success {
		for _, deposit := range investment.Deposits {
			if testDatabase.MakeDeposit(db, investment.CustomerID, deposit) == false {success = false}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*");
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorisation");

	if success {
		fmt.Fprint(w, `{"success" : true}`)
	} else {
		fmt.Fprint(w, `{"success" : false}`)
	}	
}
