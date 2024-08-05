package address

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gorm.io/gorm"
)

type UpdateParams struct {
	ID      uint   `json:"id" validate:"required"` //	收获地址
	Name    string `json:"name"`
	Contact string `json:"contact"`
	City    string `json:"city"`
	Address string `json:"address"`
	IsShow  int    `json:"isShow"`
}

// Update 编辑收获地址
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	addressInfo := &shopsModel.ShippingAddress{}
	database.Db.Where("id = ?", params.ID).Where("user_id = ?", ctx.UserId).Where("status = ?", shopsModel.ShippingAddressStatusActivate).Find(addressInfo)
	if addressInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("findError")
	}

	err := database.Db.Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&shopsModel.ShippingAddress{}).Where("id = ?", params.ID).Updates(&shopsModel.ShippingAddress{
			Name:    params.Name,
			Contact: params.Contact,
			City:    params.City,
			Address: params.Address,
			IsShow:  params.IsShow,
		})
		if result.Error != nil {
			return result.Error
		}

		//	如果选中的是默认状态, 就修改该用户其他地址为非默认
		if params.IsShow == shopsModel.ShippingAddressIsShowYes {
			result = tx.Model(&shopsModel.ShippingAddress{}).Where("id <> ?", addressInfo.ID).
				Where("type = ?", shopsModel.ShippingAddressTypeReceiving).
				Where("user_id = ?", addressInfo.UserId).
				Update("is_show", shopsModel.ShippingAddressIsShowNo)
			if result.Error != nil {
				return result.Error
			}
		}

		return nil
	})
	if err != nil {
		return ctx.ErrorJsonTranslate("abnormalOperation")
	}

	return ctx.SuccessJsonOK()
}
