package entity

import (
	"errors"
	"fmt"
	"html"
	"strings"

	"github.com/wibowo-id/go_core/app/models"
	"github.com/wibowo-id/go_core/utils/token"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id       uint   `gorm:"primarykey:auto_increment" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Email    string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string `gorm:"->;<-;not null" json:"-"`
	Token    string `gorm:"type:text" json:"-"`
}

type UserLoggedIn struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func (u *User) Save() (*User, error) {
	var err error
	err = models.DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}

	u.Name = html.EscapeString(strings.TrimSpace(u.Name))

	return
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(email, password string) (string, error) {
	var err error

	u := User{}
	err = models.DB.Model(User{}).Where("email = ?", email).Take(&u).Error
	fmt.Println("err : ", email)
	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	generateToken, err := token.GenerateToken(u.Id)
	if err != nil {
		return "", err
	}

	return generateToken, nil

}

func GetUserByID(uid uint) (User, error) {

	var u User

	if err := models.DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}

	u.PrepareGive()

	return u, nil

}

func GetUserByEmail(email string) (User, error) {

	var u User

	if err := models.DB.Where("email = ?", email).First(&u).Error; err != nil {
		return u, errors.New("User not found!")
	}

	u.PrepareGive()

	return u, nil

}

func (u *User) PrepareGive() {
	u.Password = ""
}
