package socket

import (
	"github.com/goccy/go-json"
	"github.com/gomodule/redigo/redis"
	"gofiber/app/module/cache"
)

// RedisPublish 发布消息 - 给某一个 uuid 发送消息
func (_Socket *Socket) RedisPublish(op string, receiver string, data interface{}) {
	dataBytes, _ := json.Marshal(&ConsumerMessage{
		Op:   op,
		UUID: receiver,
		Data: data,
	})
	cache.Instance.Publish(_Socket.key, dataBytes)
}

// RedisUserPublish 发布消息 给某个用户发送消息
func (_Socket *Socket) RedisUserPublish(rdsConn redis.Conn, op string, userId uint, data interface{}) {
	connInfoList := _Socket.ConnMaps.RedisGetConnInfo(rdsConn, userId)

	for _, connInfo := range connInfoList {
		_Socket.RedisPublish(op, connInfo.UUID, data)
	}
}

// ConnWriteJson 因为多线程的不允许直接调用当前方法
func (_Socket *Socket) ConnWriteJson(uuidStr string, data interface{}) error {
	connInfo := _Socket.ConnMaps.GetConn(uuidStr)
	_Socket.sync.Lock()
	defer _Socket.sync.Unlock()

	if connInfo != nil && connInfo.Conn != nil {
		return connInfo.Conn.WriteJSON(data)
	}
	return nil
}

// ConnWriteMessage 连接对象写入消息
func (_Socket *Socket) ConnWriteMessage(uuidStr string, messageType int, data []byte) error {
	connInfo := _Socket.ConnMaps.GetConn(uuidStr)

	_Socket.sync.Lock()
	defer _Socket.sync.Unlock()

	if connInfo != nil && connInfo.Conn != nil {
		return connInfo.Conn.WriteMessage(messageType, data)
	}
	return nil
}
