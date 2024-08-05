package order

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/utils"
	"time"
)

// CreateParams 新增参数
type CreateParams struct {
	UserName  string `json:"userName"`  // 用户账户
	ProductId int    `json:"productId"` // 产品ID
	Type      int    `json:"type"`      // 类型
}

// Create 新增接口
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	adminIds := ctx.GetAdminChildIds()
	userInfo := &usersModel.User{}
	result := database.Db.Model(userInfo).Where("user_name = ?", params.UserName).Where("admin_id IN ?", adminIds).Find(userInfo)
	if result.Error != nil || userInfo.ID == 0 {
		return ctx.ErrorJson("查询不到用户")
	}

	productInfo := &productsModel.Product{}
	result = database.Db.Model(&productsModel.Product{}).Where("id = ?", params.ProductId).Where("admin_id IN ?", adminIds).Find(productInfo)
	if result.Error != nil || productInfo.ID == 0 {
		return ctx.ErrorJson("查询不到产品")
	}

	createInfo := &productsModel.ProductOrder{
		AdminId:   userInfo.AdminId,
		UserId:    userInfo.ID,
		ProductId: productInfo.ID,
		OrderSn:   utils.NewRandom().OrderSn(),
		Money:     productInfo.Money,
		Type:      params.Type,
		ExpiredAt: time.Now().Add(30 * 24 * time.Hour),
	}

	result = database.Db.Create(createInfo)
	if result.Error != nil {
		return ctx.ErrorJson("添加失败, 原因 => " + result.Error.Error())
	}

	return ctx.SuccessJsonOK()
}
