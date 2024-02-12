package txcos

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

// 定义UploadType类型
type UploadType string

// 定义UploadType的可能值
const (
	VIDEO     UploadType = "video"
	THUMBNAIL UploadType = "thumbnail"
)

// Upload 上传视频或者缩略图的接口
func Upload(uploadType UploadType, Bytes []byte, objName string, uID int) (string, error) {
	// 生成这个obj唯一的id
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	videoHash, _ := hashVariables(Bytes, uID, objName, timestamp)
	objectKey := string(uploadType) + "/" + videoHash
	// 上传对象到腾讯云
	err := PutObject(objectKey, Bytes)
	if err != nil {
		return "", err
	}
	return objectKey, nil
}

// PutObject 添加一个obj的接口
func PutObject(objectKey string, bs []byte) error {
	client := getCosClient()
	// Convert byte slice to io.Reader
	f := bytes.NewReader(bs)
	_, err := client.Object.Put(context.Background(), objectKey, f, nil)
	log_status(err)
	return err
}

// GetObject 获取一个obj的接口
func GetObject(objectKey string) ([]byte, error) {
	client := getCosClient()
	resp, err := client.Object.Get(context.Background(), objectKey, nil)
	if err != nil {
		log_status(err)
		return nil, err
	}
	bs, _ := io.ReadAll(resp.Body)
	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return bs, nil
}

// GetObjectAccessURL 获取对象的访问link
func GetObjectAccessURL(objectKey string) (string, error) {
	client := getCosClient()
	// 生成预签名URL
	// HTTP 方法, 存储桶名称, 对象键, 过期时间
	presignedURL, err := client.Object.GetPresignedURL(
		context.Background(),
		http.MethodGet, objectKey,
		COS_SECRETID,
		COS_SECRETKEY,
		time.Hour*24,
		nil)
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL%v", err)
	}
	return presignedURL.String(), nil
}

func DeleteObject(objectKey string) error {
	client := getCosClient()
	_, err := client.Object.Delete(context.Background(), objectKey, nil)
	log_status(err)
	return err
}
