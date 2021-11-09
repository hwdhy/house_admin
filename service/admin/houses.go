package admin

import (
	"fmt"
	"gin_test/db"
	"gin_test/models"
	"gin_test/utils/basic"
	utilsToken "gin_test/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HousesInput struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	Status   string `json:"status"`
}

//用户房源列表
func Houses(c *gin.Context) {
	var input HousesInput
	if err := c.Bind(&input); err != nil {
		fmt.Println("bind err:", err)
		return
	}
	status := basic.StringToInt(input.Status)

	userId := utilsToken.GetContextUserId(c)
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	Offset := (input.Page - 1) * input.PageSize
	var Total int64
	var houses []models.House

	db.DB.Debug().Model(models.House{}).
		Where("status = ? and admin_id = ?", status, userId).
		Count(&Total).Offset(Offset).Limit(input.PageSize).Find(&houses)

	c.JSON(http.StatusOK, gin.H{
		"total": Total,
		"list":  houses,
	})
}
