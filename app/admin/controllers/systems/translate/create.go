package translate

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// CreateParams 新增参数
type CreateParams struct {
	Type  int    `validate:"required" json:"type"`  // 类型
	Lang  string `validate:"required" json:"lang"`  // 语言标识
	Name  string `validate:"required" json:"name"`  // 名称
	Field string `validate:"required" json:"field"` // 键名
	Value string `validate:"required" json:"value"` // 键值
}

// Create 新增接口
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	result := database.Db.Create(&systemsModel.Translate{
		AdminId: ctx.AdminId,
		Lang:    params.Lang,
		Name:    params.Name,
		Field:   params.Field,
		Value:   params.Value,
	})
	if result.Error != nil {
		return ctx.ErrorJson("添加失败, 原因 => " + result.Error.Error())
	}

	return ctx.SuccessJsonOK()
}
