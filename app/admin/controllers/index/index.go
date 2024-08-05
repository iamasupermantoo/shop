package index

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"time"
)

func Index(ctx *context.CustomCtx, bodyParams *context.NoRequestBody) error {
	//设置时间
	nowTime := time.Now()
	yesterdayTime := time.Now().AddDate(0, 0, -1)
	nowStarTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, time.Local)
	yesterdayStarTime := time.Date(yesterdayTime.Year(), yesterdayTime.Month(), yesterdayTime.Day(), 0, 0, 0, 0, time.Local)

	// 活跃人数 今日人数｜昨日人数｜总人数
	adminChildrenIds := ctx.GetAdminChildIds()
	var visitorToday, visitorYesterday, visitorSum int64
	if result := database.Db.Model(&usersModel.Access{}).
		Select("COUNT(DISTINCT ip)").
		Where("admin_id IN ?", adminChildrenIds).
		Where("route = ?", "/").
		Where("created_at BETWEEN ? AND ?", nowStarTime, nowTime).Count(&visitorToday); result.Error != nil {
		return ctx.ErrorJson(result.Error.Error())
	}

	if result := database.Db.Model(&usersModel.Access{}).
		Select("COUNT(DISTINCT ip)").
		Where("admin_id IN ?", adminChildrenIds).
		Where("route = ?", "/").
		Where("created_at BETWEEN ? AND ?", yesterdayStarTime, nowStarTime).Count(&visitorYesterday); result.Error != nil {
		return ctx.ErrorJson(result.Error.Error())
	}

	if result := database.Db.Model(&usersModel.Access{}).
		Select("COUNT(DISTINCT ip)").
		Where("route = ?", "/").
		Where("admin_id IN ?", adminChildrenIds).Count(&visitorSum); result.Error != nil {
		return ctx.ErrorJson(result.Error.Error())
	}

	// 用户人数 今日人数｜昨日人数｜总人数
	var userToday, userYesterday, userSum int64
	if result := database.Db.Model(&usersModel.User{}).
		Where("admin_id IN ?", adminChildrenIds).
		Where("type > ?", usersModel.UserTypeVirtual).
		Where("created_at BETWEEN ? AND ?", nowStarTime, nowTime).Count(&userToday); result.Error != nil {
		return ctx.ErrorJson(result.Error.Error())
	}

	if result := database.Db.Model(&usersModel.User{}).
		Where("admin_id IN ?", adminChildrenIds).
		Where("type > ?", usersModel.UserTypeVirtual).
		Where("created_at BETWEEN ? AND ?", yesterdayStarTime, nowStarTime).Count(&userYesterday); result.Error != nil {
		return ctx.ErrorJson(result.Error.Error())
	}

	if result := database.Db.Model(&usersModel.User{}).
		Where("type > ?", usersModel.UserTypeVirtual).
		Where("admin_id IN ?", adminChildrenIds).Count(&userSum); result.Error != nil {
		return ctx.ErrorJson(result.Error.Error())
	}

	//	今日充值｜昨日充值｜总充值
	var depositToday, depositYesterday, depositSum float64
	if result := database.Db.Model(&walletsModel.WalletUserOrder{}).
		Select("IFNULL(SUM(money),0)").
		Where("admin_id IN ?", adminChildrenIds).
		Where("type = ?", walletsModel.WalletUserOrderTypeDeposit).
		Where("status = ?", walletsModel.WalletUserOrderStatusComplete).
		Where("created_at BETWEEN ? AND ?", nowStarTime, nowTime).
		Find(&depositToday); result.Error != nil {
		return ctx.ErrorJson(result.Error.Error())
	}

	if result := database.Db.Model(&walletsModel.WalletUserOrder{}).
		Select("IFNULL(SUM(money),0)").
		Where("admin_id IN ?", adminChildrenIds).
		Where("type = ?", walletsModel.WalletUserOrderTypeDeposit).
		Where("status = ?", walletsModel.WalletUserOrderStatusComplete).
		Where("created_at BETWEEN ? AND ?", yesterdayStarTime, nowStarTime).
		Find(&depositYesterday); result.Error != nil {
		return ctx.ErrorJson(result.Error.Error())
	}

	if result := database.Db.Model(&walletsModel.WalletUserOrder{}).
		Select("IFNULL(SUM(money),0)").
		Where("admin_id IN ?", adminChildrenIds).
		Where("type = ?", walletsModel.WalletUserOrderTypeDeposit).
		Where("status = ?", walletsModel.WalletUserOrderStatusComplete).
		Find(&depositSum); result.Error != nil {
		return ctx.ErrorJson(result.Error.Error())
	}

	// 今日提现｜昨日提现｜总提现
	var withdrawToday, withdrawYesterday, withdrawSum float64

	if result := database.Db.Model(&walletsModel.WalletUserOrder{}).
		Select("IFNULL(SUM(money),0)").
		Where("admin_id IN ?", adminChildrenIds).
		Where("type = ?", walletsModel.WalletUserOrderTypeWithdraw).
		Where("status = ?", walletsModel.WalletUserOrderStatusComplete).
		Where("created_at BETWEEN ? AND ?", nowStarTime, nowTime).
		Find(&withdrawToday); result.Error != nil {
		return ctx.ErrorJson(result.Error.Error())
	}

	database.Db.Model(&walletsModel.WalletUserOrder{}).
		Select("IFNULL(SUM(money),0)").
		Where("admin_id IN ?", adminChildrenIds).
		Where("type = ?", walletsModel.WalletUserOrderTypeWithdraw).
		Where("status = ?", walletsModel.WalletUserOrderStatusComplete).
		Where("created_at BETWEEN ? AND ?", yesterdayStarTime, nowStarTime).
		Find(&withdrawYesterday)

	if result := database.Db.Model(&walletsModel.WalletUserOrder{}).
		Select("IFNULL(SUM(money),0)").
		Where("admin_id IN ?", adminChildrenIds).
		Where("type = ?", walletsModel.WalletUserOrderTypeWithdraw).
		Where("status = ?", walletsModel.WalletUserOrderStatusComplete).
		Find(&withdrawSum); result.Error != nil {
		return ctx.ErrorJson(result.Error.Error())
	}

	var category []string
	var visitorNumList []any
	var userNumList []any
	var depositNumList []any
	var withdrawNumList []any

	for i := -14; i <= 0; i++ {
		nowTimeTmp := time.Now().AddDate(0, 0, i)
		sourceTime := time.Date(nowTimeTmp.Year(), nowTimeTmp.Month(), nowTimeTmp.Day(), 0, 0, 0, 0, time.Local)
		staTime := sourceTime
		endTime := staTime.Add(24 * time.Hour)
		category = append(category, sourceTime.Format("01/02"))

		// 访客数
		var visitorNum int64
		if result := database.Db.Model(&usersModel.Access{}).Select("COUNT(DISTINCT ip)").
			Where("admin_id IN ?", adminChildrenIds).
			Where("route = ?", "/").
			Where("created_at BETWEEN ? AND ?", staTime, endTime).Count(&visitorNum); result.Error != nil {
			return ctx.ErrorJson(result.Error.Error())
		}
		visitorNumList = append(visitorNumList, visitorNum)

		// 用户数
		var userNum int64
		if result := database.Db.Model(&usersModel.User{}).
			Where("admin_id IN ?", adminChildrenIds).
			Where("type > ?", usersModel.UserTypeVirtual).
			Where("created_at BETWEEN ? AND ?", staTime, endTime).Count(&userNum); result.Error != nil {
			return ctx.ErrorJson(result.Error.Error())
		}
		userNumList = append(userNumList, userNum)

		// 充值量
		var depositNum float64
		if result := database.Db.Model(&walletsModel.WalletUserOrder{}).
			Select("IFNULL(SUM(money),0)").
			Where("admin_id IN ?", adminChildrenIds).
			Where("type = ?", walletsModel.WalletUserOrderTypeDeposit).
			Where("status = ?", walletsModel.WalletUserOrderStatusComplete).
			Where("created_at BETWEEN ? AND ?", staTime, endTime).
			Find(&depositNum); result.Error != nil {
			return ctx.ErrorJson(result.Error.Error())
		}
		depositNumList = append(depositNumList, depositNum)

		// 提现量
		var withdrawNum float64
		if result := database.Db.Model(&walletsModel.WalletUserOrder{}).
			Select("IFNULL(SUM(money),0)").
			Where("admin_id IN ?", adminChildrenIds).
			Where("type = ?", walletsModel.WalletUserOrderTypeWithdraw).
			Where("status = ?", walletsModel.WalletUserOrderStatusComplete).
			Where("created_at BETWEEN ? AND ?", staTime, endTime).
			Find(&withdrawNum); result.Error != nil {
			return ctx.ErrorJson(result.Error.Error())
		}
		withdrawNumList = append(withdrawNumList, withdrawNum)
	}

	data := &indexPageData{
		Items: [][]*statistics{
			{
				&statistics{
					Name:      "访客数",
					Icon:      "sym_o_person_pin_circle",
					Color:     "bg-primary",
					Today:     visitorToday,
					Yesterday: visitorYesterday,
					Total:     visitorSum,
				},
				&statistics{
					Name:      "用户数",
					Icon:      "sym_o_person_add",
					Color:     "bg-secondary",
					Today:     userToday,
					Yesterday: userYesterday,
					Total:     userSum,
				},
				&statistics{
					Name:      "充值量",
					Icon:      "sym_o_credit_card",
					Color:     "bg-accent",
					Today:     depositToday,
					Yesterday: depositYesterday,
					Total:     depositSum,
				},
				&statistics{
					Name:      "提现量",
					Icon:      "sym_o_payments",
					Color:     "bg-dark",
					Today:     withdrawToday,
					Yesterday: withdrawYesterday,
					Total:     withdrawSum,
				},
			},
		},
		EchartsList: &echarts{
			Category: category,
			SeriesList: []*series{
				{
					Name:   "访客数",
					Type:   "line",
					Smooth: true,
					Data:   visitorNumList,
				}, {
					Name:   "用户数",
					Type:   "line",
					Smooth: true,
					Data:   userNumList,
				}, {
					Name:   "充值量",
					Type:   "line",
					Smooth: true,
					Data:   depositNumList,
				}, {
					Name:   "提现量",
					Type:   "line",
					Smooth: true,
					Data:   withdrawNumList,
				},
			},
		},
	}

	return ctx.SuccessJson(data)
}

type statistics struct {
	Name      string `json:"name"`      //	名称
	Icon      string `json:"icon"`      //	图标
	Color     string `json:"color"`     //	背景颜色
	Today     any    `json:"today"`     //	今日
	Yesterday any    `json:"yesterday"` //	昨日
	Total     any    `json:"total"`     //	总数
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

// indexPageData 首页信息数据
type indexPageData struct {
	Items       [][]*statistics `json:"items"`   //	图标信息
	EchartsList *echarts        `json:"echarts"` //	图标数据
}
