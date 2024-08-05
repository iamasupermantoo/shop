package userAccount

import (
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// UpdateParams 更新参数
type UpdateParams struct {
	ID       uint   `gorm:"-" validate:"required" json:"id"` //	ID
	Name     string `json:"name"`                            // 银行名称｜Token
	RealName string `json:"realName"`                        // 真实姓名
	Number   string `json:"number"`                          // 卡号|地址
	Code     string `json:"code"`                            // 银行代码
	Remark   string `json:"remark"`                          // 备注信息
	Status   int    `json:"status"`                          // 状态  -1禁用 10开启
}

// Update 更新接口
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	result := database.Db.Model(&walletsModel.WalletUserAccount{}).
		Where("id = ?", params.ID).Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Updates(params)
	return ctx.IsErrorJson(result.Error)
}
