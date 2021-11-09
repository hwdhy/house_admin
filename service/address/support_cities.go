package address

import (
	"gin_test/db"
	"gin_test/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SupportCities 支持的区域列表
func SupportCities(c *gin.Context) {

	var addressData []models.SupportAddress
	db.DB.Model(models.SupportAddress{}).Where("level = ?", "city").Find(&addressData)

	c.JSON(http.StatusOK, gin.H{
		"list": addressData,
	})
}
