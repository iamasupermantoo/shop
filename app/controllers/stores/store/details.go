package userStore

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gorm.io/gorm"
)

type DetailsParams struct {
	ID uint `json:"id" validate:"required"` //	店铺ID
}

type storeDetails struct {
	shopsModel.Store
	FollowInfo shopsModel.StoreFollow `json:"followInfo" gorm:"foreignKey:StoreId"`
}

func (storeDetails) TableName() string {
	return "store"
}

// Details 店铺详情
func Details(ctx *context.CustomCtx, params *DetailsParams) error {
	storeInfo := &storeDetails{}
	database.Db.Model(&shopsModel.Store{}).Where("id = ?", params.ID).Where("status = ?", shopsModel.StoreStatusActivate).
		Preload("FollowInfo", func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id = ?", ctx.UserId).Where("type = ?", shopsModel.StoreFollowTypeConcernStore)
		}).Find(&storeInfo)

	return ctx.SuccessJson(storeInfo)
}
