package manage

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	UserName  string                  `json:"userName"`  //	管理名称
	NickName  string                  `json:"nickName"`  //	管理昵称
	Email     string                  `json:"email"`     //	邮箱地址
	Status    int                     `json:"status"`    //	状态
	Domains   string                  `json:"domains"`   //	域名
	Role      string                  `json:"role"`      //	角色
	ExpiredAt *scopes.RangeDatePicker `json:"expiredAt"` //	过期时间
	context.IndexParams
}

type adminUser struct {
	adminsModel.AdminUser
	ParentInfo *adminsModel.AdminUser      `gorm:"foreignKey:ID;references:ParentId" json:"parentInfo"`
	RoleInfo   *adminsModel.AuthAssignment `gorm:"foreignKey:AdminId;references:ID" json:"roleInfo"`
}

// Index 管理列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	// 	只允许查询下级管理
	data := &context.IndexData{Items: make([]*adminUser, 0)}
	//	过滤参数
	database.Db.Model(&adminsModel.AdminUser{}).Where("id IN ?", ctx.GetAdminChildIds()).
		Preload("ParentInfo").Preload("RoleInfo").
		Scopes(scopes.NewScopes().FindModeIn("id", &adminsModel.AuthAssignment{}, "admin_id", "name = ?", params.Role).
			Eq("user_name", params.UserName).
			Eq("nick_name", params.NickName).
			Eq("email", params.Email).
			Like("domains", "%"+params.Domains+"%").
			Eq("status", params.Status).
			Between("expired_at", params.ExpiredAt).
			Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	return ctx.SuccessJson(data)
}
