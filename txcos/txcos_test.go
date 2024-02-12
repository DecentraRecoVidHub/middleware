package txcos

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestHash(t *testing.T) {
	videoBytes := []byte{1, 2, 5, 4, 8}
	fmt.Println(videoBytes)
	uID := 2
	videoName := "testname"
	// 获取当前时间戳，并转换为字节切片
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	hash, err := hashVariables(videoBytes, uID, videoName, timestamp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hash)
}

func TestUploadVideo(t *testing.T) {
	videoBytes := []byte{1, 2, 5, 4, 8}
	uID := 2
	videoName := "testname"
	videoObjectKey, err := Upload(VIDEO, videoBytes, videoName, uID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(videoObjectKey) // "video/e1f0d6e4d0a2b824f28cc1e48feeabc0f9b8ce9c0d0729a3c6dbaf4637d9b816"
}

// 这个速度过慢
func TestGetObject(t *testing.T) {
	objkey := "video/e1f0d6e4d0a2b824f28cc1e48feeabc0f9b8ce9c0d0729a3c6dbaf4637d9b816"
	bytes, err := GetObject(objkey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bytes)
}

func TestGetObjectAccessURL(t *testing.T) {
	objkey := "video/e1f0d6e4d0a2b824f28cc1e48feeabc0f9b8ce9c0d0729a3c6dbaf4637d9b816"

	url, err := GetObjectAccessURL(objkey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(url)
}

func TestDeleteObject(t *testing.T) {
	objkey := "video/e1f0d6e4d0a2b824f28cc1e48feeabc0f9b8ce9c0d0729a3c6dbaf4637d9b816"
	err := DeleteObject(objkey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success delete")
}
