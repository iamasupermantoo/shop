package usersService

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/database"
	"gofiber/app/module/views"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

// GetUserVirtualOption 获取虚拟用户选项
func (_User *User) GetUserVirtualOption(adminSettingId uint) (options []*views.InputOptions) {
	userList := make([]*usersModel.User, 0)
	database.Db.Where("admin_id = ?", adminSettingId).
		Where("type = ?", usersModel.UserTypeVirtual).
		Find(&userList)
	for _, user := range userList {
		options = append(options, &views.InputOptions{
			Label: user.UserName,
			Value: user.ID,
		})
	}
	return options
}
