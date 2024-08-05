package systemsService

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/module/database"
	"gofiber/app/module/views"
	"strconv"
)

type SystemLevel struct {
}

func NewSystemLevel() *SystemLevel {
	return &SystemLevel{}
}

// GetAdminOptions 获取管理Options
func (_SystemLevel *SystemLevel) GetAdminOptions(adminIds []uint) []*views.InputOptions {
	data := make([]*views.InputOptions, 0)
	levelList := make([]*systemsModel.Level, 0)

	filterSymbolFunc := func(symbol int) bool {
		for _, datum := range data {
			if datum.Value.(int) == symbol {
				return true
			}
		}
		return false
	}

	database.Db.Model(&systemsModel.Level{}).Where("admin_id IN ?", adminIds).Where("status = ?", systemsModel.LevelStatusActive).Find(&levelList)
	for _, levelInfo := range levelList {
		if !filterSymbolFunc(levelInfo.Symbol) {
			data = append(data, &views.InputOptions{
				Label: levelInfo.Name + "." + strconv.Itoa(int(levelInfo.AdminId)),
				Value: levelInfo.Symbol,
			})
		}
	}
	return data
}
