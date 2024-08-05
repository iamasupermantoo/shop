package address

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

type CreateParams struct {
	Name    string `json:"name" validate:"required"`    //	收件人
	Contact string `json:"contact" validate:"required"` //	联系方式
	City    string `json:"city"`                        //	国家城市
	Address string `json:"address"`                     //	详细地址
	IsShow  int    `json:"isShow"`                      //	是否默认地址
}

// Create 新增收获地址
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	addressInfo := &shopsModel.ShippingAddress{
		AdminId: ctx.AdminId,
		UserId:  ctx.UserId,
		Name:    params.Name,
		Contact: params.Contact,
		City:    params.City,
		Address: params.Address,
		IsShow:  params.IsShow,
	}
	database.Db.Create(addressInfo)

	//	如果插入数据为默认地址，则将该用户其他地址改为非默认
	if addressInfo.IsShow == shopsModel.ShippingAddressIsShowYes {
		database.Db.Model(&shopsModel.ShippingAddress{}).Where("id <> ?", addressInfo.ID).
			Where("type = ?", shopsModel.ShippingAddressTypeReceiving).
			Where("user_id = ?", addressInfo.UserId).
			Update("is_show", shopsModel.ShippingAddressIsShowNo)
	}

	return ctx.SuccessJsonOK()
}
