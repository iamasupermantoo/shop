package walletsService

import (
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/module/database"
	"gofiber/app/module/views"
	"strconv"
)

type WalletsAssets struct {
}

func NewWalletsAssets() *WalletsAssets {
	return &WalletsAssets{}
}

// AdminAssetsOptions 获取管理下级资产Options
func (_WalletsAssets *WalletsAssets) AdminAssetsOptions(adminIds []uint) []*views.InputOptions {
	assetList := make([]*walletsModel.WalletAssets, 0)
	database.Db.Model(&walletsModel.WalletAssets{}).Where("admin_id IN ?", adminIds).Find(&assetList)

	data := make([]*views.InputOptions, 0)
	for _, assets := range assetList {
		data = append(data, &views.InputOptions{
			Label: assets.Name + "." + strconv.Itoa(int(assets.AdminId)),
			Value: assets.ID,
		})
	}
	return data
}

// GetWithdrawOptions 获取管理提现Options
func (_WalletsAssets *WalletsAssets) GetWithdrawOptions(adminIds []uint) []*views.InputOptions {
	options := _WalletsAssets.GetBalanceWithdrawOptions(adminIds)
	options = append(options, _WalletsAssets.GetAssetsWithdrawOptions(adminIds)...)
	return options
}

// GetAssetsWithdrawOptions 获取管理资产提现Options
func (_WalletsAssets *WalletsAssets) GetAssetsWithdrawOptions(adminIds []uint) []*views.InputOptions {
	paymentList := make([]*walletsModel.WalletPayment, 0)
	database.Db.Model(&walletsModel.WalletPayment{}).Where("admin_id IN ?", adminIds).
		Where("mode = ?", walletsModel.WalletPaymentModeAssetsWithdraw).Find(&paymentList)

	data := make([]*views.InputOptions, 0)
	for _, payment := range paymentList {
		data = append(data, &views.InputOptions{
			Label: payment.Name + "." + strconv.Itoa(int(payment.AdminId)),
			Value: payment.ID,
		})
	}
	return data
}

// GetBalanceWithdrawOptions 获取管理余额提现Options
func (_WalletsAssets *WalletsAssets) GetBalanceWithdrawOptions(adminIds []uint) []*views.InputOptions {
	paymentList := make([]*walletsModel.WalletPayment, 0)
	database.Db.Model(&walletsModel.WalletPayment{}).Where("admin_id IN ?", adminIds).
		Where("mode = ?", walletsModel.WalletPaymentModeWithdraw).Find(&paymentList)

	data := make([]*views.InputOptions, 0)
	for _, payment := range paymentList {
		data = append(data, &views.InputOptions{
			Label: payment.Name + "." + strconv.Itoa(int(payment.AdminId)),
			Value: payment.ID,
		})
	}
	return data
}
