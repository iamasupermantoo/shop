package context

import (
	validatorV10 "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gomodule/redigo/redis"
	"gofiber/app/module/cache"
)

type Handler[T any] func(*CustomCtx, *T) error
type NoRequestBody struct{}

type CustomCtx struct {
	*fiber.Ctx                //	fiberCtx 对象
	Rds            redis.Conn //	Redis缓存对象
	OriginHost     string     // 	来源域名
	Lang           string     // 	当前语言
	AdminId        uint       //	管理ID
	AdminSettingId uint       //	管理设置ID
	UserId         uint       //	用户ID
}

func NewCustomCtx(ctx *fiber.Ctx) *CustomCtx {
	return &CustomCtx{Ctx: ctx}
}

// WrapCustomHandler 包装自定义函数
func WrapCustomHandler[T any](h Handler[T]) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// 将 *fiber.Ctx 转换为 CustomCtx
		customCtx := &CustomCtx{Ctx: ctx}

		// 使用工厂函数创建一个新的实例来解析请求体
		var body T
		if _, ok := any(&body).(*NoRequestBody); !ok {
			if err := ctx.BodyParser(&body); err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(customCtx.ErrorJson(err.Error()))
			}
		}

		// 初始化需要使用的参数
		customCtx.Rds = cache.Rds.Get()
		customCtx.initLang()
		customCtx.initOriginHost()
		customCtx.initClaims()

		// 检测参数结构体
		err := validate.Struct(&body)
		if err != nil {
			for _, errs := range err.(validatorV10.ValidationErrors) {
				return customCtx.validateError(errs.Tag(), errs.Field(), errs.Param())
			}
		}

		// 调用自定义处理函数
		err = h(customCtx, &body)

		// 释放Rds
		_ = customCtx.Rds.Close()
		return err
	}
}
