package article

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/models/service/systemsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

type DetailsParams struct {
	ID int `json:"id" validate:"required"` //文章ID
}

// Details 文章详情
func Details(ctx *context.CustomCtx, params *DetailsParams) error {
	data := systemsModel.SystemArticleInfo{}
	database.Db.Model(&systemsModel.Article{}).Where("id = ?", params.ID).Where("admin_id = ?", ctx.AdminSettingId).Find(&data)

	// 翻译文章详情
	sysCache := systemsService.NewSystemTranslate(ctx.Rds, ctx.AdminSettingId)
	data.Name = sysCache.GetRedisAdminTranslateLangField(ctx.Lang, data.Name)
	data.Content = sysCache.GetRedisAdminTranslateLangField(ctx.Lang, data.Content)
	return ctx.SuccessJson(data)
}
