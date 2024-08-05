package datas

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/systemsModel"
)

// InitLevel 初始化用户等级
func InitLevel() []*systemsModel.Level {
	return []*systemsModel.Level{
		{AdminId: adminsModel.SuperAdminId, Symbol: 1, Name: "Lv1", Icon: "/assets/diamond.png", Money: 120, Days: -1, Status: 10, Desc: "Description of membership level benefits"},
		{AdminId: adminsModel.SuperAdminId, Symbol: 2, Name: "Lv2", Icon: "/assets/diamond.png", Money: 220, Days: -1, Status: 10, Desc: "Description of membership level benefits"},
		{AdminId: adminsModel.SuperAdminId, Symbol: 3, Name: "Lv3", Icon: "/assets/diamond.png", Money: 320, Days: -1, Status: 10, Desc: "Description of membership level benefits"},
		{AdminId: adminsModel.SuperAdminId, Symbol: 4, Name: "Lv4", Icon: "/assets/diamond.png", Money: 580, Days: -1, Status: 10, Desc: "Description of membership level benefits"},
		{AdminId: adminsModel.SuperAdminId, Symbol: 5, Name: "Lv5", Icon: "/assets/diamond.png", Money: 888, Days: -1, Status: 10, Desc: "Description of membership level benefits"},
	}
}
