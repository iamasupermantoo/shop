package walletsService

import (
	"github.com/gomodule/redigo/redis"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/systemsService"
	"gofiber/app/module/views"
)

type UserBill struct {
}

func NewUserBill() *UserBill {
	return &UserBill{}
}

// GetUserBillName 获取账单名称
func (_UserBill *UserBill) GetUserBillName(billType int) string {
	for _, options := range walletsModel.WalletUserBillTypeList {
		if billType == options.Value.(int) {
			return options.Label
		}
	}
	return ""
}

// GetBalanceOptions 获取余额Options
func (_UserBill *UserBill) GetBalanceOptions() []*views.InputOptions {
	data := make([]*views.InputOptions, 0)
	for _, options := range walletsModel.WalletUserBillTypeList {
		if (options.Value.(int) < 99 && options.Value.(int) > 0) || (options.Value.(int) > -100 && options.Value.(int) < 0) {
			data = append(data, &views.InputOptions{Label: options.Field, Value: options.Value})
		}
	}
	return data
}

// GetAssetsOptions 获取资产Options
func (_UserBill *UserBill) GetAssetsOptions() []*views.InputOptions {
	data := make([]*views.InputOptions, 0)
	for _, options := range walletsModel.WalletUserBillTypeList {
		if options.Value.(int) > 99 || options.Value.(int) < -100 {
			data = append(data, &views.InputOptions{Label: options.Field, Value: options.Value})
		}
	}
	return data
}

// GetViewsOptions 获取视图Options
func (_UserBill *UserBill) GetViewsOptions(rdsConn redis.Conn, adminSettingId uint) []*views.InputOptions {
	optionsList := make([]*views.InputOptions, 0)
	translateService := systemsService.NewSystemTranslate(rdsConn, adminSettingId)
	for _, options := range walletsModel.WalletUserBillTypeList {
		optionsList = append(optionsList, &views.InputOptions{
			Label: translateService.GetRedisAdminTranslateLangField("zh-CN", options.Field),
			Value: options.Value,
		})
	}
	return optionsList
}
