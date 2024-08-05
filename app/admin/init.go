package admin

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"gofiber/app/admin/controllers"
	"gofiber/app/config"
)

// InitAdminApp 初始化后端App
func InitAdminApp() *fiber.App {
	adminApp := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// 静态资源文件
	adminApp.Static("/", config.Conf.FileRoot)

	// 载入后端路由
	controllers.InitAdminRouter(adminApp)

	return adminApp
}
