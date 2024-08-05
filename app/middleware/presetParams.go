package middleware

import "github.com/gofiber/fiber/v2"

// PresetParams 预设参数中间件
func PresetParams(preset interface{}) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.Locals("preset", preset)
		return ctx.Next()
	}
}
