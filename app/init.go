package app

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"gofiber/app/config"
	"gofiber/app/controllers"
	"gofiber/app/middleware"
)

// InitWebApp 初始化前台App
func InitWebApp() *fiber.App {
	webApp := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// 静态资源文件
	webApp.Static("/", config.Conf.FileRoot)

	// 速率限制中间件
	webApp.Use(middleware.InitLimiterMiddleware())

	// 初始化前台路由
	controllers.InitWebRouter(webApp)

	return webApp
}
