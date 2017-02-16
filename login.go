package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	session, err := cookieStore.Get(r, "mysqlProxy")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if vars["token"] == accessToken {
		session.Values["isLoggedIn"] = true
		session.Save(r, w)
		fmt.Fprintf(w, "%d", 1)
	} else {
		delete(session.Values, "isLoggedIn")
		session.Save(r, w)
		w.WriteHeader(http.StatusForbidden)
	}
}

func isLoggedIn(r *http.Request) bool {
	session, err := cookieStore.Get(r, "mysqlProxy")
	if err != nil {
		log.Printf("Session get error %s", err)
		return false
	}

	data, ok := session.Values["isLoggedIn"].(bool)
	if ok {
		return data
	} else {
		return false
	}
}
