package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	c "testgo11/controllers"
	m "testgo11/models"
	u "testgo11/utils"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func main() {
	port := os.Getenv("APP_PORT")
	u.SetAuthSecret("1GN1T3CH")
	u.SetNoAuth([]string{"/login"})

	router := mux.NewRouter()
	router.HandleFunc("/", handlerIndex)
	router.HandleFunc("/index", handlerIndex)
	router.HandleFunc("/hello", handlerHello)
	router.HandleFunc("/usr", listUsr).Methods("GET")
	router.HandleFunc("/login", c.LoginController).Methods("POST")

	router.Use(u.JwtAuthentication)
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
	usr := make([]*m.Usr, 0)
	db.Find(&usr)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usr)
}

func getDB() *gorm.DB {
	once.Do(func() {
		db = u.GetDB()
	})
	return db
}
