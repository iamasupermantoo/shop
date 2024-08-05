package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"gofiber/app/config"
)

// InitFaviconMiddleware 初始化 favicon 中间
func InitFaviconMiddleware() fiber.Handler {
	return favicon.New(favicon.Config{
		File: config.Conf.FileRoot + "/logo.png",
		URL:  "/logo.png",
	})
}
