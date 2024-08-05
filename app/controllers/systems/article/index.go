package article

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/models/service/systemsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// Index 文章基础列表
func Index(ctx *context.CustomCtx, params *context.NoRequestBody) error {
	articleList := make([]*systemsModel.SystemArticleInfo, 0)
	database.Db.Model(&systemsModel.Article{}).Where("admin_id = ?", ctx.AdminSettingId).Where("type = ?", systemsModel.ArticleTypeDefault).Find(&articleList)

	sysCache := systemsService.NewSystemTranslate(ctx.Rds, ctx.AdminSettingId)
	for _, article := range articleList {
		article.Name = sysCache.GetRedisAdminTranslateLangField(ctx.Lang, article.Name)
		article.Content = sysCache.GetRedisAdminTranslateLangField(ctx.Lang, article.Content)
	}
	return ctx.SuccessJson(articleList)
}
