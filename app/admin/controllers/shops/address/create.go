package address

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// CreateParams 新增参数
type CreateParams struct {
	UserName string `json:"userName"` // 用户名
	Name     string `json:"name"`     // 收件人名称
	Contact  string `json:"contact"`  // 联系方式
	City     string `json:"city"`     // 国家城市
	Address  string `json:"address"`  // 详细地址
	Type     int    `json:"type"`     // 类型 1收货地址 2发货地址
}

// Create 新增接口
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	userInfo := &usersModel.User{}
	result := database.Db.Where("user_name = ?", params.UserName).Where("admin_id In ?", ctx.GetAdminChildIds()).Find(userInfo)
	if result.Error != nil || userInfo.ID == 0 {
		return ctx.ErrorJson("找不到用户信息")
	}

	createInfo := &shopsModel.ShippingAddress{
		AdminId: userInfo.AdminId,
		UserId:  userInfo.ID,
		Name:    params.Name,
		Contact: params.Contact,
		City:    params.City,
		Address: params.Address,
		Type:    params.Type,
	}

	result = database.Db.Create(createInfo)
	if result.Error != nil {
		return ctx.ErrorJson("添加失败, 原因 => " + result.Error.Error())
	}

	return ctx.SuccessJsonOK()
}
