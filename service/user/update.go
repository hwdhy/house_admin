package user

import (
	"fmt"
	"gin_test/db"
	"gin_test/models"
	utilsToken "gin_test/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdateInput struct {
	Avatar       string `json:"avatar"`
	Introduction string `json:"introduction"`
	NickName     string `json:"nickName"`
}

//用户信息更新
func Update(c *gin.Context) {
	var input UpdateInput
	err := c.Bind(&input)
	if err != nil {
		fmt.Println("bind err:", err)
		return
	}

	userId := utilsToken.GetContextUserId(c)
	if userId == 0 {
		c.JSON(http.StatusOK, nil)
		return
	}

	db.DB.Model(models.User{}).Where("id = ?", userId).Updates(
		map[string]interface{}{
			"avatar":       input.Avatar,
			"introduction": input.Introduction,
			"nick_name":    input.NickName,
		})
}
