package userAccount

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// CreateParams 新增参数
type CreateParams struct {
	UserName  string `json:"userName"`  //	用户名称
	PaymentId uint   `json:"paymentId"` // 支付ID
	Name      string `json:"name"`      // 银行名称｜Token
	RealName  string `json:"realName"`  // 真实姓名
	Number    string `json:"number"`    // 卡号|地址
	Code      string `json:"code"`      // 银行代码
	Remark    string `json:"remark"`    // 备注信息
}

// Create 新增接口
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	adminChildIds := ctx.GetAdminChildIds()
	userInfo := &usersModel.User{}
	result := database.Db.Model(userInfo).Where("user_name = ?", params.UserName).Where("admin_id IN ?", adminChildIds).Find(userInfo)
	if result.Error != nil || userInfo.ID == 0 {
		return ctx.ErrorJson("用户不存在")
	}

	paymentInfo := &walletsModel.WalletPayment{}
	result = database.Db.Model(paymentInfo).Where("id = ?", params.PaymentId).
		Where("admin_id IN ?", adminsService.NewAdminUser(ctx.Rds, userInfo.AdminId).GetRedisChildrenIds()).Find(paymentInfo)
	if result.Error != nil || paymentInfo.ID == 0 {
		return ctx.ErrorJson("提现类型不存在")
	}

	createInfo := &walletsModel.WalletUserAccount{
		AdminId:   userInfo.AdminId,
		UserId:    userInfo.ID,
		PaymentId: paymentInfo.ID,
		Name:      params.Name,
		RealName:  params.RealName,
		Number:    params.Number,
		Code:      params.Code,
		Remark:    params.Remark,
	}

	result = database.Db.Create(createInfo)
	if result.Error != nil {
		return ctx.ErrorJson("添加失败, 原因 => " + result.Error.Error())
	}

	return ctx.SuccessJsonOK()
}
