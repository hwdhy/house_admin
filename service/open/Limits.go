package open

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LimitOutput struct {
	UserPasswordRegex   string   `json:"userPasswordRegex"`
	PhoneRegex          string   `json:"phoneRegex"`
	AvatarSizeLimit     uint     `json:"avatarSizeLimit"`
	AvatarTypeLimit     []string `json:"avatarTypeLimit"`
	HousePhotoSizeLimit uint     `json:"housePhotoSizeLimit"`
	HousePhotoTypeLimit []string `json:"housePhotoTypeLimit"`
}

// Limits 获取规则
func Limits(c *gin.Context) {
	c.JSON(http.StatusOK, LimitOutput{
		UserPasswordRegex:   "^(?=.*\\d)((?=.*[a-z])|(?=.*[A-Z])).{8,16}$",
		PhoneRegex:          "^(1[3-9]\\d{9}$)",
		AvatarSizeLimit:     5242880,
		AvatarTypeLimit:     []string{"jpg", "png", "jpeg"},
		HousePhotoSizeLimit: 10485760,
		HousePhotoTypeLimit: []string{"jpg", "png", "jpeg", "webp"},
	})
}
