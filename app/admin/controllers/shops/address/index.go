package address

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	AdminName string `json:"adminName"` // 管理名
	UserName  string `json:"userName"`  // 用户名
	Name      string `json:"name"`      // 收件人名称
	Contact   string `json:"contact"`   // 联系方式
	City      string `json:"city"`      // 国家城市
	Address   string `json:"address"`   // 详细地址
	Type      int    `json:"type"`      // 类型 1收货地址 2发货地址
	Status    int    `json:"status"`    // 状态 -2删除 -1禁用 10激活
	IsShow    int    `json:"isShow"`    // 1不默认 2默认
	context.IndexParams
}

type IndexData struct {
	shopsModel.ShippingAddress
	AdminInfo adminsModel.AdminUser `json:"adminInfo" gorm:"foreignKey:AdminId;"`
	UserInfo  usersModel.User       `json:"userInfo" gorm:"foreignKey:UserId;"`
}

// Index 管理列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*IndexData, 0)}

	database.Db.Model(&shopsModel.ShippingAddress{}).
		Preload("AdminInfo").
		Preload("UserInfo").
		Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Scopes(scopes.NewScopes().
			FindModeIn("admin_id", &adminsModel.AdminUser{}, "id", "user_name = ?", params.AdminName).
			FindModeIn("user_id", &usersModel.User{}, "id", "user_name = ?", params.UserName).
			Like("name", params.Name+"%").
			Like("contact", params.Contact+"%").
			Like("city", params.City+"%").
			Like("address", params.Address+"%").
			Eq("type", params.Type).
			Eq("status", params.Status).
			Eq("is_show", params.IsShow).
			Between("updated_at", params.UpdatedAt).
			Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
