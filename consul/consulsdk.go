package consul

import (
	"fmt"
	"os"
	"sync"
)

type ConsulSDK struct {
	Client       *ConsulClient
	ConfigCenter *ConsulConfigCenter
}

var (
	instance *ConsulSDK
	once     sync.Once
)

func NewConsulSDK(consulAddress string, serverPort int) (*ConsulSDK, error) {
	client, err := NewConsulClient(consulAddress, serverPort)
	if err != nil {
		return nil, err
	}

	configCenter, err := NewConsulConfigCenter(consulAddress)
	if err != nil {
		return nil, err
	}

	return &ConsulSDK{
		Client:       client,
		ConfigCenter: configCenter,
	}, nil
}

// GetInstance 给出一个单例展示
func GetInstance() *ConsulSDK {
	once.Do(func() {
		consulAddress := "127.0.0.1:8500"
		serverPort := 8080

		client, err := NewConsulClient(consulAddress, serverPort)
		if err != nil {
			fmt.Println("Failed to create Consul client:", err)
			os.Exit(1)
		}

		configCenter, err := NewConsulConfigCenter(consulAddress)
		if err != nil {
			fmt.Println("Failed to create Consul config center:", err)
			os.Exit(1)
		}

		instance = &ConsulSDK{
			Client:       client,
			ConfigCenter: configCenter,
		}
	})
	return instance
}
