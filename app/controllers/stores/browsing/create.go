package browsing

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"time"
)

type CreateParams struct {
	Id uint `json:"id" validate:"required"` // 商品Id
}

// Create 商品浏览记录
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	productInfo := &productsModel.Product{}
	result := database.Db.Model(productInfo).Where("id = ?", params.Id).Where("status = ?", productsModel.ProductStatusActive).Find(productInfo)
	if result.RowsAffected > 0 {
		nowTime := time.Now()
		todayTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, time.Local)
		var browsingNums int64
		database.Db.Model(&productsModel.ProductBrowsing{}).
			Where("product_id = ?", productInfo.ID).
			Where("created_at BETWEEN ? AND ?", todayTime, nowTime).
			Count(&browsingNums)
		if browsingNums == 0 {
			database.Db.Create(&productsModel.ProductBrowsing{
				AdminId:   ctx.AdminId,
				UserId:    ctx.UserId,
				StoreId:   productInfo.StoreId,
				ProductId: productInfo.ID,
			})
		}
	}

	return ctx.SuccessJsonOK()
}
