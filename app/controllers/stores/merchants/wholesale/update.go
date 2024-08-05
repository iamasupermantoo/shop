package wholesale

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/types"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

type UpdateParams struct {
	ID     uint              `json:"id" validate:"required,gt=0"` // 产品ID
	Images types.GormStrings `json:"images"`                      // 图片
	Name   string            `json:"name"`                        // 商品名称
	Desc   string            `json:"desc"`                        // 商品描述
}

// Update 更新店铺商品
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	storeInfo := &shopsModel.Store{}
	result := database.Db.Model(storeInfo).
		Where("user_id = ?", ctx.UserId).
		Where("status = ?", shopsModel.StoreStatusActivate).
		Find(storeInfo)
	if result.Error != nil || storeInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("findError")
	}

	productInfo := &productsModel.Product{}
	result = database.Db.Model(productInfo).Where("store_id = ?", storeInfo.ID).Find(productInfo)
	if result.Error != nil || productInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("findError")
	}

	database.Db.Where("id = ?", params.ID).Updates(&productsModel.Product{
		Images: params.Images, Name: params.Name, Desc: params.Desc,
	})
	return ctx.SuccessJsonOK()
}
