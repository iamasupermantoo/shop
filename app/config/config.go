package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

// Conf 全局配置文件
var Conf *Config

func init() {
	conFile, err := os.ReadFile("./app.yaml")
	if err != nil {
		panic(err)
	}

	//	映射配置文件
	err = yaml.Unmarshal(conFile, &Conf)
	if err != nil {
		panic(err)
	}
}

// Config 配置文件
type Config struct {
	Name       string      `yaml:"Name"`       //	项目名称
	Port       string      `yaml:"Port"`       //	启动端口
	Debug      bool        `yaml:"Debug"`      //	Debug模式
	TimeZone   string      `yaml:"TimeZone"`   //	时区
	Lang       string      `yaml:"Lang"`       //	默认语言
	FileRoot   string      `yaml:"FileRoot"`   //	对外开放文件路径
	IsPreFork  bool        `yaml:"IsPreFork"`  //	是否开启子进程
	Database   *Database   `yaml:"Database"`   //	数据库配置
	Redis      *Redis      `yaml:"Redis"`      //	缓存配置
	Middleware *Middleware `yaml:"Middleware"` //	中间件配置
}

// Middleware 中间件
type Middleware struct {
	Limiter *MiddlewareLimiter `yaml:"Limiter"` //	速率限制
	Logger  *MiddlewareLogger  `yaml:"Logger"`  //	日志记录
}

// MiddlewareLimiter 速率中间件
type MiddlewareLimiter struct {
	Max        int `yaml:"Max"`        //	速率最大值
	Expiration int `yaml:"Expiration"` //	过期时间
}

// MiddlewareLogger 记录中间件
type MiddlewareLogger struct {
}

// Database 数据库配置
type Database struct {
	Network string `yaml:"Network"` //	网络
	Server  string `yaml:"Server"`  //	服务地址
	Port    int    `yaml:"Port"`    //	端口
	User    string `yaml:"User"`    //	用户名
	Pass    string `yaml:"Pass"`    //	密码
	DbName  string `yaml:"DbName"`  //	数据库
}

// Redis Redis配置
type Redis struct {
	Network         string `yaml:"Network"`         //	网络
	Server          string `yaml:"Server"`          //	服务器地址
	Port            int    `yaml:"Port"`            //	端口
	Pass            string `yaml:"Pass"`            //	密码
	DbName          int    `yaml:"DbName"`          //	数据库
	ConnectTimeout  int64  `yaml:"ConnectTimeout"`  //	连接超时时间
	ReadTimeout     int64  `yaml:"ReadTimeout"`     //	读取超时时间
	WriteTimeout    int64  `yaml:"WriteTimeout"`    //	写入超时时间
	MaxOpenConn     int    `yaml:"MaxOpenConn"`     // 	设置最大连接数
	ConnMaxIdleTime int64  `yaml:"ConnMaxIdleTime"` // 	空闲连接超时
	MaxIdleConn     int    `yaml:"MaxIdleConn"`     // 	最大空闲连接数
	Wait            bool   `yaml:"Wait"`            // 	如果超过最大连接数是否等待
}
