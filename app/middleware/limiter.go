package middleware

import (
	"gofiber/app/config"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// InitLimiterMiddleware 初始化速率限制
func InitLimiterMiddleware() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        config.Conf.Middleware.Limiter.Max,
		Expiration: time.Duration(config.Conf.Middleware.Limiter.Expiration) * time.Second,
	})
}
