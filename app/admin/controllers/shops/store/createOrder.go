package store

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/service/shopsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/utils"
	"gorm.io/gorm"
)

// CreateOrderParams 新增参数
type CreateOrderParams struct {
	ID            uint `json:"id"`            // 店铺Id
	ProductNumber int  `json:"productNumber"` // 产品数量
	BuyNumber     int  `json:"buyNumber"`     // 单量
}

type userInfo struct {
	usersModel.User
	Address *shopsModel.ShippingAddress ` json:"address" gorm:"foreignKey:UserId"` // 用户地址
}

type ProductAndSkuInfo struct {
	productsModel.ProductAttrsSku
	ProductInfo *productsModel.Product `json:"productInfo" gorm:"foreignKey:ProductId;references:ID"`
}

// CreateOrder 创建产品订单
func CreateOrder(ctx *context.CustomCtx, params *CreateOrderParams) error {
	userList := make([]*userInfo, 0)
	database.Db.Model(&usersModel.User{}).Preload("Address").
		Where("admin_id = ?", ctx.AdminId).
		Where("type = ?", usersModel.UserTypeVirtual).
		Find(&userList)
	if len(userList) == 0 {
		return ctx.ErrorJson("没有找到虚拟用户！！！")
	}

	var skuIds []uint
	database.Db.Raw("select pas.id from product as p left join product_attrs_sku as pas on p.id = pas.product_id where p.store_id = ?", params.ID).Scan(&skuIds)
	if len(skuIds) == 0 {
		return ctx.ErrorJson("没有找到产品！！！")
	}

	err := database.Db.Transaction(func(tx *gorm.DB) error {
		indexs := utils.NewRandom().IntArray(params.ProductNumber, 0, len(skuIds)-1)
		for _, index := range indexs {
			productAttrsSku := productsModel.ProductAttrsSku{}
			database.Db.Where("id = ?", skuIds[index]).Find(&productAttrsSku)
			skuParams := []*shopsService.SkuParam{
				{SkuInfo: &productAttrsSku, Nums: params.BuyNumber},
			}

			index = utils.NewRandom().Intn(0, len(userList)-1)
			err := shopsService.NewStoreOrder(tx).CreatOrder(ctx.Rds, 0, params.ID, &userList[index].User, userList[index].Address, skuParams)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return ctx.ErrorJson(err.Error())
	}
	return ctx.SuccessJsonOK()
}
