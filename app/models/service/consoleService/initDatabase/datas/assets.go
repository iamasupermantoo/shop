package datas

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/types"
	"gofiber/app/models/model/walletsModel"
)

func InitWalletAssets() []*walletsModel.WalletAssets {
	return []*walletsModel.WalletAssets{
		{GormModel: types.GormModel{ID: 1}, AdminId: adminsModel.SuperAdminId, Name: "USDT", Symbol: "USDT", Icon: "/assets/icon/usdt.png", Type: walletsModel.WalletAssetsTypeDigitalCurrency, Rate: 1},
		{GormModel: types.GormModel{ID: 2}, AdminId: adminsModel.SuperAdminId, Name: "USDC", Symbol: "USDC", Icon: "/assets/icon/usdc.png", Type: walletsModel.WalletAssetsTypeDigitalCurrency, Rate: 1},
		{GormModel: types.GormModel{ID: 3}, AdminId: adminsModel.SuperAdminId, Name: "BTC", Symbol: "BTC", Icon: "/assets/icon/btc.png", Type: walletsModel.WalletAssetsTypeDigitalCurrency, Rate: 38546},
		{GormModel: types.GormModel{ID: 4}, AdminId: adminsModel.SuperAdminId, Name: "ETH", Symbol: "ETH", Icon: "/assets/icon/eth.png", Type: walletsModel.WalletAssetsTypeDigitalCurrency, Rate: 2090},
		{GormModel: types.GormModel{ID: 5}, AdminId: adminsModel.SuperAdminId, Name: "TRX", Symbol: "TRX", Icon: "/assets/icon/trx.png", Type: walletsModel.WalletAssetsTypeDigitalCurrency, Rate: 0.1},
	}
}
