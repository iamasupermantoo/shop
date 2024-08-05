package userAuth

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gorm.io/gorm"
)

type AgreeParams struct {
	ID uint `validate:"required" json:"id"` //	ID
}

// Agree 同意钱包订单操作
func Agree(ctx *context.CustomCtx, params *AgreeParams) error {
	authInfo := &usersModel.UserAuth{}
	result := database.Db.Model(authInfo).Where("id = ?", params.ID).
		Where("status = ?", usersModel.UserAuthStatusActive).
		Where("admin_id IN ?", ctx.GetAdminChildIds()).Find(authInfo)
	if result.Error != nil || authInfo.ID == 0 {
		return ctx.ErrorJson("找不到当前认证数据")
	}

	return ctx.IsErrorJson(database.Db.Transaction(func(tx *gorm.DB) error {
		result = tx.Model(&usersModel.UserAuth{}).Where("id = ?", authInfo.ID).Update("status", usersModel.UserAuthStatusComplete)
		return result.Error
	}))
}
