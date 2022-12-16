package users

type ExistingCustomer struct {
	customerName  string `json:"customerName"`
	customerID    string `json:"customerID"`
	initialCredit int    `json:"initialCredit"`
}
type CurrentAccountUser struct {
	customerName string `json:"customerName"`
	newAccountID string `json:"newAccountID"`
	newBalance   int    `json:"newBalance"`
}

var existingCustomers []ExistingCustomer
var CurrentAccountUsers []CurrentAccountUser
