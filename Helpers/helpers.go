package Helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HandleError (err error){
	if err != nil {
		panic(err.Error())
	}
}

func HashAndSaltPassword(password []byte) string {
	 hashed,err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost) 
	 HandleError(err)
	 return string(hashed)
	 
}