package UserController

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/wibowo-id/go_core/app/models/entity"
	"github.com/wibowo-id/go_core/config"
	"github.com/wibowo-id/go_core/utils/token"
)

func CurrentUser(c *gin.Context) {
	userId, err := token.ExtractTokenID(c)
	Session := config.Session{}

	session := sessions.Default(c)
	fmt.Println("session :", session.Get(Session.Email))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := entity.GetUserByID(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

func Profile(c *gin.Context) {
	Session := config.Session{}
	session := sessions.Default(c)
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": session.Get(Session.Email)})
}

func Index(c *gin.Context) {

}

func Show(c *gin.Context) {

}

func Store(c *gin.Context) {

}

func Edit(c *gin.Context) {

}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}

func UpdateStatus(c *gin.Context) {

}

func Import(c *gin.Context) {

}

func ExportCsv(c *gin.Context) {

}

func SaveProfile(c *gin.Context) {

}
