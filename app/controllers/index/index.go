package index

import (
	"gofiber/app/module/context"
)

// Index 根路径
func Index(ctx *context.CustomCtx, params *context.NoRequestBody) error {
	return ctx.SuccessJsonOK()
}
