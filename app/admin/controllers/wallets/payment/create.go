package payment

import (
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// CreateParams 新增参数
type CreateParams struct {
	AssetsId uint   `json:"assetsId"`
	Type     int    `json:"type"`                       //	类型
	Mode     int    `json:"mode" validate:"required"`   // 模式
	Name     string `validate:"required" json:"name"`   // 名称
	Icon     string `validate:"required" json:"icon"`   // 图标
	Symbol   string `validate:"required" json:"symbol"` // 标识
}

// Create 新增接口
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	assetsInfo := &walletsModel.WalletAssets{}
	assetsType := params.Type
	if assetsType == 0 {
		assetsType = walletsModel.WalletAssetsTypeFiatCurrency
	}

	assetsAdminId := ctx.AdminId
	result := database.Db.Model(assetsInfo).Where("id = ?", params.AssetsId).Find(assetsInfo)
	if (result.Error != nil || assetsInfo.ID == 0) && (params.Mode != walletsModel.WalletPaymentModeDeposit && params.Mode != walletsModel.WalletPaymentModeWithdraw) {
		return ctx.ErrorJson("找到资产信息")
	}
	if assetsInfo.ID > 0 {
		assetsType = walletsModel.WalletAssetsTypeDigitalCurrency
		assetsAdminId = assetsInfo.AdminId
	}

	dataJson := make([]*walletsModel.WalletPaymentData, 0)
	switch params.Mode {
	case walletsModel.WalletPaymentModeDeposit, walletsModel.WalletPaymentModeAssetsDeposit:
		switch assetsType {
		case walletsModel.WalletAssetsTypeFiatCurrency:
			dataJson = []*walletsModel.WalletPaymentData{{Label: "银行名称", Field: "bankName", Value: "建设银行", IsShow: true}, {Label: "真实姓名", Field: "realName", Value: "隔壁老王", IsShow: true}, {Label: "银行卡号", Field: "bankNumber", Value: "888866665555", IsShow: true}, {Label: "银行代码", Field: "bankCode", Value: "xxx666", IsShow: true}}
		case walletsModel.WalletAssetsTypeDigitalCurrency:
			dataJson = []*walletsModel.WalletPaymentData{{Label: "公链名称", Field: "name", Value: "ETH", IsShow: true}, {Label: "Token", Field: "realName", Value: "USDT", IsShow: true}, {Label: "数字地址", Field: "number", Value: "0x1Bdd1742C5dEd48bAa7F5D71dba59628D3A3130c", IsShow: true}}
		case walletsModel.WalletPaymentTypeChannel:
			dataJson = []*walletsModel.WalletPaymentData{{Label: "渠道名称", Field: "name", Value: "xx渠道余额", IsShow: false}, {Label: "渠道标识", Field: "realName", Value: "es", IsShow: false}, {Label: "渠道ID", Field: "number", Value: "1", IsShow: true}}
		}
	case walletsModel.WalletPaymentModeWithdraw, walletsModel.WalletPaymentModeAssetsWithdraw:
		switch assetsType {
		case walletsModel.WalletAssetsTypeFiatCurrency:
			dataJson = []*walletsModel.WalletPaymentData{{Label: "bankName", Field: "name", IsShow: true}, {Label: "realName", Field: "realName", IsShow: true}, {Label: "bankNumber", Field: "number", IsShow: true}, {Label: "bankCode", Field: "code", IsShow: true}, {Label: "bankRemark", Field: "remark", IsShow: true}}
		case walletsModel.WalletAssetsTypeDigitalCurrency:
			dataJson = []*walletsModel.WalletPaymentData{{Label: "bankAddress", Field: "number", IsShow: true}}
		}
	}

	createInfo := &walletsModel.WalletPayment{
		AdminId:  assetsAdminId,
		AssetsId: params.AssetsId,
		Name:     params.Name,
		Symbol:   params.Symbol,
		Icon:     params.Icon,
		Type:     assetsType,
		Mode:     params.Mode,
		Data:     dataJson,
	}

	result = database.Db.Create(createInfo)
	if result.Error != nil {
		return ctx.ErrorJson("添加失败, 原因 => " + result.Error.Error())
	}

	return ctx.SuccessJsonOK()
}
