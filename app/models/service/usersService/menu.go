package usersService

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/module/database"
	"gofiber/app/module/views"
	"strconv"
)

type UserMenu struct {
}

func NewUserMenu() *UserMenu {
	return &UserMenu{}
}

// GetAdminOptions 获取管理options
func (_UserMenu *UserMenu) GetAdminOptions(adminIds []uint) []*views.InputOptions {
	menuList := make([]*systemsModel.Menu, 0)
	database.Db.Model(&systemsModel.Menu{}).Where("admin_id in ?", adminIds).Where("status = ?", systemsModel.MenuStatusActive).Find(&menuList)

	data := make([]*views.InputOptions, 0)
	for _, menu := range menuList {
		data = append(data, &views.InputOptions{
			Label: menu.Name + "." + strconv.Itoa(int(menu.AdminId)),
			Value: menu.ID,
		})
	}
	return data
}
