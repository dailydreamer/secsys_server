package libs

import (
	"log"
	"golang.org/x/crypto/bcrypt"
	"secsys/models"
)

// CreateAdmin create admin user
func CreateAdmin() {
	user := models.User{Phone:"1",Password:"1"}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error on encrypt password: " + err.Error())
		return
	}
	user.ID, err = models.CreateUser(user, string(hashedPassword), true)
	if err != nil {
		log.Println("Error on create user: " + err.Error())
		return
	}
	log.Println("Create admin user success!")
}