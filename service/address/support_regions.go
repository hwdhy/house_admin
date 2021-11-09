package address

import (
	"gin_test/db"
	"gin_test/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SupportRegionsList 通过支持地区的英语简称查询 belongTo  和level
func SupportRegionsList(c *gin.Context) {
	enName := c.Param("enName")

	var regions []models.SupportAddress
	db.DB.Model(models.SupportAddress{}).Where("belong_to = ? and level = ?", enName, "region").Find(&regions)

	c.JSON(http.StatusOK, gin.H{
		"list": regions,
	})
}
