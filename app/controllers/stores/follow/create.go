package follow

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

type CreateParams struct {
	ID   uint `json:"id" validate:"required"`   // 产品ID
	Type int  `json:"type" validate:"required"` // 收藏类型  1收藏店铺 2收藏产品
}

// Create 收藏产品创建
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	adminIds := ctx.GetAdminChildIds()

	switch params.Type {
	case shopsModel.StoreFollowTypeCollectionProduct:
		// 收藏商品
		productInfo := &productsModel.Product{}
		database.Db.Model(productInfo).Where("id = ?", params.ID).Where("admin_id IN ?", adminIds).Find(productInfo)
		if productInfo.ID > 0 {
			_ = userFollowProduct(ctx.AdminId, ctx.UserId, productInfo)
		}
	case shopsModel.StoreFollowTypeConcernStore:
		// 关注店铺
		storeInfo := &shopsModel.Store{}
		database.Db.Model(storeInfo).Where("id = ?", params.ID).Where("admin_id IN ?", adminIds).Find(storeInfo)
		if storeInfo.ID > 0 {
			_ = userFollowShop(ctx.AdminId, ctx.UserId, storeInfo)
		}
	}

	return ctx.SuccessJsonOK()
}

func userFollowProduct(adminId, userId uint, productInfo *productsModel.Product) error {
	storeFollowInfo := &shopsModel.StoreFollow{}
	// 查询用户是否关注过商品
	result := database.Db.Where("user_id = ?", userId).
		Where("type = ?", shopsModel.StoreFollowTypeCollectionProduct).
		Where("product_id = ?", productInfo.ID).
		Find(&storeFollowInfo)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		database.Db.Create(&shopsModel.StoreFollow{
			AdminId:   adminId,
			UserId:    userId,
			StoreId:   productInfo.StoreId,
			ProductId: productInfo.ID,
			Type:      shopsModel.StoreFollowTypeCollectionProduct,
		})
	} else {
		// 之前该用户就关注了该产品，更新关注信息
		followStatus := shopsModel.StoreFollowStatusConcern
		if storeFollowInfo.Status == shopsModel.StoreFollowStatusConcern {
			followStatus = shopsModel.StoreFollowStatusCancels
		}
		database.Db.Model(&shopsModel.StoreFollow{}).Where("id = ?", storeFollowInfo.ID).Update("status", followStatus)
	}
	return nil
}

func userFollowShop(adminId, userId uint, storeInfo *shopsModel.Store) error {
	storeFollowInfo := &shopsModel.StoreFollow{}
	result := database.Db.Where("user_id = ?", userId).
		Where("type = ?", shopsModel.StoreFollowTypeConcernStore).
		Where("store_id = ?", storeInfo.ID).
		Find(&storeFollowInfo)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		database.Db.Create(&shopsModel.StoreFollow{
			AdminId: adminId,
			UserId:  userId,
			StoreId: storeInfo.ID,
			Type:    shopsModel.StoreFollowTypeConcernStore,
		})
	} else {
		followStatus := shopsModel.StoreFollowStatusConcern
		if storeFollowInfo.Status == shopsModel.StoreFollowStatusConcern {
			followStatus = shopsModel.StoreFollowStatusCancels
		}

		database.Db.Model(&shopsModel.StoreFollow{}).Where("id = ?", storeFollowInfo.ID).Update("status", followStatus)
	}
	return nil
}
