package cache

import (
	"fmt"
	"gofiber/app/config"
	"time"

	"github.com/gomodule/redigo/redis"
)

var Rds *redis.Pool

// 初始化Redis
func init() {
	Rds = &redis.Pool{
		MaxIdle:     config.Conf.Redis.MaxIdleConn,
		MaxActive:   config.Conf.Redis.MaxOpenConn,
		IdleTimeout: time.Duration(config.Conf.Redis.ConnectTimeout) * time.Second,
		Wait:        config.Conf.Redis.Wait,
		Dial: func() (redis.Conn, error) {
			host := fmt.Sprintf("%s:%d", config.Conf.Redis.Server, config.Conf.Redis.Port)
			conn, err := redis.Dial(
				config.Conf.Redis.Network,
				host,
				redis.DialPassword(config.Conf.Redis.Pass),
				redis.DialDatabase(config.Conf.Redis.DbName),
				redis.DialConnectTimeout(time.Duration(config.Conf.Redis.ConnectTimeout)*time.Second),
				redis.DialReadTimeout(time.Duration(config.Conf.Redis.ReadTimeout)*time.Second),
				redis.DialWriteTimeout(time.Duration(config.Conf.Redis.WriteTimeout)*time.Second),
			)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
	// 初始化订阅消息
	InitSubscribe()
}
