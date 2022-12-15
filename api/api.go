package api

import (
	"Maryjane_Roava_Assessment/interfaces"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

var DB *gorm.DB
var err error

func UserDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var currentAccountCustomers []interfaces.CurrentAccountCustomer
	DB.Find(&currentAccountCustomers)
	json.NewEncoder(w).Encode(currentAccountCustomers)
}

type Login struct {
	Username string
	Password string
}

func NewLogin(username string, password string) *Login {
	return &Login{Username: username, Password: password}
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
