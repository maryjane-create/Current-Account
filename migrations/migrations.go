package migrations

import (
	"Maryjane_Roava_Assessment/helpers"
	"Maryjane_Roava_Assessment/interfaces"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=bankapp password=postgres sslmode=disable")
	HandleErr(err)
	return db
}

func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func CreateCurrentAccount(customerID string, initialCredit uint) string {
	db, _ := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=bankapp password=postgres sslmode=disable")
	newCustomer := CreateCurrentAccount(customerID, initialCredit)
	if db.Model(&newCustomer).Where("customerID=?", customerID).Updates(&newCustomer).RowsAffected == 0 {
		db.Create(&newCustomer)
		return customerID
	}
	return "error"

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
