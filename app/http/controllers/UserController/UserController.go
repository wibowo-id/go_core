package UserController

import (
	"github.com/gin-gonic/gin"
	"github.com/wibowo-id/sms-backend/app/models"
	"github.com/wibowo-id/sms-backend/utils/token"
	"net/http"
)

func CurrentUser(c *gin.Context) {
	userId, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUserByID(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
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

func Profile(c *gin.Context) {

}

func SaveProfile(c *gin.Context) {

}
