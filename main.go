package main

import (
	"github.com/dchest/captcha"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gomodule/redigo/redis"
	"gofiber/app"
	"gofiber/app/admin"
	"gofiber/app/config"
	"gofiber/app/crontab"
	"gofiber/app/middleware"
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/module/cache"
	"gofiber/app/websocket"
)

func main() {
	//	项目APP
	fiberApp := fiber.New(fiber.Config{
		AppName:      config.Conf.Name,
		ServerHeader: config.Conf.Name,
		Prefork:      config.Conf.IsPreFork,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})

	// 记录中间件
	fiberApp.Use(middleware.InitLoggerMiddleware())

	// 跨域中间件
	fiberApp.Use(middleware.InitCorsMiddleware())

	// 异常中间件
	fiberApp.Use(middleware.InitRecoverMiddleware())

	// 缓存Favicon图标
	fiberApp.Use(middleware.InitFaviconMiddleware())

	// 验证码使用Redis存储
	captcha.SetCustomStore(&CaptchaStore{})

	// 初始化订阅通道
	websocket.InitWebSocket()

	// 主进程运行方法
	if !fiber.IsChild() {
		// 运行定时任务
		crontab.InitCrontab()

		// 初始化发布消息
		websocket.InitPublishMessage()
	}

	// 初始化运行项目
	middleware.InitService(map[string]*middleware.Service{
		adminsModel.ServiceAdminRouteName: {Fiber: admin.InitAdminApp()},
		adminsModel.ServiceHomeName:       {Fiber: app.InitWebApp()},
	})
	fiberApp.Use(middleware.InitDomainRoutingMiddleware())

	// 启动监听端口
	_ = fiberApp.Listen(":" + config.Conf.Port)
}

type CaptchaStore struct {
}

func (_CaptchaStore *CaptchaStore) Set(id string, digits []byte) {
	rds := cache.Rds.Get()
	defer rds.Close()

	_, _ = rds.Do("SETEX", "captcha_"+id, 300, digits)
}

func (_CaptchaStore *CaptchaStore) Get(id string, clear bool) (digits []byte) {
	rds := cache.Rds.Get()
	defer rds.Close()

	digits, _ = redis.Bytes(rds.Do("GET", "captcha_"+id))
	return digits
}
