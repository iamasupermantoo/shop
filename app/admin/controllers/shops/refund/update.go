package refund

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/types"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// UpdateParams 更新参数
type UpdateParams struct {
	ID     uint              `json:"id" gorm:"-" validate:"required"` //	ID
	Name   string            `json:"name"`                            // 申请理由
	Images types.GormStrings `json:"images"`                          // 申请凭证
	Status int               `json:"status"`                          // 售后状态 -2删除 10申请中 11处理中 12拒绝 13完成
}

// Update 更新接口
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	//	查询申请退款的用户是否存在
	refundInfo := &shopsModel.StoreRefund{}
	if result := database.Db.
		Where(params.ID).
		Take(refundInfo); result.Error != nil {
		return result.Error
	}

	if result := database.Db.
		Where(refundInfo.UserId).
		Take(&usersModel.User{}); result.Error != nil {
		return result.Error
	}

	result := database.Db.Model(&shopsModel.StoreRefund{}).
		Where("id = ?", params.ID).Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Updates(params)
	return ctx.IsErrorJson(result.Error)
}
