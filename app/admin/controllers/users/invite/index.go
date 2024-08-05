package invite

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	AdminName string `json:"adminName"` // 管理账户
	UserName  string `json:"userName"`  // 用户账户
	Code      string `json:"code"`      // 邀请码
	Status    int    `json:"status"`    // 状态 -1禁用 10开启
	context.IndexParams
}

type invite struct {
	usersModel.Invite
	AdminInfo *adminsModel.AdminUser `gorm:"foreignKey:AdminId" json:"adminInfo"`
	UserInfo  *usersModel.User       `gorm:"foreignKey:UserId" json:"userInfo"`
}

// Index 管理列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*invite, 0)}
	database.Db.Model(&usersModel.Invite{}).Preload("AdminInfo").Preload("UserInfo").
		Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Scopes(scopes.NewScopes().
			FindModeIn("admin_id", &adminsModel.AdminUser{}, "id", "user_name = ?", params.AdminName).
			FindModeIn("user_id", &usersModel.User{}, "id", "user_name = ?", params.UserName).
			Eq("code", params.Code).
			Eq("status", params.Status).
			Between("updated_at", params.UpdatedAt).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
