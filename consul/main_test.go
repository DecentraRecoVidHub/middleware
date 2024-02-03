package consul

import (
	"fmt"
	"os"
	"testing"
)

func TestConsul(t *testing.T) {
	//创建新的 consul sdk
	sdk := GetInstance()

	// 注册服务
	serviceID := "my_service"
	serviceName := "my_service"
	serviceHost := "127.0.0.1"
	servicePort := 8080

	err := sdk.Client.RegisterService(serviceID, serviceName, serviceHost, servicePort)
	if err != nil {
		fmt.Printf("Error registering service: %v\n", err)
		os.Exit(1)
	}

	// 发现服务
	serviceAddress, err := sdk.Client.DiscoverService(serviceName)
	if err != nil {
		fmt.Printf("Error discovering service: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Discovered service address: %s\n", serviceAddress)

	sdk.ConfigCenter.SetValue("test_key", "test_value")
	//使用 ConsulConfig 获取键对应的值
	value, err := sdk.ConfigCenter.GetValue("test_key")
	if err != nil {
		fmt.Println("Failed to get value:", err)
		return
	}
	fmt.Println("Value:", value)
}
