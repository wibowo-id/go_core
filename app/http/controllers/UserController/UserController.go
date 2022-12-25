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
