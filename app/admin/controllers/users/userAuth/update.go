package userAuth

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// UpdateParams 更新参数
type UpdateParams struct {
	ID       uint   `gorm:"-" validate:"required" json:"id"` //	ID
	RealName string `json:"realName"`                        // 真实姓名
	Number   string `json:"number"`                          // 卡号
	Photo1   string `json:"photo1"`                          // 证件照1
	Photo2   string `json:"photo2"`                          // 证件照2
	Photo3   string `json:"photo3"`                          // 证件照3
	Address  string `json:"address"`                         // 详细地址
	Type     int    `json:"type"`                            // 类型 1身份证
	Status   int    `json:"status"`                          // 状态 -1拒绝 10审核 20完成
}

// Update 更新接口
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	result := database.Db.Model(&usersModel.UserAuth{}).
		Where("id = ?", params.ID).Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Updates(params)
	return ctx.IsErrorJson(result.Error)
}
