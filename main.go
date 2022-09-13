package models

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	db := GetDB()
	usr := make([]*Usr, 0)
	db.Find(&usr)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usr)
}
