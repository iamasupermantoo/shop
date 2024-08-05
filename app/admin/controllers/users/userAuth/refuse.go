package userAuth

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

type RefuseParams struct {
	ID   uint   `validate:"required" json:"id"`   //	ID
	Data string `validate:"required" json:"data"` //	拒绝理由
}

// Refuse 拒绝钱包订单
func Refuse(ctx *context.CustomCtx, params *RefuseParams) error {
	authInfo := &usersModel.UserAuth{}
	result := database.Db.Model(authInfo).Where("id = ?", params.ID).
		Where("status = ?", usersModel.UserAuthStatusActive).
		Where("admin_id IN ?", ctx.GetAdminChildIds()).Find(authInfo)
	if result.Error != nil || authInfo.ID == 0 {
		return ctx.ErrorJson("找不到可操作的认证")
	}

	return ctx.IsErrorJson(database.Db.Model(&usersModel.UserAuth{}).Where("id = ?", authInfo.ID).Updates(map[string]interface{}{
		"status": usersModel.UserAuthStatusRefuse,
		"data":   params.Data,
	}).Error)
}
