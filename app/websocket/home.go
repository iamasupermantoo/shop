package websocket

import (
	"github.com/goccy/go-json"
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/service/commonService"
	"gofiber/app/module/cache"
	"gofiber/module/socket"
	"time"
)

const (
	SubscribeChannelExample = "example"
)

var HomeWebSocket *socket.Socket

// InitHomeWebSocket 初始化前台WebSocket
func InitHomeWebSocket() {
	HomeWebSocket = socket.NewSocket(adminsModel.ServiceHomeName)
	HomeWebSocket.SetEventMessage(OnWebSocketMessageFunc)
	HomeWebSocket.SetEventClose(OnWebSocketCloseFunc)
	HomeWebSocket.SetEventOpen(OnWebSocketOpenFunc)

	// 默认通道初始化
	channels := make([]*socket.RedisSubscribeChannel, 0)
	channels = append(channels, &socket.RedisSubscribeChannel{
		Channel: SubscribeChannelExample, Args: []string{"example1", "example2"},
		ConsumerFunc: exampleReceiveMessage,
	})
	HomeWebSocket.InitSubscribeChannel(channels...)
}

// InitPublishMessage 初始化发布消息
func InitPublishMessage() {
	go examplePublishMessage()
}

// exampleReceiveMessage 通道接收消息
func exampleReceiveMessage(data []byte) {
	//fmt.Println("例子消息通道处理 ====> ", string(data))
}

// examplePublishMessage 通道消息发布
func examplePublishMessage() {
	ch := time.NewTicker(30 * time.Second)
	for {
		//cache.Instance.Publish(SubscribeChannelExample, "例子消息测试数据")
		<-ch.C
	}
}

// OnWebSocketMessageFunc 消息事件
func OnWebSocketMessageFunc(socketInterface *socket.Socket, uuidStr string, msg []byte) error {
	dataMsg := &socket.ClientMessage{}
	err := json.Unmarshal(msg, &dataMsg)
	if err != nil {
		return err
	}
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	switch dataMsg.Op {
	case socket.MessageOperateSubscribe:
		// 订阅通道
		subscribe := &socket.SubscribeInfo{}
		dataBytes, _ := json.Marshal(dataMsg.Data)
		err = json.Unmarshal(dataBytes, &subscribe)
		if err != nil {
			return err
		}

		if subscribe.Channel != "" && len(subscribe.Args) > 0 {
			socketInterface.Subscribe(rdsConn, uuidStr, subscribe.Channel, subscribe.Args)
			socketInterface.RedisPublish(socket.MessageOperateSubscribe, uuidStr, &socket.ClientMessage{
				Op:   socket.MessageOperateSubscribe,
				Data: "ok",
			})
		}

	case socket.MessageOperateUnSubscribe:
		// 取消订阅通道
		subscribe := &socket.SubscribeInfo{}
		dataBytes, _ := json.Marshal(dataMsg.Data)
		err = json.Unmarshal(dataBytes, &subscribe)
		if err != nil {
			return err
		}

		if subscribe.Channel != "" && len(subscribe.Args) > 0 {
			socketInterface.UnSubscribe(rdsConn, uuidStr, subscribe.Channel, subscribe.Args)
			socketInterface.RedisPublish(socket.MessageOperateUnSubscribe, uuidStr, &socket.ClientMessage{
				Op:   socket.MessageOperateUnSubscribe,
				Data: "ok",
			})
		}
	case socket.MessageOperateBindUser:
		adminId, userId := commonService.NewServiceToken(rdsConn).VerifyToken(adminsModel.ServiceHomeName, dataMsg.Data.(string))

		// 绑定用户ID
		socketInterface.ConnMaps.RedisSetConnInfo(rdsConn, userId, &socket.UserConnInfo{
			UUID: uuidStr, Key: socketInterface.GetSocketKey(), UserId: userId, AdminId: adminId,
			Origin: "", Device: "", IP: "",
		})
		socketInterface.RedisPublish(socket.MessageOperateBindUser, uuidStr, &socket.ClientMessage{
			Op:   socket.MessageOperateBindUser,
			Data: "ok",
		})

		// 上线通知
		socketInterface.EventOpen(socketInterface, uuidStr)
	default:
	}

	return nil
}

// OnWebSocketCloseFunc 关闭事件方法
func OnWebSocketCloseFunc(socketInterface *socket.Socket, uuidStr string) {

}

// OnWebSocketOpenFunc 打开事件方法
func OnWebSocketOpenFunc(socketInterface *socket.Socket, uuidStr string) {

}
