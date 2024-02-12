package txcos

import (
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
	"net/http"
	"net/url"
	"sync"
)

// 声明一个包级别的私有变量，存储单例实例
var cosClient *cos.Client

// 使用 sync.Once 确保实例化操作只执行一次
var once sync.Once

// GetCosClient 返回单例对象的实例，公开访问点
func getCosClient() *cos.Client {
	once.Do(func() {
		// 访问域名
		u, _ := url.Parse("https://dvideo-1302745585.cos.ap-singapore.myqcloud.com")
		b := &cos.BaseURL{BucketURL: u}
		cosClient = cos.NewClient(b, &http.Client{
			Transport: &cos.AuthorizationTransport{
				// 通过环境变量获取密钥
				SecretID:  COS_SECRETID,
				SecretKey: COS_SECRETKEY,
				// Debug 模式，把对应 请求头部、请求内容、响应头部、响应内容 输出到标准输出
				Transport: &debug.DebugRequestTransport{
					RequestHeader: false,
					// Notice when put a large file and set need the request body, might happend out of memory error.
					RequestBody:    false,
					ResponseHeader: false,
					ResponseBody:   false,
				},
			},
		})
	})
	return cosClient
}
