package socket

import (
	"github.com/goccy/go-json"
	"github.com/gomodule/redigo/redis"
	"gofiber/app/module/cache"
	"sync"
)

const (
	// MessageOperateInit 初始化操作
	MessageOperateInit = "init"

	// MessageOperateErr 错误操作
	MessageOperateErr = "err"

	// MessageOperateSubscribe 消息操作订阅
	MessageOperateSubscribe = "subscribe"

	// MessageOperateUnSubscribe 消息操作取消订阅
	MessageOperateUnSubscribe = "unsubscribe"

	// MessageOperateBindUser 绑定用户ID
	MessageOperateBindUser = "bindUser"

	// MessageOperateMessage 传递消息
	MessageOperateMessage = "message"

	// MessageOperateAudio 音频消息
	MessageOperateAudio = "audio"

	// MessageOperateOnline 上线操作
	MessageOperateOnline = "online"

	// MessageOperateOffline 下线操作
	MessageOperateOffline = "offline"

	// MessageOperateCancel 消息撤销
	MessageOperateCancel = "messageCancel"

	// MessageOperateReadMsg 消息已读
	MessageOperateReadMsg = "readMessage"
)

// Socket WebSocket
type Socket struct {
	sync          sync.Mutex                                                     //	锁操作
	key           string                                                         //	当前标识
	subscribeList []*RedisSubscribeChannel                                       // Redis 初始化订阅通道
	ConnMaps      *ConnMaps                                                      //	客户连接对象
	EventMessage  func(socketInstance *Socket, uuidStr string, msg []byte) error //	websocket 消息事件
	EventOpen     func(socketInstance *Socket, uuidStr string)                   //	websocket 打开事件
	EventClose    func(socketInstance *Socket, uuidStr string)                   //	websocket 关闭事件
}

// ClientMessage 客户端消息结构
type ClientMessage struct {
	Op   string      `json:"op"`   //	方法名称
	Data interface{} `json:"data"` //	方法参数
}

// ConsumerMessage 消费者消息结构
type ConsumerMessage struct {
	Op   string      `json:"op"`   //	操作
	UUID string      `json:"uuid"` //	连接标识
	Data interface{} `json:"data"` //	发送数据
}

// RedisSubscribeChannel Redis 订阅通道信息
type RedisSubscribeChannel struct {
	Channel      string            //	通道名称
	Args         []string          //	通道参数 ｜ 用于过滤没有的参数
	ConsumerFunc func(data []byte) //	通道消费方法
}

// NewSocket 创建socket
func NewSocket(socketKey string) *Socket {
	socket := &Socket{
		sync:          sync.Mutex{},
		key:           socketKey,
		subscribeList: make([]*RedisSubscribeChannel, 0),
		ConnMaps: &ConnMaps{
			key:  socketKey,
			maps: map[string]*ConnInfo{},
		},
		EventOpen:    OnWebSocketOpenFunc,
		EventClose:   OnWebSocketCloseFunc,
		EventMessage: OnWebSocketMessageFunc,
	}

	// 借用Redis 开通订阅通道, 接收消息队列, 给当前用户发送消息消费
	_ = cache.Instance.Subscribe(socketKey, func(data []byte) {
		writeMsg := &ConsumerMessage{}
		err := json.Unmarshal(data, &writeMsg)
		if err == nil {
			_ = socket.ConnWriteJson(writeMsg.UUID, &ClientMessage{
				Op:   writeMsg.Op,
				Data: writeMsg.Data,
			})
		}
	})
	return socket
}

// InitSubscribeChannel 初始化订阅通道
func (_Socket *Socket) InitSubscribeChannel(channelList ...*RedisSubscribeChannel) {
	_Socket.subscribeList = channelList
	for _, channelInfo := range channelList {
		_ = cache.Instance.Subscribe(channelInfo.Channel, channelInfo.ConsumerFunc)
	}
}

// SetEventMessage 设置消息事件
func (_Socket *Socket) SetEventMessage(fun func(socketInstance *Socket, uuidStr string, msg []byte) error) *Socket {
	_Socket.EventMessage = fun
	return _Socket
}

// SetEventOpen 设置打开事件
func (_Socket *Socket) SetEventOpen(fun func(socketInstance *Socket, uuidStr string)) *Socket {
	_Socket.EventOpen = fun
	return _Socket
}

// SetEventClose 设置关闭事件
func (_Socket *Socket) SetEventClose(fun func(socketInstance *Socket, uuidStr string)) *Socket {
	_Socket.EventClose = fun
	return _Socket
}

// IsOnline 是否在线
func (_Socket *Socket) IsOnline(rdsConn redis.Conn, userId uint) bool {
	userConnList := _Socket.ConnMaps.RedisGetConnInfo(rdsConn, userId)
	return len(userConnList) > 0
}

// GetSocketKey 获取socket 标识
func (_Socket *Socket) GetSocketKey() string {
	return _Socket.key
}

// CloseConn 关闭当前连接
func (_Socket *Socket) CloseConn(uuidStr string) {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	// 关闭连接并且关闭用户关联
	_Socket.ConnMaps.CloseConn(rdsConn, uuidStr)

	// 关闭当前订阅数据
	_Socket.RedisDelSubscribe(rdsConn, uuidStr)
}
