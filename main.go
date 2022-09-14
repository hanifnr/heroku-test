package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

var (
	db   *sql.DB
	once sync.Once
)

func main() {

	port := "8080"

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)
	http.HandleFunc("/usr", listUsr)

	err := http.ListenAndServe(":"+port, nil)
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
	rows, err := db.Query("SELECT * FROM votes ORDER BY created_at DESC LIMIT 5")
	if err != nil {
		fmt.Errorf("DB.Query: %v", err)
	}
	defer rows.Close()

	var usr []Usr
	for rows.Next() {
		var (
			id   int64
			name string
		)
		err := rows.Scan(&id, &name)
		if err != nil {
			fmt.Errorf("Rows.Scan: %v", err)
		}
		usr = append(usr, Usr{Id: id, Name: name})
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usr)
}

func getDB() *sql.DB {
	once.Do(func() {
		db = GetDB()
	})
	return db
}
