package store

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	AdminName string `json:"adminName"` // 管理账户
	UserName  string `json:"userName"`  // 用户账户
	Name      string `json:"name"`      // 店铺名称
	Contact   string `json:"contact"`   // 联系方式
	Keywords  string `json:"keywords"`  // 关键词
	Status    int    `json:"status"`    // 状态 -2删除 -1禁用 10激活 20关店
	context.IndexParams
}

type IndexData struct {
	shopsModel.Store
	AdminInfo adminsModel.AdminUser `json:"adminInfo" gorm:"foreignKey:AdminId;"`
	UserInfo  usersModel.User       `json:"userInfo" gorm:"foreignKey:UserId;"`
}

// Index 店铺列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*IndexData, 0)}
	database.Db.Model(&shopsModel.Store{}).
		Preload("AdminInfo").
		Preload("UserInfo").
		Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Scopes(scopes.NewScopes().
			FindModeIn("admin_id", &adminsModel.AdminUser{}, "id", "user_name = ?", params.AdminName).
			FindModeIn("user_id", &usersModel.User{}, "id", "user_name = ?", params.UserName).
			Like("name", "%"+params.Name+"%").
			Like("contact", params.Contact+"%").
			Like("keywords", "%"+params.Keywords+"%").
			Eq("status", params.Status).
			Between("updated_at", params.UpdatedAt).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
