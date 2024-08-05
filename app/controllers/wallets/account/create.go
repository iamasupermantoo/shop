package account

import (
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"strconv"
)

type CreateParams struct {
	PaymentId uint   `json:"paymentId" validate:"required"` //	提现类型
	Name      string `json:"name"`                          // 银行名称
	RealName  string `json:"realName"`                      // 真实姓名
	Number    string `json:"number"`                        // 证件卡号
	Code      string `json:"code"`                          // 银行代码
	Remark    string `json:"remark"`                        // 备注信息
}

// Create 提现账户新增
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	// 查询支付类型
	paymentInfo := &walletsModel.WalletPayment{}
	result := database.Db.Model(paymentInfo).Where("id = ?", params.PaymentId).
		Where("status = ?", walletsModel.WalletPaymentStatusActive).
		Where("mode IN ?", []uint{walletsModel.WalletPaymentModeWithdraw, walletsModel.WalletPaymentModeAssetsWithdraw}).
		Where("admin_id = ?", ctx.AdminSettingId).Find(paymentInfo)
	if result.Error != nil || paymentInfo.ID == 0 {
		return ctx.ErrorJson("abnormalOperation")
	}

	// 验证是否超过绑定数量
	var accountNums int64
	database.Db.Model(&walletsModel.WalletUserAccount{}).Where("user_id = ?", ctx.UserId).Count(&accountNums)
	accountLimitNums := adminsService.NewAdminSetting(ctx.Rds, ctx.AdminSettingId).GetRedisAdminSettingField("walletAccountNums")
	nums, _ := strconv.ParseInt(accountLimitNums, 10, 64)
	if accountNums >= nums {
		return ctx.ErrorJsonTranslate("limitExceeded", strconv.FormatInt(nums, 10))
	}
	if params.Name == "" {
		params.Name = paymentInfo.Name
	}
	result = database.Db.Create(&walletsModel.WalletUserAccount{
		PaymentId: params.PaymentId,
		AdminId:   ctx.AdminId,
		UserId:    ctx.UserId,
		Name:      params.Name,
		RealName:  params.RealName,
		Number:    params.Number,
		Code:      params.Code,
		Remark:    params.Remark,
	})
	if result.Error != nil {
		return ctx.ErrorJsonTranslate("abnormalOperation", result.Error.Error())
	}

	return ctx.SuccessJsonOK()
}
