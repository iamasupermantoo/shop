package user

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
	"gofiber/app/module/views"
)

type IndexParams struct {
	AdminName  string `json:"adminName"`  // 管理账户
	ParentName string `json:"parentName"` // 上级账户
	UserName   string `json:"userName"`   // 用户名
	NickName   string `json:"nickName"`   // 昵称
	Email      string `json:"email"`      // 邮箱
	Telephone  string `json:"telephone"`  // 手机号码
	Type       int    `json:"type"`       // 类型 -1虚拟用户 1默认用户 10会员用户
	Status     int    `json:"status"`     // 状态 -1冻结 10激活
	context.IndexParams
}

type user struct {
	usersModel.User
	AdminInfo   adminsModel.AdminUser `gorm:"foreignKey:AdminId;" json:"adminInfo"`
	ParentInfo  usersModel.User       `gorm:"foreignKey:ID;references:ParentId" json:"parentInfo"`
	SettingList []usersModel.Setting  `gorm:"foreignKey:UserId" json:"settingList"`
	SettingJson map[string]any        `gorm:"-" json:"settingJson"`
}

// Index 管理列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*user, 0)}
	database.Db.Model(&usersModel.User{}).Preload("AdminInfo").Preload("ParentInfo").Preload("SettingList").
		Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Scopes(scopes.NewScopes().
			FindModeIn("admin_id IN ?", &adminsModel.AdminUser{}, "id", "user_name = ?", params.AdminName).
			FindModeIn("parent_id IN ?", &usersModel.User{}, "parent_id", "user_name = ?", params.ParentName).
			Eq("user_name", params.UserName).
			Eq("nick_name", params.NickName).
			Eq("email", params.Email).
			Eq("telephone", params.Telephone).
			Eq("type", params.Type).
			Eq("status", params.Status).
			Between("created_at", params.CreatedAt).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	userSettingList := make([]*usersModel.Setting, 0)
	database.Db.Where("admin_id = ?", 0).Find(&userSettingList)

	for _, v := range data.Items.([]*user) {
		v.SettingJson = map[string]interface{}{}
		// 基础配置模版值
		for _, settingInfo := range userSettingList {
			v.SettingJson[settingInfo.Field] = views.InputViewsStringToData(settingInfo.Type, settingInfo.Value)
		}

		// 用户已设置的值
		for _, settingInfo := range v.SettingList {
			v.SettingJson[settingInfo.Field] = views.InputViewsStringToData(settingInfo.Type, settingInfo.Value)
		}
	}

	return ctx.SuccessJson(data)
}
