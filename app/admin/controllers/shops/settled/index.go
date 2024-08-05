package settled

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
	Type      int    `json:"type"`      // 类型 1营业执照
	Name      string `json:"name"`      // 证件名字
	Number    string `json:"number"`    // 证件号
	Contact   string `json:"contact"`   // 联系方式
	Status    int    `json:"status"`    // 状态  -1拒绝 10审核中 20 通过
	context.IndexParams
}

type IndexData struct {
	shopsModel.StoreSettled
	AdminInfo adminsModel.AdminUser `json:"adminInfo" gorm:"foreignKey:AdminId;"`
	UserInfo  usersModel.User       `json:"userInfo" gorm:"foreignKey:UserId;"`
}

// Index 管理列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*IndexData, 0)}
	database.Db.Model(&shopsModel.StoreSettled{}).
		Preload("AdminInfo").
		Preload("UserInfo").
		Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Scopes(scopes.NewScopes().
			FindModeIn("admin_id", &adminsModel.AdminUser{}, "id", "user_name = ?", params.AdminName).
			FindModeIn("user_id", &usersModel.User{}, "id", "user_name = ?", params.UserName).
			Like("name", params.Name+"%").
			Like("contact", params.Contact+"%").
			Like("number", params.Number+"%").
			Eq("status", params.Status).
			Eq("type", params.Type).
			Between("updated_at", params.UpdatedAt).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
