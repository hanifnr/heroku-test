package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	m "testgo11/models"
	u "testgo11/utils"

	"github.com/gorilla/mux"
)

var (
	db   *sql.DB
	once sync.Once
)

func main() {
	port := os.Getenv("APP_PORT")

	router := mux.NewRouter()
	router.HandleFunc("/", handlerIndex)
	router.HandleFunc("/index", handlerIndex)
	router.HandleFunc("/hello", handlerHello)
	router.HandleFunc("/usr", listUsr).Methods("GET")

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("listening to port %s", port)
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Welcome"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello world!"
	w.Write([]byte(message))
}

func listUsr(w http.ResponseWriter, r *http.Request) {
	db = getDB()
	rows, err := db.Query("SELECT * FROM usr")
	if err != nil {
		fmt.Printf("DB.Query: %v", err)
	}
	defer rows.Close()

	var usr []m.Usr
	for rows.Next() {
		var (
			id   int64
			name string
		)
		err := rows.Scan(&id, &name)
		if err != nil {
			fmt.Printf("Rows.Scan: %v", err)
		}
		usr = append(usr, m.Usr{Id: id, Name: name})
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usr)
}

func getDB() *sql.DB {
	once.Do(func() {
		db = u.GetDB()
	})
	return db
}
