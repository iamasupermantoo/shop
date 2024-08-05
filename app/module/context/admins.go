package context

import (
	"gofiber/app/models/service/adminsService"
)

// GetAdminChildIds 获取管理子级Ids
func (_CustomCtx *CustomCtx) GetAdminChildIds() []uint {
	adminService := adminsService.NewAdminUser(_CustomCtx.Rds, _CustomCtx.AdminId)
	return adminService.GetRedisChildrenIds()
}
