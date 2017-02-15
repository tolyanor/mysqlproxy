package main

import (
  "fmt"
  "os"
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"strings"
)

var db *sql.DB

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Need mysql connection string in pararm. For example user:pass@/dbname")
	}
  db1, err := sql.Open("mysql", os.Args[1])
  checkErr(err)
  db = db1

	r := mux.NewRouter()
	r.HandleFunc("/query/{qr}", queryHandler).Methods("GET")	

	srv := &http.Server{
    Handler:      r,
    Addr:         "127.0.0.1:7000",
    WriteTimeout: 15 * time.Second,
    ReadTimeout:  15 * time.Second,
  }

  log.Printf("mysqlproxy started")
  log.Fatal(srv.ListenAndServe())
}


func queryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	qr := vars["qr"]

	var res string
	var err error

	if strings.Index(qr, "select") == 0 {
  	res, err = query(qr, db)
	} else {
		res, err = exec(qr, db)
	}
	if (err != nil) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err)	
	  log.Printf("query %s %s", qr, err)
	} else {
		fmt.Fprintf(w, "%s", res)	
	  log.Printf("query %s", qr)
	}
}

func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}

