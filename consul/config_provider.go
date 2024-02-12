package consul

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
)

type ConsulConfigCenter struct {
	client *consulapi.Client
}

func NewConsulConfigCenter(consulAddress string) (*ConsulConfigCenter, error) {
	config := consulapi.DefaultConfig()
	config.Address = consulAddress
	client, err := consulapi.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &ConsulConfigCenter{client: client}, nil
}

// GetValue 获取特定键对应的值
func (cc *ConsulConfigCenter) GetValue(key string) (string, error) {
	kv := cc.client.KV()
	pair, _, err := kv.Get(key, nil)
	if err != nil {
		return "", err
	}
	if pair == nil {
		return "", fmt.Errorf("key '%s' not found", key)
	}
	return string(pair.Value), nil
}

// SetValue 键值存储中设置一个键值对
func (cc *ConsulConfigCenter) SetValue(key, value string) error {
	kv := cc.client.KV()
	p := &consulapi.KVPair{Key: key, Value: []byte(value)}
	_, err := kv.Put(p, nil)
	return err
}

// DeleteValue 删除键值对内容
func (cc *ConsulConfigCenter) DeleteValue(key string) error {
	kv := cc.client.KV()
	_, err := kv.Delete(key, nil)
	return err
}
