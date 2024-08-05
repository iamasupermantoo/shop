package wholesale

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/service/productsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gorm.io/gorm"
)

type ShelvesParams struct {
	Ids []uint `json:"ids" validate:"required"` //	批发商品Ids
}

// Shelves 批发中心上架产品
func Shelves(ctx *context.CustomCtx, params *ShelvesParams) error {
	storeInfo := &shopsModel.Store{}
	result := database.Db.Model(storeInfo).Where("user_id = ?", ctx.UserId).Where("status = ?", shopsModel.StoreStatusActivate).Find(storeInfo)
	if result.Error != nil || storeInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("findError")
	}

	err := database.Db.Transaction(func(tx *gorm.DB) error {
		for _, id := range params.Ids {
			wholesaleInfo := &productsModel.Product{}
			database.Db.Model(wholesaleInfo).Where("id = ?", id).Where("admin_id = ?", ctx.AdminSettingId).Find(wholesaleInfo)

			if wholesaleInfo.ID > 0 {
				storeProductInfo := &productsModel.Product{}
				result = database.Db.Model(storeProductInfo).Where("parent_id = ?", wholesaleInfo.ID).Where("store_id = ?", storeInfo.ID).Find(storeProductInfo)
				if result.Error == nil && storeProductInfo.ID == 0 {
					newProductInfo := &productsModel.Product{
						AdminId:    ctx.AdminId,
						ParentId:   wholesaleInfo.ID,
						StoreId:    storeInfo.ID,
						CategoryId: wholesaleInfo.CategoryId,
						AssetsId:   wholesaleInfo.AssetsId,
						Name:       wholesaleInfo.Name,
						Images:     wholesaleInfo.Images,
						Discount:   wholesaleInfo.Discount,
						Money:      wholesaleInfo.Money,
						Type:       productsModel.ProductTypeDefault,
						Desc:       wholesaleInfo.Desc,
					}
					result = tx.Create(newProductInfo)
					if result.Error != nil {
						return ctx.ErrorJsonTranslate("abnormalOperation")
					}

					// 产品的sku信息插入
					err := productsService.NewProduct(tx).InserterProductAttrs(productsService.ProductAttrsSkuWholesale, id, newProductInfo.ID)
					if err != nil {
						return err
					}
				}
			}
		}
		return nil
	})
	if err != nil {
		return ctx.ErrorJsonTranslate(err.Error())
	}

	return ctx.SuccessJsonOK()
}
