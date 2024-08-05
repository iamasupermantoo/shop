package socket

// OnWebSocketMessageFunc 消息事件
func OnWebSocketMessageFunc(socketInstance *Socket, uuidStr string, msg []byte) error {
	return nil
}

// OnWebSocketOpenFunc 打开事件方法
func OnWebSocketOpenFunc(socketInstance *Socket, uuidStr string) {

}

// OnWebSocketCloseFunc 关闭事件方法
func OnWebSocketCloseFunc(socketInstance *Socket, uuidStr string) {

}
