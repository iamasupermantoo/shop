package assets

import (
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"strconv"
	"time"
)

// Index 用户资产列表
func Index(ctx *context.CustomCtx, params *context.NoRequestBody) error {
	data := assetsIndex{}
	database.Db.Model(&walletsModel.WalletAssets{}).Where("admin_id = ?", ctx.AdminSettingId).
		Preload("UserAssets", database.Db.Where("user_id = ?", ctx.UserId)).Where("status = ?", walletsModel.WalletAssetsStatusActive).
		Find(&data.AssetsList)
	adminSettingCache := adminsService.NewAdminSetting(ctx.Rds, ctx.AdminSettingId)
	amountRateStr := adminSettingCache.GetRedisAdminSettingField("amountRate")
	amountRate, _ := strconv.ParseFloat(amountRateStr, 64)

	// 查询每个资产信息
	data.Echarts = &echarts{
		Category:   make([]string, 0),
		SeriesList: make([]*series, 0),
	}

	staTime := time.Now().Add(-14 * 24 * time.Hour)
	endTime := time.Now()
	userBillList := make([]*statisticsDays, 0)
	database.Db.Model(&walletsModel.WalletUserBill{}).Select("assets_id", "DATE_FORMAT(created_at, '%Y%m%d') as days", "IFNULL(SUM(IF(type>0,money,-ABS(money))), 0) as balance").
		Where("user_id = ?", ctx.UserId).Where("assets_id > 0").Where("created_at BETWEEN ? AND ?", staTime, endTime).Group("days,assets_id").Find(&userBillList)

	for assetsRangeIndex, assetsData := range data.AssetsList {
		data.SumMoney += assetsData.UserAssets.Money * assetsData.Rate
		data.SumAmount += assetsData.UserAssets.Money * assetsData.Rate * amountRate

		seriesInfo := &series{
			Name:   assetsData.Name,
			Type:   "line",
			Smooth: true,
			Data:   make([]any, 0),
		}

		n := -14
		for i := n; i <= 0; i++ {
			starTime := time.Now().AddDate(0, 0, i)

			// 名称
			if assetsRangeIndex == 0 {
				data.Echarts.Category = append(data.Echarts.Category, starTime.Format("01/02"))
			}

			// 获取用户账单
			var maxBalance float64
			for _, userBill := range userBillList {
				if userBill.Days == starTime.Format("20060102") && userBill.AssetsId == assetsData.ID {
					maxBalance = userBill.Balance
				}
			}

			seriesInfo.Data = append(seriesInfo.Data, maxBalance)
		}

		data.Echarts.SeriesList = append(data.Echarts.SeriesList, seriesInfo)
	}

	return ctx.SuccessJson(data)
}

type assetsIndex struct {
	SumMoney   float64       `json:"sumMoney"`   // 总资产
	SumAmount  float64       `json:"sumAmount"`  // 总金额
	Echarts    *echarts      `json:"echarts"`    // 图标信息
	AssetsList []*assetsInfo `json:"assetsList"` // 资产列表
}

type series struct {
	Name   string `json:"name"`   //	名称
	Type   string `json:"type"`   //	线类型
	Smooth bool   `json:"smooth"` // 	平滑
	Data   []any  `json:"data"`   //	数据
}

type echarts struct {
	Category   []string  `json:"category"` //	日期
	SeriesList []*series `json:"series"`   //	数据
}

type statisticsDays struct {
	Days     string  `json:"days"`     // 时间
	Balance  float64 `json:"balance"`  // 金额
	AssetsId uint    `json:"assetsId"` // 资产ID
}
