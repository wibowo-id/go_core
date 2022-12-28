package LoginController

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/wibowo-id/go_core/app/models"
	"github.com/wibowo-id/go_core/app/models/entity"
	"github.com/wibowo-id/go_core/config"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateToken struct {
	Token string `json:"token"`
}

type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required" gorm:"unique;not null"`
}

func Login(c *gin.Context) {
	var input LoginInput
	var user entity.User
	var SetSessions config.Session

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := entity.User{}
	u.Email = input.Email
	u.Password = input.Password

	token, err := entity.LoginCheck(u.Email, u.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email or password is incorrect."})
		return
	}

	if models.DB.Model(&user).Where("email = ?", u.Email).Update("token", token).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Gagal membperbarui token"})
		return
	}

	getUser, _ := entity.GetUserByEmail(u.Email)
	fmt.Println("get User :", getUser)

	session := sessions.Default(c)
	session.Set(SetSessions.Id, getUser.Id)
	session.Set(SetSessions.Name, getUser.Name)
	session.Set(SetSessions.Token, getUser.Token)
	session.Set(SetSessions.Email, getUser.Email)
	err = session.Save()
	if err != nil {
		return
	}

	UserLoggedIn := entity.UserLoggedIn{}
	UserLoggedIn.Id = getUser.Id
	UserLoggedIn.Name = getUser.Name
	UserLoggedIn.Email = getUser.Email
	UserLoggedIn.Token = getUser.Token

	c.JSON(http.StatusOK, UserLoggedIn)
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := entity.User{}
	u.Name = input.Name
	u.Email = input.Email
	u.Password = input.Password

	_, err := u.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "validated"})
}

func Logout(c *gin.Context) {
	//session, _ := config.Config()
	//var user models.User
	//
	//if models.DB.Model(&user).Where("email = ?", u.Email).Update("token", "").RowsAffected == 0 {
	//	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Gagal membperbarui token"})
	//	return
	//}
}
