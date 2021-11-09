package main

import (
	"crypto/md5"
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	result := md5.Sum([]byte("123456"))
	fmt.Println(fmt.Sprintf("%x", result))
}

func TestGet(t *testing.T) {

}