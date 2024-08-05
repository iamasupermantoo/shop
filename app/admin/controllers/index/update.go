package index

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// UpdateParams 更新管理
type UpdateParams struct {
	Avatar   string `json:"avatar"`   //	头像
	Email    string `json:"email"`    //	邮箱
	NickName string `json:"nickName"` //	昵称
	Domains  string `json:"domains"`  //	域名
	Online   string `json:"online"`   //	客服链接
	SeatLink string `json:"seatLink"` // 坐席链接
}

// Update 更新管理
// func Update(ctx *fiber.Ctx) error {
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	adminInfo := &adminsModel.AdminUser{}
	result := database.Db.Model(adminInfo).Where("id = ?", ctx.AdminId).Find(adminInfo)
	if result.Error != nil || adminInfo.ID == 0 {
		return ctx.ErrorJson("找不到管理信息")
	}

	// 更新当前域名
	if params.Domains != "" {
		err := adminsService.NewAdminUser(ctx.Rds, adminInfo.ID).UpdateDomains(adminInfo.Domains, params.Domains)
		if err != nil {
			return ctx.ErrorJson(err.Error())
		}
	}

	if err := database.Db.Model(&adminsModel.AdminUser{}).Where("id = ?", adminInfo.ID).Updates(params).Error; err != nil {
		return ctx.ErrorJson(err.Error())
	}

	return ctx.SuccessJsonOK()
}
