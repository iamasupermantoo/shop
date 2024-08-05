package client

import (
	"github.com/fasthttp/websocket"
	"github.com/goccy/go-json"
	"github.com/gomodule/redigo/redis"
	"gofiber/app/module/cache"
	"sync"
	"time"
)

// SocketClient socket客户端
type SocketClient struct {
	sync                 sync.Mutex                                 //	锁
	err                  error                                      //	错误信息
	addr                 string                                     // 	连接地址
	conn                 *websocket.Conn                            //	连接对象
	subscribes           []*Subscribe                               //	订阅数据
	heartbeatTime        time.Duration                              //	心跳时间
	reconnectTime        time.Duration                              //	重连时间
	websocketMessageFunc func(rdsConn redis.Conn, msg []byte) error //	websocket消息处理方法
}

// NewSocketClient 创建新的连接对象
func NewSocketClient(addr string) SocketClientInterface {
	client := &SocketClient{
		addr:          addr,
		sync:          sync.Mutex{},
		heartbeatTime: 10 * time.Second,
		reconnectTime: 10 * time.Second,
		subscribes:    make([]*Subscribe, 0),
		websocketMessageFunc: func(rdsConn redis.Conn, msg []byte) error {
			return nil
		},
	}
	client.Connect(addr)
	return client
}

// Run 运行读取
func (_SocketClient *SocketClient) Run() error {
	if _SocketClient.err != nil {
		return _SocketClient.err
	}

	// 订阅数据
	for _, subscribe := range _SocketClient.subscribes {
		if subscribe.Data == "" || subscribe.Data == nil {
			continue
		}
		dataBytes, _ := json.Marshal(subscribe.Data)
		_ = _SocketClient.ConnWriteMessage(websocket.TextMessage, dataBytes)
	}

	go func() {
		rdsCon := cache.Rds.Get()
		defer func(rdsCon redis.Conn) {
			// 重新连接
			_SocketClient.Reconnect()
			_ = rdsCon.Close()
		}(rdsCon)

		var msg []byte
		var err error
		for {
			_, msg, err = _SocketClient.conn.ReadMessage()
			if err != nil {
				return
			}

			// 读取消息
			_ = _SocketClient.websocketMessageFunc(rdsCon, msg)
		}
	}()
	return nil
}

// SetWebSocketMessageFunc 设置消息处理
func (_SocketClient *SocketClient) SetWebSocketMessageFunc(fun func(rdsConn redis.Conn, msg []byte) error) SocketClientInterface {
	_SocketClient.websocketMessageFunc = fun
	return _SocketClient
}

// InitSubscribes 初始化订阅数据
func (_SocketClient *SocketClient) InitSubscribes(subscribeList []*Subscribe) SocketClientInterface {
	_SocketClient.subscribes = subscribeList
	return _SocketClient
}

// Subscribe 订阅消息
func (_SocketClient *SocketClient) Subscribe(subscribe *Subscribe) error {
	oldSubscribe := _SocketClient.GetSubscribe(subscribe.Name)
	if oldSubscribe == nil {
		_SocketClient.subscribes = append(_SocketClient.subscribes, subscribe)
	}
	dataBytes, _ := json.Marshal(subscribe.Data)
	return _SocketClient.ConnWriteMessage(websocket.TextMessage, dataBytes)
}

// UnSubscribe 取消订阅
func (_SocketClient *SocketClient) UnSubscribe(subscribe *Subscribe) error {
	for i := 0; i < len(_SocketClient.subscribes); i++ {
		if _SocketClient.subscribes[i].Name == subscribe.Name {
			_SocketClient.subscribes = append(_SocketClient.subscribes[:i], _SocketClient.subscribes[i+1:]...)
			dataBytes, _ := json.Marshal(subscribe.Data)
			return _SocketClient.ConnWriteMessage(websocket.TextMessage, dataBytes)
		}
	}
	return nil
}

// GetSubscribe 获取订阅通道对象
func (_SocketClient *SocketClient) GetSubscribe(name string) *Subscribe {
	for _, subscribe := range _SocketClient.subscribes {
		if subscribe.Name == name {
			return subscribe
		}
	}
	return nil
}

// ConnWriteJson 发送JSON消息
func (_SocketClient *SocketClient) ConnWriteJson(data interface{}) error {
	_SocketClient.sync.Lock()
	defer _SocketClient.sync.Unlock()

	return _SocketClient.conn.WriteJSON(data)
}

// ConnWriteMessage 发送字节消息
func (_SocketClient *SocketClient) ConnWriteMessage(messageType int, data []byte) error {
	_SocketClient.sync.Lock()
	defer _SocketClient.sync.Unlock()

	return _SocketClient.conn.WriteMessage(messageType, data)
}

// SetHeartbeatTime 设置心跳时间
func (_SocketClient *SocketClient) SetHeartbeatTime(second int) SocketClientInterface {
	_SocketClient.heartbeatTime = time.Duration(second) * time.Second
	return _SocketClient
}

// SetReconnectTime 设置重连时间
func (_SocketClient *SocketClient) SetReconnectTime(second int) SocketClientInterface {
	_SocketClient.reconnectTime = time.Duration(second) * time.Second
	return _SocketClient
}

// Connect 连接
func (_SocketClient *SocketClient) Connect(addr string) SocketClientInterface {
	_SocketClient.conn, _, _SocketClient.err = websocket.DefaultDialer.Dial(addr, nil)

	//	如果连接成功启动心跳处理
	if _SocketClient.err == nil {
		go func() {
			ch := time.NewTicker(_SocketClient.heartbeatTime)

			for {
				if err := _SocketClient.ConnWriteMessage(websocket.PingMessage, []byte{}); err != nil {
					_ = _SocketClient.conn.Close()
					return
				}
				<-ch.C
			}
		}()
	}

	return _SocketClient
}

// Reconnect 重新连接
func (_SocketClient *SocketClient) Reconnect() {
	time.AfterFunc(_SocketClient.reconnectTime, func() {
		_SocketClient.Connect(_SocketClient.addr)
		if _SocketClient.err != nil {
			_SocketClient.Reconnect()
			return
		}

		// 重新发送订阅, 并且启动
		_ = _SocketClient.InitSubscribes(_SocketClient.subscribes).Run()
	})
}
