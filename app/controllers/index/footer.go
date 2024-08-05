package index

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/models/service/systemsService"
	"gofiber/app/module/context"
)

// Footer 脚部接口
func Footer(ctx *context.CustomCtx, params *context.NoRequestBody) error {
	articleService := systemsService.NewSystemArticle(ctx.Rds, ctx.AdminSettingId)

	data := []*footerInfo{
		{Label: "about", Items: articleService.GetArticleList(systemsModel.ArticleTypeAbout)},
		{Label: "product", Items: articleService.GetArticleList(systemsModel.ArticleTypeProduct)},
		{Label: "service", Items: articleService.GetArticleList(systemsModel.ArticleTypeService)},
		{Label: "helper", Items: articleService.GetArticleList(systemsModel.ArticleTypeHelpers)},
		{Label: "social", Items: articleService.GetArticleList(systemsModel.ArticleTypeSocial)},
	}
	return ctx.SuccessJson(data)
}

type footerInfo struct {
	Label string                            `json:"label"`
	Items []*systemsModel.SystemArticleInfo `json:"items"`
}
