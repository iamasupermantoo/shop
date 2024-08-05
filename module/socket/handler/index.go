package handler

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gofiber/module/socket"
	"gofiber/utils"
)

// NewSocketConn 创建socket 连接
func NewSocketConn(socketInterface *socket.Socket) fiber.Handler {
	return websocket.New(func(conn *websocket.Conn) {
		// 设置当前连接对象
		uuidStr := uuid.New().String()

		// 设置连接
		socketInterface.ConnMaps.SetConn(uuidStr, conn)

		// 返回UUID
		_ = conn.WriteMessage(websocket.TextMessage, []byte(utils.JsonMarshal(&socket.ClientMessage{
			Op:   socket.MessageOperateInit,
			Data: uuidStr,
		})))

		// 心跳包设置
		go sendHeartbeat(socketInterface, uuidStr)
		defer func() {
			socketInterface.CloseConn(uuidStr)
			socketInterface.EventClose(socketInterface, uuidStr)
		}()

		var (
			msg []byte
			err error
		)

		// 处理业务
		for {
			// 读取消息
			if _, msg, err = conn.ReadMessage(); err != nil {
				break
			}

			// 消息事件
			err = socketInterface.EventMessage(socketInterface, uuidStr, msg)
			if err != nil {
				_ = socketInterface.ConnWriteJson(uuidStr, &socket.ClientMessage{
					Op:   socket.MessageOperateErr,
					Data: err.Error(),
				})
			}
		}
	})
}
