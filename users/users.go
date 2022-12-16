package users

type ExistingCustomer struct {
	customerName  string `json:"customerName"`
	customerID    string `json:"customerID"`
	initialCredit int    `json:"initialCredit"`
}
type CurrentAccountUser struct {
	customerName  string `json:"customerName"`
	customerID    string `json:"customerID"`
	initialCredit int    `json:"initialCredit"`
}

var ExistingCustomers []ExistingCustomer
var CurrentAccountUsers []CurrentAccountUser
