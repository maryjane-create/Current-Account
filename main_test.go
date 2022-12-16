package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ExistingUser struct {
	name          string
	customerID    string
	initialCredit int
}

var ExistingCustomers []ExistingUser

func TestUsers(t *testing.T) {
	users1 := ExistingUser{"ruth", "ui", 9}
	newList := append(ExistingCustomers, users1)

	assert.Equal(t, users1, newList)

}
