package adminsService

import (
	"github.com/goccy/go-json"
	"github.com/gomodule/redigo/redis"
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/module/database"
	"gofiber/app/module/views"
	"gofiber/utils"
)

const (
	// RedisAuthAdminRouter 管理路由列表
	RedisAuthAdminRouter = "RedisAuthAdminRouter"
)

type AdminAuth struct {
	rdsConn redis.Conn
	adminId uint
}

func NewAdminAuth(rdsConn redis.Conn, adminId uint) *AdminAuth {
	return &AdminAuth{rdsConn: rdsConn, adminId: adminId}
}

// GetRedisAdminRouterList 获取管理路由列表
func (_AdminAuth *AdminAuth) GetRedisAdminRouterList() []string {
	routerListBytes, err := redis.Bytes(_AdminAuth.rdsConn.Do("HGET", RedisAuthAdminRouter, _AdminAuth.adminId))
	routerList := make([]string, 0)
	if err != nil {
		// 获取所有角色
		rolesList := _AdminAuth.GetAdminRoles()

		// 获取所有角色对应的路由列表
		if len(rolesList) > 0 {
			routerNameList := make([]string, 0)
			database.Db.Model(&adminsModel.AuthChild{}).Select("child").
				Where("parent IN ?", rolesList).Where("type = ?", adminsModel.AuthChildTypeRoleRouteName).
				Find(&routerNameList)

			// 路由名称对应的路由列表
			if len(routerNameList) > 0 {
				database.Db.Model(&adminsModel.AuthChild{}).Select("child").
					Where("parent IN ?", routerNameList).Where("type = ?", adminsModel.AuthChildTypeRouteNameRoute).
					Find(&routerList)

				// 设置当前路由列表
				if len(routerList) > 0 {
					routerListBytes, _ = json.Marshal(routerList)
					_, _ = _AdminAuth.rdsConn.Do("HSET", RedisAuthAdminRouter, _AdminAuth.adminId, routerListBytes)
				}

			}
			return routerList
		}
	}
	_ = json.Unmarshal(routerListBytes, &routerList)
	return routerList
}

// GetAdminRoles 获取管理角色列表
func (_AdminAuth *AdminAuth) GetAdminRoles() []string {
	rolesList := make([]string, 0)
	database.Db.Model(&adminsModel.AuthAssignment{}).Select("name").Where("admin_id = ?", _AdminAuth.adminId).Find(&rolesList)
	return rolesList
}

// GetAdminRolesRouterSelectOptions 获取全部权限, 角色是否选中
func (_AdminAuth *AdminAuth) GetAdminRolesRouterSelectOptions(rolesList []string) []*views.InputOptions {
	routerNameList := make([]string, 0)
	database.Db.Model(&adminsModel.AuthChild{}).Select("child").
		Where("parent IN ?", rolesList).Where("type = ?", adminsModel.AuthChildTypeRoleRouteName).
		Find(&routerNameList)

	// 获取所有权限列表
	authChildList := make([]*adminsModel.AuthChild, 0)
	database.Db.Model(&adminsModel.AuthChild{}).Where("type = ?", adminsModel.AuthChildTypeRouteNameRoute).Find(&authChildList)

	data := make([]*views.InputOptions, 0)
	for _, child := range authChildList {
		selectEd := false
		if utils.ArrayStringIndexOf(routerNameList, child.Parent) > -1 {
			selectEd = true
		}

		data = append(data, &views.InputOptions{
			Label: child.Parent,
			Value: selectEd,
		})
	}

	return data
}

// GetRolesChildrenOptions 获取管理角色下级列表
func (_AdminAuth *AdminAuth) GetRolesChildrenOptions() []*views.InputOptions {
	rolesList := _AdminAuth.GetAdminRoles()
	childrenRoles := make([]*views.InputOptions, 0)

	childList := make([]*adminsModel.AuthChild, 0)
	database.Db.Model(&adminsModel.AuthChild{}).Where("type = ?", adminsModel.AuthChildTypeRoleParentRole).Where("parent IN ?", rolesList).Find(&childList)

	for _, child := range childList {
		childrenRoles = append(childrenRoles, &views.InputOptions{
			Label: child.Child,
			Value: child.Child,
		})
	}

	return childrenRoles
}

// GetRolesOptions 获取角色列表
func (_AdminAuth *AdminAuth) GetRolesOptions() []*views.InputOptions {
	rolesList := make([]*views.InputOptions, 0)

	childList := make([]*adminsModel.AuthChild, 0)
	database.Db.Model(&adminsModel.AuthChild{}).Where("type = ?", adminsModel.AuthChildTypeRoleParentRole).Find(&childList)

	for _, child := range childList {
		rolesList = append(rolesList, &views.InputOptions{
			Label: child.Child,
			Value: child.Child,
		})
	}

	return rolesList
}
