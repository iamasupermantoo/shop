package auth

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

type CreateParams struct {
	Type     int    `json:"type" validate:"required"`     // 类型
	RealName string `json:"realName" validate:"required"` // 证件姓名
	Photo1   string `json:"photo1" validate:"required"`   // 证件照1
	Number   string `json:"number"`                       // 证件号码
	Photo2   string `json:"photo2"`                       // 证件照2
	Address  string `json:"address"`                      // 详细地址
	Photo3   string `json:"photo3"`                       // 证件照3
	Status   int    `json:"-"`                            // 状态
}

// Create  申请认证
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	//	用户信息
	userInfo := &usersModel.User{}
	result := database.Db.Where("id = ?", ctx.UserId).Find(userInfo)
	if result.Error != nil || userInfo.ID == 0 {
		return ctx.ErrorJson("abnormalOperation")
	}

	currentAuthInfo := &usersModel.UserAuth{}
	result = database.Db.Where("user_id = ?", userInfo.ID).Find(currentAuthInfo)
	if currentAuthInfo.Status == usersModel.UserAuthStatusActive || currentAuthInfo.Status == usersModel.UserAuthStatusComplete {
		return ctx.SuccessJsonOK()
	}

	// 没有记录, 新增记录
	if result.RowsAffected == 0 {
		database.Db.Create(&usersModel.UserAuth{
			AdminId:  userInfo.AdminId,
			UserId:   userInfo.ID,
			RealName: params.RealName,
			Number:   params.Number,
			Photo1:   params.Photo1,
			Photo2:   params.Photo2,
			Photo3:   params.Photo3,
			Address:  params.Address,
		})
	} else {
		params.Status = usersModel.UserAuthStatusActive
		database.Db.Model(&usersModel.UserAuth{}).Where(currentAuthInfo.ID).Updates(params)
	}
	return ctx.SuccessJsonOK()
}
