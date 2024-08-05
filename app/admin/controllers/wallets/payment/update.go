package payment

import (
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// UpdateParams 更新参数
type UpdateParams struct {
	ID        uint    `gorm:"-" validate:"required" json:"id"` // ID
	Name      string  `json:"name"`                            // 名称
	Symbol    string  `json:"symbol"`                          // 标识
	Icon      string  `json:"icon"`                            // 图标
	Rate      float64 `json:"rate"`                            // 汇率
	Type      int     `json:"type"`                            // 类型 1银行卡类型 11数字货币类型 21第三方支付
	Mode      int     `json:"mode"`                            // 模式 1充值模式 2资产充值模式 11提现模式 12资产提现模式
	IsVoucher int     `json:"isVoucher"`                       // 显示凭证
	Status    int     `json:"status"`                          // 状态
	Desc      string  `json:"desc"`                            // 详情
}

// Update 更新接口
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	// 获取钱包支付信息
	walletPaymentInfo := &walletsModel.WalletPayment{}
	if result := database.Db.
		Where(params.ID).
		Find(walletPaymentInfo); result.RowsAffected == 0 {
		return ctx.ErrorJson(result.Error.Error())
	}

	result := database.Db.Model(&walletsModel.WalletPayment{}).
		Where("id = ?", params.ID).Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Updates(params)
	return ctx.IsErrorJson(result.Error)
}
