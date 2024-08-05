package socket

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gomodule/redigo/redis"
	"sync"
)

// ConnMaps 客户端Maps
type ConnMaps struct {
	sync.Mutex
	key  string // SocketKey
	maps map[string]*ConnInfo
}

type ConnInfo struct {
	Conn   *websocket.Conn // 连接对象
	UserId uint            //	用户ID
}

// SetConn 设置客户端
func (_ConnMaps *ConnMaps) SetConn(uuidStr string, conn *websocket.Conn) *ConnMaps {
	_ConnMaps.Lock()
	defer _ConnMaps.Unlock()

	_ConnMaps.maps[uuidStr] = &ConnInfo{Conn: conn}
	return _ConnMaps
}

// SetConnUserId 设置连接用户ID
func (_ConnMaps *ConnMaps) SetConnUserId(uuidStr string, userId uint) *ConnMaps {
	connInfo := _ConnMaps.GetConn(uuidStr)

	_ConnMaps.Lock()
	defer _ConnMaps.Unlock()
	if connInfo != nil {
		connInfo.UserId = userId
	}
	return _ConnMaps
}

// GetConn 获取客户端
func (_ConnMaps *ConnMaps) GetConn(uuidStr string) *ConnInfo {
	_ConnMaps.Lock()
	defer _ConnMaps.Unlock()

	if _, ok := _ConnMaps.maps[uuidStr]; ok {
		return _ConnMaps.maps[uuidStr]

	}
	return nil
}

// CloseConn 关闭客户端
func (_ConnMaps *ConnMaps) CloseConn(rdsConn redis.Conn, uuidStr string) {
	connInfo := _ConnMaps.GetConn(uuidStr)
	_ = connInfo.Conn.Close()

	// 删除 用户连接列表中的当前连接
	_ConnMaps.RedisDelUUIDConnInfo(rdsConn, 0, uuidStr)

	_ConnMaps.Lock()
	defer _ConnMaps.Unlock()
	delete(_ConnMaps.maps, uuidStr)
}
