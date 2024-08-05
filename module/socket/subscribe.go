package socket

import (
	"github.com/goccy/go-json"
	"github.com/gomodule/redigo/redis"
	"gofiber/utils"
)

const (
	RedisSubscribeChannelList = "RedisSubscribeChannelList"
)

// SubscribeInfo 订阅信息
type SubscribeInfo struct {
	Channel string   `json:"channel"` //	订阅通道
	Args    []string `json:"args"`    //	订阅标识组
}

// Subscribe 订阅
func (_Socket *Socket) Subscribe(rdsConn redis.Conn, uuidStr string, channel string, args []string) {
	subscribeInfo := _Socket.GetRedisSubscribeInfo(channel)
	if subscribeInfo != nil {
		subChannelList := _Socket.GetSubscribeChannelList(rdsConn, channel)

		if _, ok := subChannelList[uuidStr]; !ok {
			subChannelList[uuidStr] = make([]string, 0)
		}

		for _, arg := range args {
			if utils.ArrayStringIndexOf(subChannelList[uuidStr], arg) == -1 && utils.ArrayStringIndexOf(subscribeInfo.Args, arg) > -1 {
				subChannelList[uuidStr] = append(subChannelList[uuidStr], arg)
			}
		}

		// 更新订阅通道数据
		channelListBytes, _ := json.Marshal(subChannelList)
		_, _ = rdsConn.Do("HSET", RedisSubscribeChannelList+_Socket.key, channel, channelListBytes)
	}
}

// UnSubscribe 取消订阅
func (_Socket *Socket) UnSubscribe(rdsConn redis.Conn, uuidStr string, channel string, args []string) {
	subChannelList := _Socket.GetSubscribeChannelList(rdsConn, channel)

	if _, ok := subChannelList[uuidStr]; ok {
		currentArgs := make([]string, 0)
		for _, arg := range subChannelList[uuidStr] {
			if utils.ArrayStringIndexOf(args, arg) == -1 {
				currentArgs = append(currentArgs, arg)
			}
		}

		// 如果没有数据, 那么删除当前对象
		if len(currentArgs) == 0 {
			delete(subChannelList, uuidStr)
		} else {
			subChannelList[uuidStr] = currentArgs
		}

		// 更新订阅通道数据
		channelListBytes, _ := json.Marshal(subChannelList)
		_, _ = rdsConn.Do("HSET", RedisSubscribeChannelList+_Socket.key, channel, channelListBytes)
	}
}

// RedisDelSubscribe 取消当前所有订阅
func (_Socket *Socket) RedisDelSubscribe(rdsConn redis.Conn, uuidStr string) {
	channelNameList, _ := redis.Strings(rdsConn.Do("HKEYS", RedisSubscribeChannelList+_Socket.key))

	for _, channel := range channelNameList {
		channelList := _Socket.GetSubscribeChannelList(rdsConn, channel)
		if _, ok := channelList[uuidStr]; ok {
			delete(channelList, uuidStr)

			channelListBytes, _ := json.Marshal(channelList)
			_, _ = rdsConn.Do("HSET", RedisSubscribeChannelList+_Socket.key, channel, channelListBytes)
		}
	}
}

// GetSubscribeChannelList 获取订阅通道列表
func (_Socket *Socket) GetSubscribeChannelList(rdsConn redis.Conn, channel string) map[string][]string {
	data := map[string][]string{}
	dataBytes, _ := redis.Bytes(rdsConn.Do("HGET", RedisSubscribeChannelList+_Socket.key, channel))
	_ = json.Unmarshal(dataBytes, &data)
	return data
}

// GetRedisSubscribeInfo 获取Redis订阅通道信息
func (_Socket *Socket) GetRedisSubscribeInfo(channel string) *RedisSubscribeChannel {
	for _, subscribeInfo := range _Socket.subscribeList {
		if subscribeInfo.Channel == channel {
			return subscribeInfo
		}
	}
	return nil
}
