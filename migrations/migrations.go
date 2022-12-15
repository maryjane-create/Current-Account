package migrations

import (
	"Maryjane_Roava_Assessment/helpers"
	"Maryjane_Roava_Assessment/interfaces"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateCurrentAccount(customerID string, initialCredit uint) (string, uint, int) {

	db := connectDB()
	currentAccountUser := &[2]interfaces.User{}
	for i := 0; i < len(currentAccountUser); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(currentAccountUser[i].Username))
		currentUser := &interfaces.User{Username: currentAccountUser[i].Username, Email: currentAccountUser[i].Email, Password: generatedPassword}
		db.Create(&currentUser)
	}
	return customerID, initialCredit, len(currentAccountUser)
}

func connectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=bankapp password=postgres sslmode=disable")
	helpers.HandleErr(err)
	return db
}

func createAccounts() {
	db := connectDB()

	users := &[2]interfaces.User{}

	for i := 0; i < len(users); i++ {
		// Correct one way
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := &interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)

		account := &interfaces.Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + " account"), Balance: uint(10000 * int(i+1)), UserID: user.ID}
		db.Create(&account)
	}
	defer db.Close()
}

func Migrate() {
	User := &interfaces.User{}
	Account := &interfaces.Account{}
	db := connectDB()
	db.AutoMigrate(&User, &Account)
	defer db.Close()

	createAccounts()
}
