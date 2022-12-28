package ProductController

import (
	"github.com/gin-gonic/gin"
	prod "github.com/wibowo-id/sms-backend/app/models"
	"github.com/wibowo-id/sms-backend/app/models/entity"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func Index(c *gin.Context) {
	var products []entity.Product

	prod.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"datas": products})
}

func Show(c *gin.Context) {
	var product entity.Product
	id := c.Param("id")

	if err := prod.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

func Create(c *gin.Context) {
	var product entity.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	prod.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"data": product})
}

func Update(c *gin.Context) {
	var product entity.Product
	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if prod.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Gagal membperbarui produk"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Berhasil memperbarui data"})
}

func Delete(c *gin.Context) {
	var product entity.Product
	input := map[string]string{"id": "0"}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := strconv.ParseInt(input["id"], 10, 64)
	if prod.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus produk"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Berhasil dihapus"})
}

func Store(c *gin.Context) {

}

func Edit(c *gin.Context) {

}
