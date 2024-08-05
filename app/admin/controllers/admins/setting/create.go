package setting

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// CreateParams 新增参数
type CreateParams struct {
	GroupId int    `json:"groupId"` // 分组ID
	Name    string `json:"name"`    // 设置名称
	Type    int    `json:"type"`    // input类型
	Field   string `json:"field"`   // 建铭
	Value   string `json:"value"`   // 键值
	Data    string `json:"data"`    // input配置
}

// Create 新增接口
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	createInfo := &adminsModel.AdminSetting{
		AdminId: ctx.AdminId,
		GroupId: params.GroupId,
		Name:    params.Name,
		Type:    params.Type,
		Field:   params.Field,
		Value:   params.Value,
		Data:    params.Data,
	}

	result := database.Db.Create(createInfo)
	if result.Error != nil {
		return ctx.ErrorJson("添加失败, 原因 => " + result.Error.Error())
	}

	return ctx.SuccessJsonOK()
}
