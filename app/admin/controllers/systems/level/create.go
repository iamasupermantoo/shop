package level

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// CreateParams 新增参数
type CreateParams struct {
	Name  string  `validate:"required" json:"name"`  // 名称
	Icon  string  `validate:"required" json:"icon"`  // 图标
	Money float64 `validate:"required" json:"money"` // 金额
}

// Create 新增接口
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	// 获取当前管理员最高等级
	var highLevel int
	if result := database.Db.Model(systemsModel.Level{}).
		Where("admin_id = ?", ctx.AdminId).
		Select("MAX(symbol)").
		Find(&highLevel); result.Error == nil && result.RowsAffected == 0 {
		return ctx.ErrorJson(result.Error.Error())
	}

	result := database.Db.Create(&systemsModel.Level{
		AdminId: ctx.AdminId,
		Name:    params.Name,
		Icon:    params.Icon,
		Money:   params.Money,
		Symbol:  highLevel + 1,
	})
	if result.Error != nil {
		return ctx.ErrorJson("添加失败, 原因 => " + result.Error.Error())
	}

	return ctx.SuccessJsonOK()
}
