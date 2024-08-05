package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gofiber/app/config"
	"gofiber/utils"
	"os"
)

// InitLoggerMiddleware 初始化记录中间件
func InitLoggerMiddleware() fiber.Handler {
	accessFile := os.Stdout
	if !config.Conf.Debug {
		accessFile, _ = os.OpenFile("./logs/access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	}

	return logger.New(logger.Config{
		TimeZone:      config.Conf.TimeZone,
		TimeFormat:    "2006-01-02T15:04:05.000Z0700",
		DisableColors: !config.Conf.Debug,
		Output:        accessFile,
		Format:        "${time} | ${status} | ${latency} | ${IP4} | ${method} | ${path} | ${error}\n",
		CustomTags: map[string]logger.LogFunc{
			"IP4": func(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
				return output.WriteString(utils.GetClientIP(c))
			},
		},
	})
}
