package category

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/service/systemsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/module/scopes"
)

type IndexParams struct {
	AdminName string `json:"adminName"` //  管理账户
	ParentId  int    `json:"parentId"`  //  分类上级ID
	Type      int    `json:"type"`      //  类型1默认类型
	Status    int    `json:"status"`    //  状态-1禁用 10启用
	context.IndexParams
}

type category struct {
	productsModel.Category
	AdminInfo  adminsModel.AdminUser  `gorm:"foreignKey:AdminId;" json:"adminInfo"`
	ParentInfo productsModel.Category `gorm:"foreignKey:ID;references:ParentId" json:"parentInfo"`
}

// Index 管理列表
func Index(ctx *context.CustomCtx, params *IndexParams) error {
	data := &context.IndexData{Items: make([]*category, 0)}

	//	过滤参数
	database.Db.Model(&productsModel.Category{}).
		Preload("AdminInfo").Preload("ParentInfo").
		Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Scopes(scopes.NewScopes().
			FindModeIn("admin_id", &adminsModel.AdminUser{}, "id", "user_name = ?", params.AdminName).
			Eq("parent_id", params.ParentId).
			Eq("type", params.Type).
			Eq("status", params.Status).
			Between("updated_at", params.UpdatedAt).Scopes()).
		Count(&data.Count).
		Scopes(params.Pagination.Scopes()).
		Find(&data.Items)

	// 翻译分类
	for _, categoryInfo := range data.Items.([]*category) {
		translateService := systemsService.NewSystemTranslate(ctx.Rds, categoryInfo.AdminId)
		categoryInfo.Name = categoryInfo.Name + "_" + translateService.GetRedisAdminTranslateLangField("zh-CN", categoryInfo.Name)
		categoryInfo.ParentInfo.Name = translateService.GetRedisAdminTranslateLangField("zh-CN", categoryInfo.ParentInfo.Name)
	}

	return ctx.SuccessJson(data)
}
