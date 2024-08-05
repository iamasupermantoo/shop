package middleware

import (
	"crypto/rsa"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

const (
	// AuthScheme 标头
	AuthScheme = "Bearer"
)

// InitJwtMiddleware 初始化Jwt中间件
func InitJwtMiddleware(privateKey *rsa.PrivateKey, successHandler func(c *fiber.Ctx) error) fiber.Handler {
	if privateKey == nil {
		panic("没有设置 RS256 Key")
	}

	var config = jwtware.Config{
		ContextKey: "token",
		AuthScheme: AuthScheme,
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.RS256,
			Key:    privateKey.Public(),
		},
		SuccessHandler: successHandler,
	}

	return jwtware.New(config)
}
