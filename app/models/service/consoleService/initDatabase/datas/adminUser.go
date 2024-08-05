package datas

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/types"
	"gofiber/utils"
	"time"
)

// InitAdminUser 初始化管理
func InitAdminUser() []*adminsModel.AdminUser {
	password := utils.PasswordEncrypt("taozijun")
	otherPassword := utils.PasswordEncrypt("abc123")

	nowTime := time.Now().Add(365 * 24 * time.Hour)
	return []*adminsModel.AdminUser{
		{GormModel: types.GormModel{ID: adminsModel.SuperAdminId}, ExpiredAt: nowTime, ParentId: 0, UserName: "admin", Avatar: "/assets/icon/online.png", SeatLink: "", NickName: "八戒网络科技", Email: "muiprosperyls15@gmail.com", Password: password, SecurityKey: password},
		{GormModel: types.GormModel{ID: adminsModel.MerchantAdminId}, ExpiredAt: nowTime, ParentId: adminsModel.SuperAdminId, Avatar: "/assets/icon/online.png", SeatLink: "https://online.ainn.us/seat?u=ceshi&p=abc123", UserName: "merchant", NickName: "八戒网络科技", Email: "muiprosperyls15@gmail.com", Domains: "localhost:9100,192.168.5.135:9100", Password: otherPassword, SecurityKey: otherPassword, Data: adminsModel.NewMerchantData()},
		{GormModel: types.GormModel{ID: adminsModel.AgentAdminId}, ExpiredAt: nowTime, ParentId: adminsModel.MerchantAdminId, Avatar: "/assets/icon/online.png", UserName: "agent", SeatLink: "", Domains: "", NickName: "八戒网络科技", Email: "muiprosperyls10@gmail.com", Password: otherPassword, SecurityKey: otherPassword},
	}
}
