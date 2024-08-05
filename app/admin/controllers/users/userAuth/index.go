package userAuth

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
	RealName  string `json:"realName"`  // 真实姓名
	Number    string `json:"number"`    // 卡号
	Address   string `json:"address"`   // 详细地址
	Type      int    `json:"type"`      // 类型 1身份证
	Status    int    `json:"status"`    // 状态 -1拒绝 10审核 20完成
	context.IndexParams
}

type userAuth struct {
	usersModel.UserAuth
	AdminInfo adminsModel.AdminUser `gorm:"foreignKey:AdminId;" json:"adminInfo"`
	UserInfo  usersModel.User       `gorm:"foreignKey:UserId;" json:"userInfo"`
}

// Index 管理列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*userAuth, 0)}
	database.Db.Model(&usersModel.UserAuth{}).Preload("AdminInfo").Preload("UserInfo").
		Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Scopes(scopes.NewScopes().
			FindModeIn("admin_id", &adminsModel.AdminUser{}, "id", "user_name = ?", params.AdminName).
			FindModeIn("user_id", &usersModel.User{}, "id", "user_name = ?", params.UserName).
			Eq("real_name", params.RealName).
			Eq("number", params.Number).
			Like("address", "%"+params.Address+"%").
			Eq("status", params.Status).
			Between("created_at", params.CreatedAt).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
