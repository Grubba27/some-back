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

func FindByEmail(email string) (User, error) {
	user := User{Email: email}
	db := db.GetDB()
	haveUser := db.First(&user, "email = ?", email)
	if errors.Is(haveUser.Error, gorm.ErrRecordNotFound) {
		return user, errors.New("User with that email was not found")
	}
	return user, nil
}

func FindByID(id uint) (User, error) {
	user := User{}
	db := db.GetDB()
	haveUser := db.First(&user, "id = ?", id)
	if errors.Is(haveUser.Error, gorm.ErrRecordNotFound) {
		return user, errors.New("User with that id was not found")
	}
	return user, nil
}

func UpdateUser(user User) (User, error) {
	db := db.GetDB()
	
	result := db.Save(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	return user, nil
}

func FindByPublicAddress(publicAddress string) (User, error) {
	user := User{}
	db := db.GetDB()
	haveUser := db.First(&user, "public_address = ?", publicAddress)
	if errors.Is(haveUser.Error, gorm.ErrRecordNotFound) {
		return user, errors.New("User with that public address was not found")
	}
	return user, nil
}