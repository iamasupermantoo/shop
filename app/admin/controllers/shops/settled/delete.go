package settled

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// Delete 删除接口
func Delete(ctx *context.CustomCtx, params *context.DeleteParams) error {
	result := database.Db.Where("id IN ?", params.Ids).
		Where("admin_id IN ?", ctx.GetAdminChildIds()).Delete(&shopsModel.StoreSettled{})

	return ctx.IsErrorJson(result.Error)
}
