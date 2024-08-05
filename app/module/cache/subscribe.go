package cache

import (
	"github.com/gomodule/redigo/redis"
	"sync"
)

var Instance *SubScribe

// InitSubscribe 初始化订阅消息
func InitSubscribe() {
	rdsConn := Rds.Get()
	Instance = &SubScribe{
		clientMaps: make(map[string]func(data []byte)),
		subConn:    redis.PubSubConn{Conn: rdsConn},
	}

	// 初始化订阅消息
	go Instance.InitConsume()
}

// SubScribe 订阅
type SubScribe struct {
	sync       sync.Mutex                   //	锁机制
	subConn    redis.PubSubConn             //	对象
	clientMaps map[string]func(data []byte) //	订阅者
}

// Subscribe 订阅
func (_SubScribe *SubScribe) Subscribe(name string, fun func(data []byte)) error {
	err := _SubScribe.subConn.Subscribe(redis.Args{}.Add(name)...)
	if err != nil {
		return err
	}

	_SubScribe.sync.Lock()
	defer _SubScribe.sync.Unlock()
	_SubScribe.clientMaps[name] = fun
	return nil
}

// UnSubscribe 取消订阅
func (_SubScribe *SubScribe) UnSubscribe(name string) error {
	err := _SubScribe.subConn.Unsubscribe(redis.Args{}.Add(name)...)
	if err != nil {
		return err
	}

	_SubScribe.sync.Lock()
	defer _SubScribe.sync.Unlock()

	delete(_SubScribe.clientMaps, name)
	return nil
}
