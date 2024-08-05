package utils

import (
	"github.com/goccy/go-json"
	"github.com/gomodule/redigo/redis"
)

const (
	RedisLocationName = "RedisLocationName"
)

// GetIP4Location 获取IP4位置信息
func GetIP4Location(rdsConn redis.Conn, ip4 string) string {
	address, _ := redis.String(rdsConn.Do("HGET", RedisLocationName, ip4))
	if address == "" {
		respData := new(struct {
			Code int               `json:"code"`
			Msg  string            `json:"msg"`
			Data map[string]string `json:"data"`
		})
		dataBytes, _ := NewClient().Request("GET", "https://searchplugin.csdn.net/api/v1/ip/get?ip="+ip4, nil)
		_ = json.Unmarshal(dataBytes, &respData)
		if _, ok := respData.Data["address"]; ok {
			address = respData.Data["address"]
			_, _ = rdsConn.Do("HSET", RedisLocationName, ip4, address)
			return address
		}
		return ""
	}
	return address
}
