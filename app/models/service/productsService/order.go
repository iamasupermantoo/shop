package productsService

import (
	"errors"
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/database"
	"gorm.io/gorm"
	"time"
)

type ProductOrder struct {
	tx *gorm.DB
}

// NewProductOrder 新建产品订单模型
func NewProductOrder(tx *gorm.DB) *ProductOrder {
	return &ProductOrder{tx: tx}
}

// CreateProductOrder 创建产品订单
func (_ProductOrder *ProductOrder) CreateProductOrder(assetsId uint, storeOrderInfo *shopsModel.ProductStoreOrder, skuInfo *productsModel.ProductAttrsSku, nums int) (*productsModel.ProductOrder, error) {
	var increase float64
	if err := database.Db.Raw("select ul.increase from store as s left join user_level as ul on s.user_id = ul.user_id where s.id = ?", storeOrderInfo.StoreId).Scan(&increase).Error; err != nil {
		return nil, errors.New("abnormalOperation")
	}

	// 店铺产品总价
	sumProductOrderMoney := skuInfo.GetTotalPrice(float64(nums))
	// 店铺产品最终价
	productOrderFinalMoney := skuInfo.GetFinalPrice(float64(nums))
	// 店铺用户收益
	productOrderEarningMoney := skuInfo.GetEarning(float64(nums), increase)
	if sumProductOrderMoney == 0 || productOrderFinalMoney == 0 {
		return nil, errors.New("abnormalOperation")
	}

	orderInfo := &productsModel.ProductOrder{
		AdminId:      storeOrderInfo.AdminId,
		UserId:       storeOrderInfo.UserId,
		StoreOrderId: storeOrderInfo.ID,
		OrderSn:      storeOrderInfo.OrderSn,
		StoreId:      storeOrderInfo.StoreId,
		ProductId:    skuInfo.ProductId,
		Money:        sumProductOrderMoney,
		FinalMoney:   productOrderFinalMoney,
		Earnings:     productOrderEarningMoney,
		AssetsId:     assetsId,
		Nums:         nums,
		SkuData:      (*productsModel.SkuData)(skuInfo),
		Status:       productsModel.ProductOrderStatusShipping,
		ExpiredAt:    time.Now().Add(3 * 24 * time.Hour),
	}
	result := _ProductOrder.tx.Create(orderInfo)
	if result.Error != nil {
		return nil, errors.New("abnormalOperation")
	}
	return orderInfo, nil
}
