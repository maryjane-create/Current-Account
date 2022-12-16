package api

import (
	"Maryjane_Roava_Assessment/users"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var err error

func UserDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func CreateCurrentAccount() {
	customers := users.ExistingCustomers
	customers = append(customers, users.ExistingCustomer{})
	for _ = range customers {
		{

		}

	}

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
