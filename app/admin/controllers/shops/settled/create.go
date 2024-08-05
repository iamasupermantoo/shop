package settled

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// CreateParams 新增参数
type CreateParams struct {
	UserName string `json:"userName" validate:"required" gorm:"-"` // 用户名
	Name     string `json:"name" validate:"required"`              // 店铺名字
	Type     int    `json:"type" validate:"required,oneof=1"`      // 类型 1营业执照
	Number   string `json:"number" validate:"required"`            // 证件号
	Photo1   string `json:"photo1" validate:"required"`            // 证件正
	Photo2   string `json:"photo2"`                                // 证件反
	Contact  string `json:"contact" validate:"required"`           // 联系方式
}

// Create 新增接口
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	//	通过UserName查询用户是否存在
	userInfo := &usersModel.User{}
	if result := database.Db.
		Where("username = ?", params.UserName).
		Where("admin_id in ?", ctx.GetAdminChildIds()).
		Take(userInfo); result.Error != nil {
		return ctx.ErrorJson("该用户不存在")
	}

	// 判断用户是否重复申请入驻
	if result := database.Db.
		Where("user_id = ?", userInfo.ID).
		Take(&shopsModel.StoreSettled{}); result.Error == nil {
		return ctx.ErrorJson("请勿重复申请入驻")
	}

	result := database.Db.Create(&shopsModel.StoreSettled{
		AdminId: ctx.AdminId,
		UserId:  userInfo.ID,
		Type:    params.Type,
		Name:    params.Name,
		Number:  params.Number,
		Contact: params.Contact,
		Photo1:  params.Photo1,
		Photo2:  params.Photo2,
	})

	if result.Error != nil {
		return ctx.ErrorJson("添加失败, 原因 => " + result.Error.Error())
	}

	return ctx.SuccessJsonOK()
}
