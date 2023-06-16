package user

import (
	"app/db"
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email         string `json:"email" binding:"required,email" gorm:"uniqueIndex"`
	Password      string `json:"password" binding:"required"`
	PublicAddress string `json:"publicAddress"`
}

// Create user in database
// If user already exists, return error and user
func Create(email string, password string) (User, error) {
	// TODO: Hash password

	user := User{Email: email, Password: password}
	db := db.GetDB()
	
	haveUser := db.First(&user, "email = ?", email)
	if haveUser.RowsAffected != 0 {
		return user, errors.New("User with that email already exists")
	}

	result := db.Create(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	return user, nil
}

func Authenticate(email string, password string) (User, error) {
	user := User{Email: email, Password: password}
	db := db.GetDB()
	haveUser := db.First(&user, "email = ?", email)
	if errors.Is(haveUser.Error, gorm.ErrRecordNotFound) {
		return user, errors.New("User with that email was not found")
	}
	// jwt
	return user, nil
}
