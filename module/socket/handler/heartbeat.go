package handler

import (
	"github.com/gofiber/contrib/websocket"
	"gofiber/module/socket"
	"time"
)

const (
	// HeartbeatTime 心跳时间
	HeartbeatTime = 30 * time.Second
)

// sendHeartbeat 心跳包处理
func sendHeartbeat(socketInterface *socket.Socket, uuidStr string) {
	ticker := time.NewTicker(HeartbeatTime)
	defer ticker.Stop()

	for range ticker.C {
		if err := socketInterface.ConnWriteMessage(uuidStr, websocket.PingMessage, []byte{}); err != nil {
			socketInterface.CloseConn(uuidStr)
			return
		}
	}
}
