package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gofiber/app/config"
)

// InitRecoverMiddleware 初始化异常捕获
func InitRecoverMiddleware() fiber.Handler {
	var cfg = recover.Config{
		EnableStackTrace: config.Conf.Debug,
	}
	return recover.New(cfg)
}
