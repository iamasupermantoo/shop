package convert

import (
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

type InfoParams struct {
	AssetsId   uint `json:"assetsId"`   // 资产ID
	ToAssetsId uint `json:"toAssetsId"` // 转换资产ID
}

// Info 资金转换信息
func Info(ctx *context.CustomCtx, params *InfoParams) error {
	var assetsRate float64 = 1
	var toAssetsRate float64 = 1
	if params.AssetsId > 0 {
		assetsInfo := &walletsModel.WalletAssets{}
		result := database.Db.Model(assetsInfo).Where("id = ?", params.AssetsId).Where("admin_id = ?", ctx.AdminSettingId).Find(assetsInfo)
		if result.Error != nil {
			return ctx.ErrorJsonTranslate("abnormalOperation", "convertInfo.assetsInfo")
		}
		assetsRate = assetsInfo.Rate
	}

	if params.ToAssetsId > 0 {
		toAssetsInfo := &walletsModel.WalletAssets{}
		result := database.Db.Model(toAssetsInfo).Where("id = ?", params.ToAssetsId).Where("admin_id = ?", ctx.AdminSettingId).Find(toAssetsInfo)
		if result.Error != nil {
			return ctx.ErrorJsonTranslate("abnormalOperation", "convertInfo.toAssetsInfo")
		}
		toAssetsRate = toAssetsInfo.Rate
	}

	return ctx.SuccessJson(&infoData{
		Rate: assetsRate * 1 / toAssetsRate,
	})
}

type infoData struct {
	Rate float64 `json:"rate"` // 兑换汇率
}
