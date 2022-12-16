package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var err error

type ExistingUser struct {
	customerName  string `json:"customerName"`
	customerID    string `json:"customerID"`
	initialCredit int    `json:"initialCredit"`
}

type CurrentAccountUser struct {
	customerName  string `json:"customerName"`
	customerID    string `json:"customerID"`
	initialCredit int    `json:"initialCredit"`
}

func UserDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func CreateCurrentAccount(customer interface{}) CurrentAccountUser {
	var CurrentAccountUsers []CurrentAccountUser
	currentAccountUser := CurrentAccountUser(customer)
	CurrentUsers := append(CurrentAccountUsers, currentAccountUser)
	fmt.Println(CurrentAccountUsers, CurrentUsers)
	return currentAccountUser
}

type ErrResponse struct {
	Message string
}

func StartApi() {
	router := mux.NewRouter()
	router.HandleFunc("/user/{customerID}", UserDetails).Methods("GET")

	fmt.Println("App is working on port :8888")
	log.Fatal(http.ListenAndServe(":8888", router))

}
