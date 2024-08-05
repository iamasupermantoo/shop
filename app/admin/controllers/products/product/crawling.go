package product

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
	CrawlInfo []CrawlInfo `json:"crawlInfo"`
}

type CrawlInfo struct {
	CategoryId uint   `json:"categoryId" validate:"required,gt=0"` //  类目ID
	Urls       string `json:"urls" validate:"required"`            //  产品URL
}

// Crawling 更新接口
func Crawling(ctx *context.CustomCtx, params *CrawlingParams) error {
	for _, v := range params.CrawlInfo {
		var adminId uint
		database.Db.Model(&productsModel.Category{}).Select("admin_id").Where("id = ?", v.CategoryId).Scan(&adminId)
		if adminId == 0 {
			continue
		}

		settingAdminId := adminsService.NewAdminUser(ctx.Rds, ctx.AdminSettingId).GetRedisAdminSettingId(adminId)
		crawl, err := crawling.NewCrawling()
		if err != nil {
			return ctx.ErrorJson(err.Error())
		}

		crawl.SetRequestHeaders(amazon.Headers()).
			SetOnHtml(amazon.ObtainDetails(settingAdminId, "body")).
			SetOnHtml(amazon.SearchClassLink("#search")).
			SetOnScraped(func(r *colly.Response) {
				if data := r.Ctx.GetAny("message"); data != nil {
					if p, ok := data.(crawling.ProductAttr); ok {

						tx := database.Db.Begin()
						_, err = productsService.NewProduct(tx).InsertCrawlingProduct(settingAdminId, v.CategoryId, &p)
						if err != nil {
							tx.Rollback()
							return
						}
						tx.Commit()
					}
				}
			})

		urlTmp := strings.Split(v.Urls, "\n")
		for _, url := range urlTmp {
			err = crawl.Run(url)
			if err != nil {
				zap.L().Error("Crawling", zap.Error(err))
			}
		}
		crawl.Wait()
		crawl.Clone()
	}

	return ctx.SuccessJsonOK()
}
