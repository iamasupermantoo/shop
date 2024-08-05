package datas

import (
	"github.com/gofiber/fiber/v2"
	"gofiber/app/admin"
	"gofiber/app/models/model/adminsModel"
	"gofiber/utils"
	"strings"
)

const (
	roleLoadRouteTypeDefault = 0 //	拥有全部路由
	roleLoadRouteTypeFilter  = 1 //	过滤路由
	roleLoadRouteTypeOwner   = 2 //	拥有的路由

)

type roleLoadRoute struct {
	Role   string   // 角色
	Type   int      // 0什么都不做 1过滤路由 2拥有路由
	Routes []string // 路由
}

// InitAdminAuth 初始化管理权限
func InitAdminAuth() ([]*adminsModel.AuthItem, []*adminsModel.AuthChild) {
	routeList := admin.InitAdminApp().GetRoutes()

	routeFilterList := make([]fiber.Route, 0)
	for _, route := range routeList {
		if route.Method != "HEAD" {
			routeFilterList = append(routeFilterList, route)
		}
	}
	routeList = routeFilterList

	authItem := make([]*adminsModel.AuthItem, 0)
	authChild := make([]*adminsModel.AuthChild, 0)
	authChild = append(authChild, &adminsModel.AuthChild{
		Parent: adminsModel.AuthRoleSuperManage, Child: adminsModel.AuthRoleMerchantManage, Type: adminsModel.AuthChildTypeRoleParentRole,
	})
	authChild = append(authChild, &adminsModel.AuthChild{
		Parent: adminsModel.AuthRoleMerchantManage, Child: adminsModel.AuthRoleAgentManage, Type: adminsModel.AuthChildTypeRoleParentRole,
	})

	roleList := []*roleLoadRoute{
		{Role: adminsModel.AuthRoleSuperManage, Type: roleLoadRouteTypeDefault, Routes: nil},
		{Role: adminsModel.AuthRoleMerchantManage, Type: roleLoadRouteTypeFilter, Routes: []string{
			"管理配置", "管理同步配置", "管理重置配置",
			"管理菜单列表", "管理菜单更新", "管理菜单新增", "管理菜单删除", "管理菜单配置", "管理菜单视图",
			"管理角色列表", "管理角色更新", "管理角色删除", "管理角色新增", "管理角色配置", "管理角色视图",
			"翻译语言翻译", "管理权限添加",
		}},
		{Role: adminsModel.AuthRoleAgentManage, Type: roleLoadRouteTypeOwner, Routes: []string{
			"上传文件", "首页信息", "管理员更新", "管理员更新密码",
			"用户提现账户列表", "用户提现账户更新", "用户提现账户视图",
			"用户资产列表", "用户资产更新", "用户资产视图",
			"钱包账单列表", "钱包账单视图",
			"余额充值订单", "余额提现订单",
			"用户列表", "用户删除", "用户更新", "用户新增", "用户视图",
			"用户访问列表", "用户访问视图",
			"用户邀请列表", "用户邀请删除", "用户邀请新增", "用户邀请更新", "用户邀请视图",
			"用户认证列表", "用户认证同意", "用户认证拒绝", "用户认证视图",
		}},
	}

	for _, route := range routeList {
		if strings.Index(route.Path, "/"+adminsModel.ServiceAdminAuthRouteName+"/") > -1 {
			authItem = append(authItem, &adminsModel.AuthItem{Name: route.Name, Type: adminsModel.AuthItemTypeName})
			authItem = append(authItem, &adminsModel.AuthItem{Name: route.Path, Type: adminsModel.AuthItemTypeRoute})

			authChild = append(authChild, &adminsModel.AuthChild{Parent: route.Name, Child: route.Path, Type: adminsModel.AuthChildTypeRouteNameRoute})

			for _, role := range roleList {
				switch role.Type {
				case roleLoadRouteTypeDefault:
					authChild = append(authChild, &adminsModel.AuthChild{Parent: role.Role, Child: route.Name, Type: adminsModel.AuthChildTypeRoleRouteName})
				case roleLoadRouteTypeFilter:
					if utils.ArrayStringIndexOf(role.Routes, route.Name) == -1 {
						authChild = append(authChild, &adminsModel.AuthChild{Parent: role.Role, Child: route.Name, Type: adminsModel.AuthChildTypeRoleRouteName})
					}
				case roleLoadRouteTypeOwner:
					if utils.ArrayStringIndexOf(role.Routes, route.Name) > -1 {
						authChild = append(authChild, &adminsModel.AuthChild{Parent: role.Role, Child: route.Name, Type: adminsModel.AuthChildTypeRoleRouteName})
					}
				}
			}
		}
	}

	return authItem, authChild
}
