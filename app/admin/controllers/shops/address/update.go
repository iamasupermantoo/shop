package address

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gorm.io/gorm"
)

// UpdateParams 更新参数
type UpdateParams struct {
	ID      uint   `json:"id" gorm:"-" validate:"required"` // ID
	Name    string `json:"name"`                            // 收件人名称
	Contact string `json:"contact"`                         // 联系方式
	City    string `json:"city"`                            // 国家城市
	Address string `json:"address"`                         // 详细地址
	Status  int    `json:"status"`                          // 状态 -2删除 -1禁用 10激活
	IsShow  int    `json:"isShow"`                          // 1不默认 2默认
}

// Update 更新接口
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	//	通过id参数查询地址
	shoppingAddressInfo := &shopsModel.ShippingAddress{}
	result := database.Db.Where(params.ID).Where("admin_id IN ?", ctx.GetAdminChildIds()).Find(&shoppingAddressInfo)
	if result.Error != nil || shoppingAddressInfo.ID == 0 {
		return ctx.ErrorJson("找不到对应的地址信息")
	}

	err := database.Db.Transaction(func(tx *gorm.DB) error {
		//	修改该用户要修改的数据
		result = tx.Where("id = ?", shoppingAddressInfo.ID).Updates(&shopsModel.ShippingAddress{
			Name:    params.Name,
			Contact: params.Contact,
			City:    params.City,
			Address: params.Address,
			Status:  params.Status,
			IsShow:  params.IsShow,
		})
		if result.Error != nil {
			return ctx.ErrorJson("更新地址信息失败")
		}

		//	如果选中的是默认状态, 就修改该用户其他地址为非默认
		if params.IsShow == shopsModel.ShippingAddressIsShowYes {
			//	修改该用户的其他地址为非默认地址
			result = tx.Where("id <> ?", params.ID).Where("type = ?", shoppingAddressInfo.Type).
				Where("user_id = ?", shoppingAddressInfo.UserId).
				Updates(&shopsModel.ShippingAddress{
					IsShow: shopsModel.ShippingAddressIsShowNo,
				})
			if result.Error != nil {
				return ctx.ErrorJson("更新默认地址失败")
			}
		}

		return nil
	})
	if err != nil {
		return ctx.ErrorJson(err.Error())
	}

	return ctx.SuccessJsonOK()
}
