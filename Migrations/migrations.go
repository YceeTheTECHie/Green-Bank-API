package Migrations

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialect/postgres"
	"github.com/yceethetechie/green-bank-backend/Helpers"
	"golang.org/x/tools/go/analysis/unitchecker"
)

type User struct {
	gorm.Model
	Username string
	Email string
	password string
}


type Account struct{
	gorm.Model
	Type string
	Name string
	Balance uint
	UserID uint
}