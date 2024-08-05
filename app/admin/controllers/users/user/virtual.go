package user

import (
	"github.com/brianvoe/gofakeit/v6"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/utils"
	"gorm.io/gorm"
	"strings"
)

// CreateVirtualParams 创建虚拟用户参数
type CreateVirtualParams struct {
	Number   int    `json:"number"`   // 添加数量
	Password string `json:"password"` // 默认密码
}

// CreateVirtual 创建虚拟用户
func CreateVirtual(ctx *context.CustomCtx, params *CreateVirtualParams) error {
	_ = database.Db.Transaction(func(tx *gorm.DB) error {
		for i := 0; i < params.Number; i++ {
			params.Password = utils.PasswordEncrypt(params.Password)
			createInfo := &usersModel.User{
				AdminId:  ctx.AdminId,
				UserName: strings.ReplaceAll(gofakeit.Name(), " ", ""),
				NickName: gofakeit.Name(),
				Money:    100000,
				Password: params.Password,
				Type:     usersModel.UserTypeVirtual,
			}

			err := tx.Create(createInfo).Error
			if err != nil {
				continue
			}

			addressInfo := gofakeit.Address()
			userAddress := shopsModel.ShippingAddress{
				AdminId: ctx.AdminId,
				UserId:  createInfo.ID,
				Name:    gofakeit.Name(),
				Contact: gofakeit.Phone(),
				City:    addressInfo.City,
				Address: addressInfo.Address,
				Type:    shopsModel.ShippingAddressTypeReceiving,
				IsShow:  shopsModel.ShippingAddressIsShowYes,
			}
			err = tx.Create(&userAddress).Error
			if err != nil {
				continue
			}
		}
		return nil
	})

	return ctx.SuccessJsonOK()
}
