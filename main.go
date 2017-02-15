package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
	"time"
)

var db *sql.DB
var accessToken string

func main() {
	mysqlString, accessToken1, sessionKey, port := checkFlags()
	accessToken = accessToken1

	log.Println(sessionKey)

	db1, err := sql.Open("mysql", mysqlString)
	checkErr(err)
	db = db1

	r := mux.NewRouter()
	r.HandleFunc("/query/{qr}", queryHandler).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:" + string(port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("mysqlproxy started")
	log.Fatal(srv.ListenAndServe())
}

func checkFlags() (string, string, string, string) {
	mysqlFlag := flag.String("mysqlFlag", "", "mysqlConnectionString -mysqlFlag=user:pass@/dbname")
	accessToken := flag.String("accessToken", "", "access token for client authorization")
	sessionKey := flag.String("sessionKey", "very secret", "session store encryption key")
	port := flag.String("port", "7000", "application listen port")

	flag.Parse()

	if *mysqlFlag == "" {
		log.Fatal("Need mysql connection string flag. For example -mysqlFlag=user:pass@/dbname")
	}

	if *accessToken == "" {
		log.Fatal("Need accessToken flag. For example -accessToken=supersecret")
	}

	return *mysqlFlag, *accessToken, *sessionKey, *port
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
	if err != nil {
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
