package Migrations

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialect/postgres"
	"github.com/yceethetechie/green-bank-backend/Helpers"
)

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

// connect to db
func connectDB() *gorm.DB {
	// db, err := gorm.Open("postgres",os.Getenv("DB_HOST") + os.Getenv("PORT") + os.Getenv(DB_USER) + os.Getenv(DB_NAME) + os.Getenv(DB_PASSWORD) + os.Getenv(SSLMODE) + " " +)
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=greenbank password=greenbank sslmode=disabled")
	Helpers.HandleError(err)
	return db
}

// createAccount

func createAccount() {
	db := connectDB()
	users := [2]User{
		{Username: "Admin", Email: "admin@greenbank.com"},
		{Username: "Admin2", Email: "admin2@greenbank.com"},
	}
	for key, _ := range users {
		//generating passwords
		generatedPassword := Helpers.HashAndSaltPassword([]byte(users[key].Username))
		user := User{Username: users[key].Username,Email:users[key].Email,Password: generatedPassword} 
		db.Create(&user)
		//account details
		account := Account{Type:"Basic",Name:string(users[key].Username + "'s" + "account"), Balance: uint(1000 * int(key+1)), UserID: user.ID}
		db.Create(&account)
	}
	//closing db conn
	defer db.Close()
}

// migrate func
func Migrate(){
	db := connectDB()
	db.AutoMigrate(&User{},&Account{})	
	defer db.Close()
	createAccount()
}
