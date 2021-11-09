package house

import (
	"gin_test/db"
	"gin_test/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type SearchAutocompleteInput struct {
	Prefix string `json:"prefix"`
}

//搜索补全
func SearchAutocomplete(c *gin.Context) {
	var input SearchAutocompleteInput
	err := c.Bind(&input)
	if err != nil {
		log.Println("bind err:", err)
		return
	}
	var titleList []string
	if input.Prefix != "" {
		db.DB.Debug().Model(models.House{}).Where("title like ?", "%"+input.Prefix+"%").Pluck("title", &titleList)
	}
	c.JSON(http.StatusOK, titleList)
}
