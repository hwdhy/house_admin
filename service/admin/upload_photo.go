package admin

import (
	"gin_test/conf"
	utilsToken "gin_test/utils/token"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

type UploadPhotoOutput struct {
	Size      int64  `json:"size"`      //图片大小
	Key       string `json:"key"`       //图片名称
	Url       string `json:"url"`       //图片URL
	CdnPrefix string `json:"cdnPrefix"` //图片CDN
}

//上传图片
func UploadPhoto(c *gin.Context) {
	userId := utilsToken.GetContextUserId(c)
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, "用户登录失效")
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		log.Println("formFile err:", err)
		return
	}
	size := file.Size
	uuidStr := strings.ReplaceAll(uuid.NewString(), "-", "")
	ext := filepath.Ext(file.Filename)
	filePath := conf.UploadFilePath + uuidStr + ext

	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		log.Println("upload File err:", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, UploadPhotoOutput{
		Size:      size,
		Key:       uuidStr + ext,
		Url:       conf.CdnPath + uuidStr + ext,
		CdnPrefix: conf.CdnPath,
	})
}
