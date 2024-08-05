package wholesale

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

type StatusParams struct {
	Ids    []uint `json:"ids" validate:"required"`                 // 店铺产品ID
	Status int    `json:"status" validate:"omitempty,oneof=-1 10"` // 店铺产品上架下架
}

// Status 店铺产品上架下架
func Status(ctx *context.CustomCtx, params *StatusParams) error {
	storeInfo := &shopsModel.Store{}
	result := database.Db.Model(storeInfo).
		Where("user_id = ?", ctx.UserId).
		Where("status = ?", shopsModel.StoreStatusActivate).
		Find(storeInfo)
	if result.Error != nil || storeInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("findError")
	}

	database.Db.Model(&productsModel.Product{}).
		Where("store_id = ?", storeInfo.ID).
		Where("id IN ?", params.Ids).
		Update("status", params.Status)
	return ctx.SuccessJsonOK()
}
