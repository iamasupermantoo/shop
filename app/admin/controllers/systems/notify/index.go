package notify

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	AdminName string `json:"adminName"` // 管理员名称
	UserName  string `json:"userName"`  // 用户名称
	Name      string `json:"name"`      // 标题
	Status    int    `json:"status"`    // 状态
	Mode      int    `json:"mode"`      // 模式 1后台 11前台
	context.IndexParams
}

type IndexData struct {
	systemsModel.Notify
	AdminInfo adminsModel.AdminUser `gorm:"foreignKey:AdminId" json:"adminInfo"`
	UserInfo  usersModel.User       `gorm:"foreignKey:UserId" json:"userInfo"`
}

// Index 通知列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*IndexData, 0)}

	//	过滤参数
	database.Db.Model(&systemsModel.Notify{}).
		Preload("AdminInfo").
		Preload("UserInfo").
		Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Scopes(scopes.NewScopes().
			FindModeIn("admin_id", &adminsModel.AdminUser{}, "id", "user_name = ?", params.AdminName).
			FindModeIn("user_id", &usersModel.User{}, "id", "user_name = ?", params.UserName).
			Eq("name", params.Name).
			Eq("status", params.Status).
			Eq("mode", params.Mode).
			Between("created_at", params.CreatedAt).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
