package main

import (
	"Maryjane_Roava_Assessment/api"
	"Maryjane_Roava_Assessment/users"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ExistingUser struct {
	name          string
	customerID    string
	initialCredit int
}

type CurrentAccountUser struct {
	customerName  string `json:"customerName"`
	customerID    string `json:"customerID"`
	initialCredit int    `json:"initialCredit"`
}

var ExistingCustomers []api.ExistingUser
var CurrentAccountUsers []users.CurrentAccountUser

func TestExistingUsers(t *testing.T) {
	user1 := api.ExistingUser(ExistingUser{"y", "t", 7})
	user2 := api.ExistingUser(ExistingUser{"tope", "iu", 0})
	ListOfExistingCustomers := append(append(ExistingCustomers, user1), user2)
	assert.Equal(t, ListOfExistingCustomers, user1)
}

func TestCurrentAccountUsers(t *testing.T) {

	users1 := api.ExistingUser{"ruth", "ui", 9}
	users2 := api.ExistingUser{"bolu", "io", 8}
	ListOfExistingCustomers := append(append(ExistingCustomers, users1), users2)
	user1 := CurrentAccountUser(users1)
	api.CreateCurrentAccount(users1)
	assert.Equal(t, CurrentAccountUsers, ListOfExistingCustomers)

}
