package user

import (
	"crypto/sha1"
	"fmt"
	"gin_test/conf"
	"gin_test/db"
	"gin_test/models"
	utilsFile "gin_test/utils/file"
	utilsToken "gin_test/utils/token"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

type UploadPhotoOutput struct {
	Size      int64  `json:"size"`      //图片大小
	Key       string `json:"key"`       //图片名称
	Url       string `json:"url"`       //图片URL
	CdnPrefix string `json:"cdnPrefix"` //图片CDN
}

//用户上传图片
func UploadPhoto(c *gin.Context) {
	userId := utilsToken.GetContextUserId(c)
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, "用户登录失效")
		return
	}

	//获取表单文件流
	files, err := c.FormFile("file")
	if err != nil {
		log.Println("formFile err:", err)
		return
	}
	size := files.Size
	ext := filepath.Ext(files.Filename)

	//打开表单文件流  获取file类型文件
	file, err := files.Open()
	if err != nil {
		return
	}

	//读取文件流
	all, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		return
	}

	//文件在数据库是否存在是否存在
	hash := sha1.New()
	hash.Write(all)
	fileHashName := fmt.Sprintf("%x", hash.Sum(nil))

	//判断hashName是否存在
	var picture models.HousePicture
	db.DB.Model(models.HousePicture{}).Where("location = ?", fileHashName).First(&picture)
	if picture.ID != 0 {
		c.JSON(http.StatusOK, UploadPhotoOutput{
			Size:      size,                              //图片大小
			Key:       fileHashName + ext,                //文件名 + 文件后缀
			Url:       conf.CdnPath + fileHashName + ext, //文件全路径
			CdnPrefix: conf.CdnPath,                      //CDN
		})
	}
	//判断路径是否存在
	utilsFile.FileCreatePath(conf.UploadFilePath)

	fileName := conf.UploadFilePath + fileHashName + ext

	//文件上传
	err = c.SaveUploadedFile(files, fileName)
	if err != nil {
		log.Println("upload File err:", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	_, _ = utilsFile.ImagesResize(fileName, []string{"370x270", "140x90"})

	c.JSON(http.StatusOK, UploadPhotoOutput{
		Size:      size,                              //图片大小
		Key:       fileHashName + ext,                //文件名 + 文件后缀
		Url:       conf.CdnPath + fileHashName + ext, //文件全路径
		CdnPrefix: conf.CdnPath,                      //CDN
	})
}
