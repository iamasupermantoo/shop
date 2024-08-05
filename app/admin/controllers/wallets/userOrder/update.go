package userOrder

import (
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// UpdateParams 更新参数
type UpdateParams struct {
	ID       uint    `gorm:"-" validate:"required" json:"id"` //	ID
	AdminId  uint    `json:"adminId"`                         // 管理ID
	UserId   uint    `json:"userId"`                          // 用户ID
	AssetsId uint    `json:"assetsId"`                        // 资产ID
	SourceId uint    `json:"sourceId"`                        // 来源ID
	Type     int     `json:"type"`                            // 类型 1充值类型 2资产充值类型 11提现类型 12资产提现类型
	OrderSn  string  `json:"orderSn"`                         // 编号
	Money    float64 `json:"money"`                           // 金额
	Fee      float64 `json:"fee"`                             // 手续费
	Voucher  string  `json:"voucher"`                         // 支付凭证
	Status   int     `json:"status"`                          // 状态  -1禁用 10审核 20完成
	Data     string  `json:"data"`                            // 数据
}

// Update 更新接口
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	result := database.Db.Model(&walletsModel.WalletUserOrder{}).
		Where("id = ?", params.ID).Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Updates(params)
	return ctx.IsErrorJson(result.Error)
}
