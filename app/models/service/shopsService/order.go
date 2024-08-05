package shopsService

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/types"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/models/service/productsService"
	"gofiber/app/models/service/walletsService"
	"gofiber/app/module/database"
	"gofiber/utils"
	"gorm.io/gorm"
)

// SkuParam Sku参数
type SkuParam struct {
	SkuInfo *productsModel.ProductAttrsSku // skuId
	Nums    int                            // 购买数量
}

type StoreOrder struct {
	tx *gorm.DB
}

// NewStoreOrder 新建产品订单模型
func NewStoreOrder(tx *gorm.DB) *StoreOrder {
	return &StoreOrder{tx: tx}
}

// CreatOrder 创建订单
func (_StoreOrder *StoreOrder) CreatOrder(conn redis.Conn, assetsId, storeId uint, userInfo *usersModel.User, addressInfo *shopsModel.ShippingAddress, SkuParams []*SkuParam) error {
	storeOrderInfo := &shopsModel.ProductStoreOrder{
		AdminId:     userInfo.AdminId,
		UserId:      userInfo.ID,
		OrderSn:     utils.NewRandom().OrderSn(),
		AddressData: (*shopsModel.AddressData)(addressInfo),
		StoreId:     storeId,
		AssetsId:    assetsId,
		Status:      shopsModel.ProductStoreOrderStatusShipping,
	}
	result := _StoreOrder.tx.Create(storeOrderInfo)
	if result.Error != nil {
		return errors.New("abnormalOperation")
	}

	// 创建产品订单
	var sumMoney, finalMoney, earningMoney float64
	for _, skuParam := range SkuParams {
		productOrderInfo, err := productsService.NewProductOrder(_StoreOrder.tx).CreateProductOrder(assetsId, storeOrderInfo, skuParam.SkuInfo, skuParam.Nums)
		if err != nil {
			return err
		}
		sumMoney += productOrderInfo.Money
		finalMoney += productOrderInfo.FinalMoney
		earningMoney += productOrderInfo.Earnings
	}

	// 更新店铺订单
	_StoreOrder.tx.Model(&shopsModel.ProductStoreOrder{}).Where("id = ?", storeOrderInfo.ID).Updates(&shopsModel.ProductStoreOrder{
		Money:      sumMoney,
		FinalMoney: finalMoney,
		Earnings:   earningMoney,
	})

	userWallet := walletsService.NewUserWallet(_StoreOrder.tx, userInfo, nil)
	assetsInfo := walletsModel.WalletAssets{}
	if assetsId > 0 {
		settingId := adminsService.NewAdminUser(conn, userInfo.AdminId).GetRedisAdminSettingId(userInfo.AdminId)
		database.Db.Where("id = ?", assetsId).Where("user_id = ?", userInfo.ID).Where("admin_id = ?", settingId).Find(&assetsInfo)
		return userWallet.SetAssetsInfo(&assetsInfo).ChangeUserAssets(walletsModel.WalletUserBillTypeBuyProduct, storeOrderInfo.ID, finalMoney)
	}

	return userWallet.ChangeUserBalance(walletsModel.WalletUserBillTypeBuyProduct, storeOrderInfo.ID, finalMoney)
}

// Shipping 产品发货
func (_StoreOrder *StoreOrder) Shipping(userInfo usersModel.User, storeOrderInfo shopsModel.ProductStoreOrder) error {
	// 更新店铺订单状态
	result := _StoreOrder.tx.Model(&shopsModel.ProductStoreOrder{}).
		Where("id = ?", storeOrderInfo.ID).
		Update("status", shopsModel.ProductStoreOrderStatusProgress)
	if result.Error != nil {
		return errors.New("abnormalOperation")
	}

	// 更新商品订单状态
	result = _StoreOrder.tx.Model(&productsModel.ProductOrder{}).
		Where("store_order_id = ?", storeOrderInfo.ID).
		Update("status", productsModel.ProductOrderStatusProgress)
	if result.Error != nil {
		return errors.New("abnormalOperation")
	}

	// 店铺商家付款
	wholesaleMoney := storeOrderInfo.FinalMoney - storeOrderInfo.Earnings
	return walletsService.NewUserWallet(_StoreOrder.tx, &userInfo, nil).ChangeUserBalance(walletsModel.WalletUserBillTypePurchaseProduct, storeOrderInfo.ID, wholesaleMoney)
}

// Complete 订单完成
func (_StoreOrder *StoreOrder) Complete(storeUserInfo usersModel.User, storeOrderInfo shopsModel.ProductStoreOrder) error {
	// 商家订单完成
	result := _StoreOrder.tx.Model(&shopsModel.ProductStoreOrder{}).
		Where("id = ?", storeOrderInfo.ID).
		Update("status", shopsModel.ProductStoreOrderStatusComplete)
	if result.Error != nil {
		return errors.New("abnormalOperation")
	}

	// 商品订单完成
	result = _StoreOrder.tx.Model(&productsModel.ProductOrder{}).
		Where("store_order_id = ?", storeOrderInfo.ID).
		Update("status", productsModel.ProductOrderStatusComplete)
	if result.Error != nil {
		return errors.New("abnormalOperation")
	}

	productsOrderList := make([]*productsModel.ProductOrder, 0)
	result = _StoreOrder.tx.Model(&productsModel.ProductOrder{}).
		Where("store_order_id = ?", storeOrderInfo.ID).
		Find(&productsOrderList)
	if result.Error != nil {
		return errors.New("abnormalOperation")
	}

	storeCommentList := make([]*shopsModel.StoreComment, 0)
	for _, order := range productsOrderList {
		storeCommentList = append(storeCommentList, &shopsModel.StoreComment{
			AdminId:   order.AdminId,
			UserId:    order.UserId,
			StoreId:   order.StoreId,
			ProductId: order.ProductId,
			OrderId:   order.ID,
			Images:    make(types.GormStrings, 0),
		})
	}
	result = _StoreOrder.tx.Create(&storeCommentList)
	if result.Error != nil {
		return errors.New("abnormalOperation")
	}
	return walletsService.NewUserWallet(_StoreOrder.tx, &storeUserInfo, nil).ChangeUserBalance(walletsModel.WalletUserBillTypeEarnings, storeOrderInfo.ID, storeOrderInfo.Earnings)
}
