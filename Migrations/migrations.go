package Migrations

import (
	"github.com/yceethetechie/green-bank-backend/Helpers"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialect/postgres"

)

type User struct {
	gorm.Model
	Username string
	Email string
	password string
}