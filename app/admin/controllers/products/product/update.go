package product

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/types"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gorm.io/gorm"
)

// UpdateParams 更新参数
type UpdateParams struct {
	ID         uint              `json:"id" gorm:"-" validate:"required"` //  ID
	CategoryId uint              `json:"categoryId"`                      //  类目ID
	AssetsId   uint              `json:"assetsId"`                        //  资产ID
	Name       string            `json:"name"`                            //  标题
	Images     types.GormStrings `json:"images"`                          //  图片组
	Money      float64           `json:"money"`                           //  金额
	Discount   float64           `json:"discount"`                        //  折扣
	Type       int               `json:"type"`                            //  类型1默认类型
	Sort       int               `json:"sort"`                            //  排序
	Status     int               `json:"status"`                          //  状态-1禁用 10启用
	Desc       string            `json:"desc"`                            //  描述
}

// Update 更新接口
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	err := database.Db.Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&productsModel.Product{}).
			Where("id = ?", params.ID).
			Where("admin_id IN ?", ctx.GetAdminChildIds()).Updates(params)
		if result.Error != nil {
			return ctx.ErrorJson("更新产品失败")
		}

		// 更新第一个sku 价格
		skuList := make([]*productsModel.ProductAttrsSku, 0)
		database.Db.Model(&productsModel.ProductAttrsSku{}).Where("product_id = ?", params.ID).Find(&skuList)
		for skuInfoIndex, skuInfo := range skuList {
			if skuInfoIndex == 0 {
				result = tx.Model(&productsModel.ProductAttrsSku{}).Where("id = ?", skuInfo.ID).Updates(&productsModel.ProductAttrsSku{
					Money: params.Money, Discount: params.Discount,
				})
				if result.Error != nil {
					return ctx.ErrorJson("更新产品规格失败")
				}
			}
		}
		return nil
	})
	if err != nil {
		return ctx.ErrorJson(err.Error())
	}

	return ctx.SuccessJsonOK()
}
