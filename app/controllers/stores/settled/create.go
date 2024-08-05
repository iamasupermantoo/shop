package settled

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gorm.io/gorm"
)

// CreateParams 新增参数
type CreateParams struct {
	Logo      string `json:"logo" validate:"required"` // 店铺logo
	Name      string `json:"name" validate:"required"` // 店铺名字
	Address   string `json:"address"`                  // 店铺详细地址
	CountryId uint   `json:"countryId"`                // 国家ID
	Number    string `json:"number"`                   // 证件号
	RealName  string `json:"realName"`                 // 用户真实名称
	Photo1    string `json:"photo1"`                   // 证件正
	Photo2    string `json:"photo2"`                   // 证件反
	Photo3    string `json:"photo3"`                   // 手持证件照
	Email     string `json:"email"`                    // email
	Contact   string `json:"contact"`                  // 电话
}

// Create 商家入驻
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	settledTemplate := adminsService.NewAdminSetting(ctx.Rds, ctx.AdminSettingId).CheckBoxToMaps("settledTemplate")
	if settledTemplate["showCountry"] && params.Contact == "" {
		return ctx.ErrorJson("")
	}
	if settledTemplate["showNumber"] && params.Number == "" {
		return ctx.ErrorJson("number is not null")
	}
	if settledTemplate["showRealName"] && params.RealName == "" {
		return ctx.ErrorJson("realName is not null")
	}
	if settledTemplate["showPhoto2"] && params.Photo2 == "" {
		return ctx.ErrorJson("photo2 is not null")
	}
	if settledTemplate["showPhoto3"] && params.Photo3 == "" {
		return ctx.ErrorJson("photo3 is not null")
	}
	if settledTemplate["showEmail"] && params.Email == "" {
		return ctx.ErrorJson("email is not null")
	}
	if settledTemplate["showContact"] && params.Contact == "" {
		return ctx.ErrorJson("contact is not null")
	}

	// 查询当前商家入驻信息
	settledInfo := &shopsModel.StoreSettled{}
	result := database.Db.Model(settledInfo).Where("user_id = ?", ctx.UserId).Find(settledInfo)
	if result.Error != nil {
		return ctx.ErrorJsonTranslate(result.Error.Error())
	}

	userInfo := &usersModel.User{}
	result = database.Db.Model(userInfo).Where("id = ?", ctx.UserId).Find(userInfo)
	if result.Error != nil || userInfo.ID == 0 {
		return ctx.ErrorJsonTranslate("abnormalOperation")
	}

	if settledInfo.ID == 0 {
		err := database.Db.Transaction(func(tx *gorm.DB) error {
			tx.Create(&shopsModel.StoreSettled{
				AdminId:   ctx.AdminId,
				UserId:    ctx.UserId,
				CountryId: params.CountryId,
				Name:      params.Name,
				Address:   params.Address,
				RealName:  params.RealName,
				Logo:      params.Logo,
				Photo1:    params.Photo1,
				Photo2:    params.Photo2,
				Photo3:    params.Photo3,
				Number:    params.Number,
				Email:     params.Email,
				Contact:   params.Contact,
			})
			return nil
		})
		if err != nil {
			return err
		}

	} else {
		if settledInfo.Status == shopsModel.StoreSettledStatusRefuse {
			database.Db.Where("id = ?", settledInfo.ID).Updates(&shopsModel.StoreSettled{
				CountryId: params.CountryId,
				Name:      params.Name,
				Address:   params.Address,
				RealName:  params.RealName,
				Logo:      params.Logo,
				Photo1:    params.Photo1,
				Photo2:    params.Photo2,
				Photo3:    params.Photo3,
				Number:    params.Number,
				Email:     params.Email,
				Contact:   params.Contact,
				Status:    shopsModel.StoreSettledStatusPending,
			})
		}
	}

	return ctx.SuccessJsonOK()
}
