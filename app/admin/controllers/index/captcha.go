package index

import (
	"github.com/dchest/captcha"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// NewCaptcha 创建验证码
func NewCaptcha(c *fiber.Ctx) error {
	err := c.SendString(captcha.NewLen(4))
	if err != nil {
		return err
	}
	return nil
}

// Captcha 显示验证码
func Captcha(c *fiber.Ctx) error {
	width, _ := strconv.Atoi(c.Params("width"))
	height, _ := strconv.Atoi(c.Params("height"))

	c.Type("png")
	return captcha.WriteImage(c.Context().Response.BodyWriter(), c.Params("id"), width, height)
}
