package Migrations

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
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

func init() {
	// loads values from .env into the system
	err := godotenv.Load()
	Helpers.HandleError(err)
}

// connect to db
func connectDB() *gorm.DB {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("PORT")
	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password) //Build connection strin
	conn, err := gorm.Open("postgres", dbUri)
	// db, err := gorm.Open("postgres",os.Getenv("DB_HOST") + os.Getenv("PORT") + os.Getenv(DB_USER) + os.Getenv(DB_NAME) + os.Getenv(DB_PASSWORD) + os.Getenv(SSLMODE) + " " +)
	// db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=greenbank password=greenbank sslmode=disable")
	Helpers.HandleError(err)
	return conn
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
		user := User{Username: users[key].Username, Email: users[key].Email, Password: generatedPassword}
		db.Create(&user)
		//account details
		account := Account{Type: "Basic", Name: string(users[key].Username + "'s" + "account"), Balance: uint(1000 * int(key+1)), UserID: user.ID}
		db.Create(&account)
	}
	//closing db conn
	defer db.Close()
}

// migrate func
func Migrate() {
	db := connectDB()
	db.AutoMigrate(&User{}, &Account{})
	defer db.Close()
	createAccount()
}
