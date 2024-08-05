package client

import "github.com/gomodule/redigo/redis"

// SocketClientInterface socket客户端接口
type SocketClientInterface interface {
	// Connect 连接websocket
	Connect(addr string) SocketClientInterface

	// Reconnect 重新连接
	Reconnect()

	// Subscribe 订阅消息
	Subscribe(subscribe *Subscribe) error

	// UnSubscribe 取消订阅
	UnSubscribe(subscribe *Subscribe) error

	// GetSubscribe 获取订阅通道对象
	GetSubscribe(name string) *Subscribe

	// InitSubscribes 初始化订阅数据
	InitSubscribes(subscribeList []*Subscribe) SocketClientInterface

	// ConnWriteJson 发送JSON消息
	ConnWriteJson(data interface{}) error

	// ConnWriteMessage 发送字节消息
	ConnWriteMessage(messageType int, data []byte) error

	// Run 启动开始
	Run() error

	// SetWebSocketMessageFunc 设置消息处理
	SetWebSocketMessageFunc(fun func(rdsConn redis.Conn, msg []byte) error) SocketClientInterface

	// SetHeartbeatTime 设置心跳时间
	SetHeartbeatTime(second int) SocketClientInterface

	// SetReconnectTime 设置重连时间
	SetReconnectTime(second int) SocketClientInterface
}
