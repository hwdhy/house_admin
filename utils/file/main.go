package utilsFile

import (
	"github.com/nfnt/resize"
	"golang.org/x/image/bmp"
	"golang.org/x/image/webp"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// FileCreatePath 目录不存在创建目录
func FileCreatePath(filePath string) {
	_, err := os.Stat(filePath)
	if err != nil {
		if !os.IsExist(err) {
			err = os.MkdirAll(filePath, os.ModePerm)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

// ImagesResize 图片裁剪
func ImagesResize(fileName string, size []string) ([]string, error) {

	var newImageList []string

	//打开文件流
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var img image.Image
	//获取文件类型
	fileType, err := FileType(fileName)
	if err != nil {
		return nil, err
	}

	switch fileType {
	case "image/png":
		img, err = png.Decode(file)
	case "image/webp":
		img, err = webp.Decode(file)
	case "bmp":
		img, err = bmp.Decode(file)
	default:
		img, err = jpeg.Decode(file)
	}

	if err != nil {
		return nil, err
	}

	for k := range size {
		var width, height uint

		switch size[k] {
		case "default":
			width = uint(img.Bounds().Dx())
			height = uint(img.Bounds().Dy())
		default:
			sizeArr := strings.Split(size[k], "x")
			widthA, _ := strconv.Atoi(sizeArr[0])
			heightA, _ := strconv.Atoi(sizeArr[1])
			width = uint(widthA)
			height = uint(heightA)
		}

		newImage := strings.Replace(fileName, ".", "_"+size[k]+".", -1)

		m := resize.Resize(width, height, img, resize.Lanczos3)

		newImageFile, err := os.Create(newImage)
		if err != nil {
			return nil, err
		}
		defer newImageFile.Close()

		err = jpeg.Encode(newImageFile, m, nil)
		if err != nil {
			return nil, err
		}
		newImageList = append(newImageList, newImage)
	}

	return newImageList, nil
}

// FileType 判断图片类型
func FileType(fileName string) (string, error) {

	open, err := os.Open(fileName)
	if err != nil {
		return "", nil
	}
	defer open.Close()

	//判断文件类型
	buffer := make([]byte, 512)
	_, err = open.Read(buffer)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
