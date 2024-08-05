package wholesale

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

type DetailsParams struct {
	ID uint `json:"id" validate:"required"` //		产品ID
}

// Details 批发中心产品详情
func Details(ctx *context.CustomCtx, params *DetailsParams) error {
	storeInfo := &shopsModel.Store{}
	result := database.Db.Model(storeInfo).Where("user_id = ?", ctx.UserId).Where("status = ?", shopsModel.StoreStatusActivate).Find(storeInfo)
	if result.Error != nil || storeInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("findError")
	}

	productInfo := &productsModel.Product{}
	database.Db.Model(productInfo).Where("id = ?", params.ID).Where("admin_id = ?", ctx.AdminSettingId).Find(productInfo)

	return ctx.SuccessJson(productInfo)
}
