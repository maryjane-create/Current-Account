package users

type ExistingCustomer struct {
	customerID    string `json:"customerID"`
	initialCredit int    `json:"initialCredit"`
}
type CurrentAccountUser struct {
	newAccountID string `json:"newAccountID"`
	newBalance   int    `json:"newBalance"`
}
