package category

import (
	"github.com/gocolly/colly"
	"go.uber.org/zap"
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/models/service/productsService"
	"gofiber/app/module/context"
	"gofiber/app/module/crawling"
	"gofiber/app/module/crawling/amazon"
	"gofiber/app/module/database"
	"strings"
)

// CrawlingParams 更新参数
type CrawlingParams struct {
	Id  uint   `json:"id"`
	Url string `json:"url"`
}

type UrlInfo struct {
	Id         uint   `json:"id"`
	CategoryId uint   `json:"categoryId"`
	ProductId  uint   `json:"productId"`
	Url        string `json:"url"`
}

func (*UrlInfo) TableName() string {
	return "urls"
}

// Crawling 更新接口
func Crawling(ctx *context.CustomCtx, params *CrawlingParams) error {
	var adminId uint
	database.Db.Model(&productsModel.Category{}).Select("admin_id").Where("id = ?", params.Id).Scan(&adminId)
	if adminId == 0 {
		return ctx.ErrorJson("not find admin")
	}

	settingAdminId := adminsService.NewAdminUser(ctx.Rds, ctx.AdminSettingId).GetRedisAdminSettingId(adminId)
	go func() {
		crawl, err := crawling.NewCrawling()
		if err != nil {
			return
		}

		crawl.SetRequestHeaders(amazon.Headers()).
			SetOnHtml(amazon.ObtainDetails(settingAdminId, "body")).
			SetOnHtml(amazon.SearchClassLink("#search")).
			SetOnScraped(func(r *colly.Response) {
				if data := r.Ctx.GetAny("message"); data != nil {
					if p, ok := data.(crawling.ProductAttr); ok {
						tx := database.Db.Begin()
						_, err = productsService.NewProduct(tx).InsertCrawlingProduct(settingAdminId, params.Id, &p)
						if err != nil {
							tx.Rollback()
							return
						}
						tx.Commit()
					}
				}

			})

		urlTmp := strings.Split(params.Url, "\n")
		for _, url := range urlTmp {
			err = crawl.Run(url)
			if err != nil {
				zap.L().Error("Crawling", zap.Error(err))
			}
		}
		crawl.Wait()
		crawl.Clone()
	}()

	return ctx.SuccessJsonOK()
}
