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

var ExistingCustomers []ExistingUser
var CurrentAccountUsers []users.CurrentAccountUser

func TestExistingUsers(t *testing.T) {
	user1 := ExistingUser{"bolu", "io", 9}
	user2 := ExistingUser{"tope", "iu", 0}
	ListOfExistingCustomers := append(append(ExistingCustomers, user1), user2)
	assert.Equal(t, ListOfExistingCustomers, user1)
}

func TestCurrentAccountUsers(t *testing.T) {

	users1 := ExistingUser{"ruth", "ui", 9}
	users2 := ExistingUser{"bolu", "io", 8}
	ListOfExistingCustomers := append(append(ExistingCustomers, users1), users2)
	api.CreateCurrentAccount(users1)
	assert.Equal(t, CurrentAccountUsers, ListOfExistingCustomers)

}
