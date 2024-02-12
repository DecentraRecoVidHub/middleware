package txcos

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
)

// 为任意长度，任意类型的变量进行hash，并返回hash后的字符串
func hashVariables(vars ...interface{}) (string, error) {
	// 初始化一个buffer用于存储所有变量的字节序列
	var buffer bytes.Buffer

	// 创建一个gob编码器
	enc := gob.NewEncoder(&buffer)

	for _, v := range vars {
		// 使用gob编码器将变量编码到buffer
		if err := enc.Encode(v); err != nil {
			return "", err
		}
	}

	// 对编码后的字节序列进行SHA-256哈希处理
	hash := sha256.New()
	hash.Write(buffer.Bytes())
	hashed := hash.Sum(nil)

	// 将哈希值转换为十六进制字符串
	return hex.EncodeToString(hashed), nil
}

func log_status(err error) {
	if err == nil {
		return
	}
	if cos.IsNotFoundError(err) {
		// WARN
		fmt.Println("WARN: Resource is not existed")
	} else if e, ok := cos.IsCOSError(err); ok {
		fmt.Printf("ERROR: Code: %v\n", e.Code)
		fmt.Printf("ERROR: Message: %v\n", e.Message)
		fmt.Printf("ERROR: Resource: %v\n", e.Resource)
		fmt.Printf("ERROR: RequestId: %v\n", e.RequestID)
		// ERROR
	} else {
		fmt.Printf("ERROR: %v\n", err)
		// ERROR
	}
}
