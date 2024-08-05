package walletsService

import (
	"errors"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/model/walletsModel"
	"gorm.io/gorm"
)

type UserWallet struct {
	tx         *gorm.DB
	userInfo   *usersModel.User
	assetsInfo *walletsModel.WalletAssets
}

func NewUserWallet(tx *gorm.DB, userInfo *usersModel.User, assetsInfo *walletsModel.WalletAssets) *UserWallet {
	return &UserWallet{tx: tx, userInfo: userInfo, assetsInfo: assetsInfo}
}

// SetAssetsInfo 设置资产信息
func (_UserWallet *UserWallet) SetAssetsInfo(assetsInfo *walletsModel.WalletAssets) *UserWallet {
	_UserWallet.assetsInfo = assetsInfo
	return _UserWallet
}

// SetUserInfo 设置用户信息
func (_UserWallet *UserWallet) SetUserInfo(userInfo *usersModel.User) *UserWallet {
	_UserWallet.userInfo = userInfo
	return _UserWallet
}

// ChangeUserBalance 变更用户余额
func (_UserWallet *UserWallet) ChangeUserBalance(billType int, sourceId uint, money float64) error {
	userBillService := NewUserBill()

	// 如果余额小于0, 或者账单类型不存在
	billTypeLabel := userBillService.GetUserBillName(billType)
	if money <= 0 || billTypeLabel == "" || _UserWallet.userInfo == nil {
		return errors.New("abnormalOperation")
	}

	// 账单类型负数, 证明这个是扣款操作 - 判断余额是否不足
	resultMoney := _UserWallet.userInfo.Money + money
	if billType <= 0 {
		resultMoney = _UserWallet.userInfo.Money - money
	}
	if resultMoney < 0 {
		return errors.New("insufficientBalance")
	}

	// 更新用户资产金额
	result := _UserWallet.tx.Model(&usersModel.User{}).Where("id = ?", _UserWallet.userInfo.ID).Update("money", resultMoney)
	if result.Error != nil {
		return result.Error
	}

	// 插入账单
	result = _UserWallet.tx.Create(&walletsModel.WalletUserBill{
		AdminId:  _UserWallet.userInfo.AdminId,
		UserId:   _UserWallet.userInfo.ID,
		SourceId: sourceId,
		Type:     billType,
		Name:     billTypeLabel,
		Money:    money,
		Balance:  _UserWallet.userInfo.Money,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// ChangeUserAssets 变更用户资产
func (_UserWallet *UserWallet) ChangeUserAssets(billType int, sourceId uint, money float64) error {
	userBillService := NewUserBill()

	// 如果余额小于0, 或者账单类型不存在
	billTypeLabel := userBillService.GetUserBillName(billType)
	if money <= 0 || billTypeLabel == "" || _UserWallet.assetsInfo == nil || _UserWallet.userInfo == nil {
		return errors.New("abnormalOperation")
	}

	// 查询用户资产信息
	userAssetsInfo := &walletsModel.WalletUserAssets{}
	result := _UserWallet.tx.Model(userAssetsInfo).Where("user_id = ?", _UserWallet.userInfo.ID).
		Where("assets_id = ?", _UserWallet.assetsInfo.ID).Find(userAssetsInfo)
	if result.Error != nil {
		return result.Error
	}

	// 如果当前资产用户没有, 那么给当前用户创建资产
	if userAssetsInfo.ID == 0 {
		userAssetsInfo = &walletsModel.WalletUserAssets{
			AdminId:  _UserWallet.userInfo.AdminId,
			UserId:   _UserWallet.userInfo.ID,
			AssetsId: _UserWallet.assetsInfo.ID,
		}
		result = _UserWallet.tx.Create(userAssetsInfo)
		if result.Error != nil {
			return result.Error
		}
	}

	// 账单类型负数, 证明这个是扣款操作 - 判断余额是否不足
	resultMoney := userAssetsInfo.Money + money
	if billType <= 0 {
		resultMoney = userAssetsInfo.Money - money
	}
	if resultMoney < 0 {
		return errors.New("insufficientAssets")
	}

	// 更新用户资产金额
	result = _UserWallet.tx.Model(&walletsModel.WalletUserAssets{}).Where("id = ?", userAssetsInfo.ID).Update("money", resultMoney)
	if result.Error != nil {
		return result.Error
	}

	// 插入账单
	result = _UserWallet.tx.Create(&walletsModel.WalletUserBill{
		AdminId:  _UserWallet.userInfo.AdminId,
		UserId:   _UserWallet.userInfo.ID,
		AssetsId: _UserWallet.assetsInfo.ID,
		SourceId: sourceId,
		Type:     billType,
		Name:     billTypeLabel,
		Money:    money,
		Balance:  userAssetsInfo.Money,
	})
	return result.Error
}
