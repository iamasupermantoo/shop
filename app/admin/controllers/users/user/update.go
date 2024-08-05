package user

import (
	"gofiber/app/models/model/types"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// UpdateParams 更新参数
type UpdateParams struct {
	ID          uint                     `gorm:"-" validate:"required" json:"id"` //	ID
	AdminId     int                      `json:"adminId"`                         // 管理ID
	ParentId    int                      `json:"parentId"`                        // 上级用户ID
	NickName    string                   `json:"nickName"`                        // 昵称
	Email       string                   `json:"email"`                           // 邮箱
	Telephone   string                   `json:"telephone"`                       // 手机号码
	Avatar      string                   `json:"avatar"`                          // 头像
	Score       int                      `json:"score"`                           // 信用分
	Sex         int                      `json:"sex"`                             // 性别0未知 1男 2女
	Password    types.GormPasswordParams `json:"password"`                        // 密码
	SecurityKey types.GormPasswordParams `json:"securityKey"`                     // 密钥
	Type        int                      `json:"type"`                            // 类型 -1虚拟用户 1默认用户 10会员用户
	Status      int                      `json:"status"`                          // 状态 -1冻结 10激活
	Desc        string                   `json:"desc"`                            // 详情
}

// Update 更新接口
// func Update(ctx *fiber.Ctx) error {
// func Update(ctx *context.CustomCtx, bodyParams *UpdateParams) error {
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	return ctx.IsErrorJson(database.Db.Model(&usersModel.User{}).
		Where("id = ?", params.ID).Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Updates(params).Error)

}
