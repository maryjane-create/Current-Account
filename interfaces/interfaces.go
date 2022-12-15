package interfaces

import "github.com/jinzhu/gorm"

type CurrentAccountCustomer struct {
	gorm.Model
	CustomerID    string `json:"customerID"`
	initialCredit int    `json:"initialCredit"`
}
type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

type Account struct {
	gorm.Model
	Type    string
	Name    string
	Balance uint
	UserID  uint
}

type ResponseAccount struct {
	ID      uint
	Name    string
	Balance int
}

type ResponseUser struct {
	ID       uint
	Username string
	Email    string
	Accounts []ResponseAccount
}
