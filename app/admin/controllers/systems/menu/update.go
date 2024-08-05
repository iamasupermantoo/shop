package menu

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// UpdateParams 更新参数
type UpdateParams struct {
	ID         uint   `gorm:"-" validate:"required" json:"id"` //	ID
	ParentId   uint   `json:"parentId"`                        // 父级ID
	Name       string `json:"name"`                            // 名称
	Route      string `json:"route"`                           // 路由
	Sort       int    `json:"sort"`                            // 排序
	Icon       string `json:"icon"`                            // 图标
	ActiveIcon string `json:"activeIcon"`                      // 选中图标
	IsDesktop  int    `json:"isDesktop"`                       // 桌面显示
	IsMobile   int    `json:"isMobile"`                        // 手机显示
	Type       int    `json:"type"`                            // 类型1导航菜单 11设置菜单 21更多菜单
	Status     int    `json:"status"`                          // 状态-1禁用 10开启
}

// Update 更新接口
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	menuInfo := &systemsModel.Menu{}
	result := database.Db.Model(menuInfo).Where("id = ?", params.ID).Where("admin_id IN ?", ctx.GetAdminChildIds()).Find(menuInfo)
	if result.RowsAffected == 0 {
		return ctx.ErrorJson("找不到可用的前台菜单")
	}

	// 更新菜单, 并且删除缓存
	database.Db.Model(&systemsModel.Menu{}).Where("id = ?", menuInfo.ID).Updates(params)
	adminsService.NewAdminMenu(ctx.Rds, ctx.AdminSettingId).DelRedisAdminMenuList()

	return ctx.SuccessJsonOK()
}
