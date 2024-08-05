package socket

import (
	"github.com/goccy/go-json"
	"github.com/gomodule/redigo/redis"
)

const (
	RedisUserConnInfoList = "RedisUserConnInfoList"
)

// UserConnInfo 连接的用户信息
type UserConnInfo struct {
	UUID    string `json:"uuid"`    //	UUID
	Key     string `json:"key"`     //	socket 标识
	UserId  uint   `json:"userId"`  //	用户ID
	AdminId uint   `json:"adminId"` // 	管理ID
	Device  string `json:"device"`  //	设备信息
	Origin  string `json:"origin"`  //	源域名
	IP      string `json:"ip"`      //	IP信息
}

// RedisSetConnInfo 设置｜更新 用户连接信息
func (_ConnMaps *ConnMaps) RedisSetConnInfo(rdsConn redis.Conn, userId uint, currentInfo *UserConnInfo) {
	_ConnMaps.Lock()
	defer _ConnMaps.Unlock()

	userConnInfoList := _ConnMaps.RedisGetConnInfo(rdsConn, userId)

	connIndexOf := -1
	for connIndex, connInfo := range userConnInfoList {
		if connInfo.UUID == currentInfo.UUID {
			connIndexOf = connIndex
		}
	}
	if connIndexOf == -1 {
		userConnInfoList = append(userConnInfoList, currentInfo)
	} else {
		userConnInfoList[connIndexOf] = currentInfo
	}

	connInfoListBytes, _ := json.Marshal(userConnInfoList)
	_, _ = rdsConn.Do("HSET", RedisUserConnInfoList+_ConnMaps.key, userId, connInfoListBytes)
}

// RedisGetConnInfo 获取用户连接信息
func (_ConnMaps *ConnMaps) RedisGetConnInfo(rdsConn redis.Conn, userId uint) []*UserConnInfo {
	userConnInfoList := make([]*UserConnInfo, 0)
	userConnInfoListBytes, err := redis.Bytes(rdsConn.Do("HGET", RedisUserConnInfoList+_ConnMaps.key, userId))
	if err == nil {
		_ = json.Unmarshal(userConnInfoListBytes, &userConnInfoList)
	}
	return userConnInfoList
}

// RedisDelUUIDConnInfo 删除用户uuid 连接数据
func (_ConnMaps *ConnMaps) RedisDelUUIDConnInfo(rdsConn redis.Conn, userId uint, uuidStr string) {
	_ConnMaps.Lock()
	defer _ConnMaps.Unlock()

	userConnInfoList := _ConnMaps.RedisGetConnInfo(rdsConn, userId)

	// 如果只有一条数据, 那么直接删除
	if len(userConnInfoList) == 1 {
		_ConnMaps.RedisDelConnInfo(rdsConn, userId)
		return
	}

	// 如果多个连接, 那么删除对应的连接
	connIndexOf := -1
	for connIndex, connInfo := range userConnInfoList {
		if connInfo.UUID == uuidStr {
			connIndexOf = connIndex
		}
	}
	if connIndexOf > 0 {
		userConnInfoList = append(userConnInfoList[:connIndexOf], userConnInfoList[connIndexOf+1:]...)
		connInfoBytes, _ := json.Marshal(userConnInfoList)
		_, _ = rdsConn.Do("HSET", RedisUserConnInfoList+_ConnMaps.key, userId, connInfoBytes)
	}
}

// RedisDelConnInfo 删除用户连接信息
func (_ConnMaps *ConnMaps) RedisDelConnInfo(rdsConn redis.Conn, userId uint) {
	_, _ = rdsConn.Do("HDEL", RedisUserConnInfoList+_ConnMaps.key, userId)
}
