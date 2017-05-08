package libs

import (
	"log"
	"golang.org/x/crypto/bcrypt"
	"secsys/models"
)

// CreateAdmin create admin user
func CreateAdmin() {
	user := models.User{
		Phone:"admin",
		Password:"admin",
		CompanyBasic: models.CompanyBasic{
			ComName:"北京保安协会评定办",
		},
	}

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