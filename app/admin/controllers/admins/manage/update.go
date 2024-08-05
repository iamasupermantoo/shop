package manage

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/types"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// UpdateParams 更新管理参数
type UpdateParams struct {
	ID          uint                     `gorm:"-" validate:"required" json:"id"` //	ID
	Avatar      string                   `json:"avatar"`                          //	头像
	NickName    string                   `json:"nickName"`                        //	昵称
	Password    types.GormPasswordParams `json:"password"`                        //	密码
	SecurityKey types.GormPasswordParams `json:"securityKey"`                     //	安全密码
	Email       string                   `json:"email"`                           //	邮箱
	Domains     string                   `json:"domains"`                         //	域名
	Online      string                   `json:"online"`                          //	客服链接
	SeatLink    string                   `json:"seatLink"`                        // 坐席链接
	Status      int                      `json:"status"`                          //	状态
	Money       float64                  `json:"money"`                           //	金额
	ExpiredAt   types.GormTimeParams     `json:"expiredAt"`                       //	过期时间
	Role        string                   `gorm:"-" json:"role"`                   //	角色
}

// Update 更新管理
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	currentAdminInfo := &adminsModel.AdminUser{}
	result := database.Db.Model(currentAdminInfo).Where("id IN ?", ctx.GetAdminChildIds()).Where("id = ?", params.ID).Find(currentAdminInfo)
	if result.Error != nil || currentAdminInfo.ID == 0 {
		return ctx.ErrorJson("没有找到对应的管理信息")
	}

	// 如果是绑定过的域名, 那么不能使用
	adminService := adminsService.NewAdminUser(ctx.Rds, currentAdminInfo.ID)
	if params.Domains != "" {
		err := adminService.UpdateDomains(currentAdminInfo.Domains, params.Domains)
		if err != nil {
			return ctx.ErrorJson(err.Error())
		}
	}

	// 如果修改角色
	if params.Role != "" {
		adminInfo := &adminsModel.AdminUser{}
		result = database.Db.Model(adminInfo).Where("id = ?", params.ID).Find(adminInfo)
		if result.Error != nil || adminInfo.ID == 0 {
			return ctx.ErrorJson("管理员不存在~")
		}

		database.Db.Model(&adminsModel.AuthAssignment{}).Where("admin_id = ?", adminInfo.ID).Update("name", params.Role)
		adminsService.NewAdminMenu(ctx.Rds, ctx.AdminId).DelRedisAdminMenuList()
	}

	database.Db.Model(&adminsModel.AdminUser{}).
		Where("id = ?", currentAdminInfo.ID).
		Updates(params)

	// 删除管理过期时间
	adminService.DelRedisExpiration()
	return ctx.SuccessJsonOK()
}
