package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// InitCorsMiddleware 初始化跨域中间件
func InitCorsMiddleware() fiber.Handler {
	var config = cors.Config{}

	return cors.New(config)
}
