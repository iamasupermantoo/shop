package store

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"time"
)

type homeData struct {
	OrderPending   int64           `json:"orderPending"`   // 未付款数
	OrderShipping  int64           `json:"orderShipping"`  // 待发货数
	OrderRefund    int64           `json:"orderRefund"`    // 待售后数
	OrderComment   int64           `json:"orderComment"`   // 待评论
	VisitorStatist statisticsInt   `json:"visitorStatist"` //	访问记录
	SalesStatist   statisticsFloat `json:"salesStatist"`   //	销售金额
	OrderStatist   statisticsInt   `json:"orderStatist"`   // 	订单数
	EarningStatist statisticsFloat `json:"earningStatist"` // 	收益金额
	FollowStatist  statisticsInt   `json:"followStatist"`  // 	店铺关注
	CollectStatist statisticsInt   `json:"collectStatist"` // 	商品收藏
	RefundStatist  statisticsInt   `json:"refundStatist"`  // 	售后订单数
}

type statisticsFloat struct {
	Today     float64 `json:"today"`     //	今日
	Yesterday float64 `json:"yesterday"` //	昨日
	Month     float64 `json:"month"`     //	本月
}

type statisticsInt struct {
	Today     int64 `json:"today"`     //	今日
	Yesterday int64 `json:"yesterday"` //	昨日
	Month     int64 `json:"month"`     //	本月
}

// Home 店铺首页信息
func Home(ctx *context.CustomCtx, params *context.NoRequestBody) error {
	data := &homeData{}

	storeInfo := &shopsModel.Store{}
	result := database.Db.Model(storeInfo).Where("user_id = ?", ctx.UserId).Where("status = ?", shopsModel.StoreStatusActivate).Find(storeInfo)
	if result.Error != nil || storeInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("findError")
	}

	// 未付款数量
	database.Db.Model(&shopsModel.ProductStoreOrder{}).Where("store_id = ?", storeInfo.ID).Where("status = ?", shopsModel.ProductStoreOrderStatusPending).Count(&data.OrderPending)

	// 待发货数量
	database.Db.Model(&shopsModel.ProductStoreOrder{}).Where("store_id = ?", storeInfo.ID).Where("status = ?", shopsModel.ProductStoreOrderStatusShipping).Count(&data.OrderShipping)

	// 待售后数量
	database.Db.Model(&shopsModel.StoreRefund{}).Where("store_id = ?", storeInfo.ID).Where("status = ?", shopsModel.StoreRefundStatusPending).Count(&data.OrderRefund)

	// 待评论数量
	database.Db.Model(&shopsModel.StoreComment{}).Where("store_id = ?", storeInfo.ID).Where("status = ?", shopsModel.StoreCommentsStatusPending).Count(&data.OrderComment)

	nowTime := time.Now()
	todayTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, time.Local)
	yesterdayTime := todayTime.Add(-24 * time.Hour)
	monthTime := todayTime.Add(-30 * 24 * time.Hour)

	// 今日访问量
	database.Db.Model(&usersModel.Access{}).Distinct("source_id").Where("type = ?", usersModel.AccessTypeStore).
		Where("created_at BETWEEN ? AND ?", todayTime, nowTime).
		Where("source_id = ?", storeInfo.ID).Count(&data.VisitorStatist.Today)
	// 昨日访问量
	database.Db.Model(&usersModel.Access{}).Distinct("source_id").Where("type = ?", usersModel.AccessTypeStore).
		Where("created_at BETWEEN ? AND ?", yesterdayTime, todayTime).
		Where("source_id = ?", storeInfo.ID).Count(&data.VisitorStatist.Yesterday)
	// 本月访问量
	database.Db.Model(&usersModel.Access{}).Distinct("source_id").Where("type = ?", usersModel.AccessTypeStore).
		Where("created_at BETWEEN ? AND ?", monthTime, nowTime).
		Where("source_id = ?", storeInfo.ID).Count(&data.VisitorStatist.Month)

	// 今日销售量
	database.Db.Model(&shopsModel.ProductStoreOrder{}).Select("IFNULL(SUM(final_money), 0)").Where("store_id = ?", storeInfo.ID).Where("status = ?", shopsModel.ProductStoreOrderStatusComplete).
		Where("created_at BETWEEN ? AND ?", todayTime, nowTime).Find(&data.SalesStatist.Today)
	// 昨日销售量
	database.Db.Model(&shopsModel.ProductStoreOrder{}).Select("IFNULL(SUM(final_money), 0)").Where("store_id = ?", storeInfo.ID).Where("status = ?", shopsModel.ProductStoreOrderStatusComplete).
		Where("created_at BETWEEN ? AND ?", yesterdayTime, todayTime).Find(&data.SalesStatist.Yesterday)
	// 本月销售量
	database.Db.Model(&shopsModel.ProductStoreOrder{}).Select("IFNULL(SUM(final_money), 0)").Where("store_id = ?", storeInfo.ID).Where("status = ?", shopsModel.ProductStoreOrderStatusComplete).
		Where("created_at BETWEEN ? AND ?", monthTime, nowTime).Find(&data.SalesStatist.Month)

	// 今日订单数
	database.Db.Model(&shopsModel.ProductStoreOrder{}).Where("store_id = ?", storeInfo.ID).
		Where("created_at BETWEEN ? AND ?", todayTime, nowTime).Count(&data.OrderStatist.Today)
	// 昨日订单数
	database.Db.Model(&shopsModel.ProductStoreOrder{}).Where("store_id = ?", storeInfo.ID).
		Where("created_at BETWEEN ? AND ?", yesterdayTime, todayTime).Count(&data.OrderStatist.Yesterday)
	// 本月订单数
	database.Db.Model(&shopsModel.ProductStoreOrder{}).Where("store_id = ?", storeInfo.ID).
		Where("created_at BETWEEN ? AND ?", monthTime, nowTime).Count(&data.OrderStatist.Month)

	// 今日收益
	database.Db.Model(&shopsModel.ProductStoreOrder{}).Select("IFNULL(SUM(earnings), 0)").Where("store_id = ?", storeInfo.ID).Where("status = ?", shopsModel.ProductStoreOrderStatusComplete).
		Where("created_at BETWEEN ? AND ?", todayTime, nowTime).Find(&data.EarningStatist.Today)
	// 昨日收益
	database.Db.Model(&shopsModel.ProductStoreOrder{}).Select("IFNULL(SUM(earnings), 0)").Where("store_id = ?", storeInfo.ID).Where("status = ?", shopsModel.ProductStoreOrderStatusComplete).
		Where("created_at BETWEEN ? AND ?", yesterdayTime, todayTime).Find(&data.EarningStatist.Yesterday)
	// 本月收益
	database.Db.Model(&shopsModel.ProductStoreOrder{}).Select("IFNULL(SUM(earnings), 0)").Where("store_id = ?", storeInfo.ID).Where("status = ?", shopsModel.ProductStoreOrderStatusComplete).
		Where("created_at BETWEEN ? AND ?", monthTime, nowTime).Find(&data.EarningStatist.Month)

	// 今日店铺关注
	database.Db.Model(&shopsModel.StoreFollow{}).Where("store_id = ?", storeInfo.ID).Where("type = ?", shopsModel.StoreFollowTypeConcernStore).Where("status = ?", shopsModel.StoreFollowStatusConcern).
		Where("created_at BETWEEN ? AND ?", todayTime, nowTime).Count(&data.FollowStatist.Today)
	// 昨日店铺关注
	database.Db.Model(&shopsModel.StoreFollow{}).Where("store_id = ?", storeInfo.ID).Where("type = ?", shopsModel.StoreFollowTypeConcernStore).Where("status = ?", shopsModel.StoreFollowStatusConcern).
		Where("created_at BETWEEN ? AND ?", yesterdayTime, todayTime).Count(&data.FollowStatist.Yesterday)
	// 本月店铺关注
	database.Db.Model(&shopsModel.StoreFollow{}).Where("store_id = ?", storeInfo.ID).Where("type = ?", shopsModel.StoreFollowTypeConcernStore).Where("status = ?", shopsModel.StoreFollowStatusConcern).
		Where("created_at BETWEEN ? AND ?", monthTime, nowTime).Count(&data.FollowStatist.Month)

	// 今日商品收藏
	database.Db.Model(&shopsModel.StoreFollow{}).Where("store_id = ?", storeInfo.ID).Where("type = ?", shopsModel.StoreFollowTypeCollectionProduct).Where("status = ?", shopsModel.StoreFollowStatusConcern).
		Where("created_at BETWEEN ? AND ?", todayTime, nowTime).Count(&data.CollectStatist.Today)
	// 昨日商品收藏
	database.Db.Model(&shopsModel.StoreFollow{}).Where("store_id = ?", storeInfo.ID).Where("type = ?", shopsModel.StoreFollowTypeCollectionProduct).Where("status = ?", shopsModel.StoreFollowStatusConcern).
		Where("created_at BETWEEN ? AND ?", yesterdayTime, todayTime).Count(&data.CollectStatist.Yesterday)
	// 本月商品收藏
	database.Db.Model(&shopsModel.StoreFollow{}).Where("store_id = ?", storeInfo.ID).Where("type = ?", shopsModel.StoreFollowTypeCollectionProduct).Where("status = ?", shopsModel.StoreFollowStatusConcern).
		Where("created_at BETWEEN ? AND ?", monthTime, nowTime).Count(&data.CollectStatist.Month)

	// 今日售后
	database.Db.Model(&shopsModel.StoreRefund{}).Where("store_id = ?", storeInfo.ID).
		Where("created_at BETWEEN ? AND ?", todayTime, nowTime).Count(&data.RefundStatist.Today)
	// 昨日售后
	database.Db.Model(&shopsModel.StoreRefund{}).Where("store_id = ?", storeInfo.ID).
		Where("created_at BETWEEN ? AND ?", yesterdayTime, todayTime).Count(&data.RefundStatist.Yesterday)
	// 本月售后
	database.Db.Model(&shopsModel.StoreRefund{}).Where("store_id = ?", storeInfo.ID).
		Where("created_at BETWEEN ? AND ?", monthTime, nowTime).Count(&data.RefundStatist.Month)

	return ctx.SuccessJson(data)
}
