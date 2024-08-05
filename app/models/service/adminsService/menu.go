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
	// RedisAuthAdminMenu 管理菜单
	RedisAuthAdminMenu = "RedisAuthAdminMenu"
)

type AdminMenu struct {
	rdsConn redis.Conn
	adminId uint
}

func NewAdminMenu(rdsConn redis.Conn, adminId uint) *AdminMenu {
	return &AdminMenu{rdsConn: rdsConn, adminId: adminId}
}

// GetRedisAdminMenuList 获取管理菜单列表
func (_AdminMenu *AdminMenu) GetRedisAdminMenuList() []*adminsModel.AdminMenuInfo {
	menuListBytes, err := redis.Bytes(_AdminMenu.rdsConn.Do("HGET", RedisAuthAdminMenu, _AdminMenu.adminId))
	menuList := make([]*adminsModel.AdminMenuInfo, 0)
	if err != nil {
		routerList := NewAdminAuth(_AdminMenu.rdsConn, _AdminMenu.adminId).GetRedisAdminRouterList()

		adminMenuList := make([]*adminsModel.AdminMenu, 0)
		database.Db.Model(&adminsModel.AdminMenu{}).Where("status = ?", adminsModel.AdminMenuStatusActive).Order("sort asc").Find(&adminMenuList)

		// 递归操作
		if len(adminMenuList) > 0 && len(routerList) > 0 {
			menuList = _AdminMenu.recursiveMenuChildren(0, adminMenuList, routerList)
			if len(menuList) > 0 {
				menuListBytes, _ = json.Marshal(menuList)
				_, _ = _AdminMenu.rdsConn.Do("HSET", RedisAuthAdminMenu, _AdminMenu.adminId, menuListBytes)
			}
		}
		return menuList
	}
	_ = json.Unmarshal(menuListBytes, &menuList)
	return menuList
}

// GetMenuOptions 获取菜单Options
func (_AdminMenu *AdminMenu) GetMenuOptions() []*views.InputOptions {
	menuList := make([]*adminsModel.AdminMenu, 0)
	database.Db.Model(&adminsModel.AdminMenu{}).Where("status = ?", adminsModel.AdminMenuStatusActive).Find(&menuList)

	data := make([]*views.InputOptions, 0)
	for _, menu := range menuList {
		data = append(data, &views.InputOptions{
			Label: menu.Name,
			Value: menu.ID,
		})
	}
	return data
}

// DelRedisAdminMenuList 删除管理菜单列表
func (_AdminMenu *AdminMenu) DelRedisAdminMenuList() {
	_, _ = _AdminMenu.rdsConn.Do("HDEL", RedisAuthAdminMenu, _AdminMenu.adminId)
	_, _ = _AdminMenu.rdsConn.Do("HDEL", RedisAuthAdminRouter, _AdminMenu.adminId)
}

// recursiveMenuChildren 递归获取管理菜单
func (_AdminMenu *AdminMenu) recursiveMenuChildren(menuId uint, menuList []*adminsModel.AdminMenu, routerList []string) []*adminsModel.AdminMenuInfo {
	var data []*adminsModel.AdminMenuInfo
	for _, menu := range menuList {
		if menu.ParentId == menuId {
			if menu.Route == "" || utils.ArrayStringIndexOf(routerList, "/"+adminsModel.ServiceAdminAuthRouteName+menu.Route) > -1 {
				childrenList := _AdminMenu.recursiveMenuChildren(menu.ID, menuList, routerList)
				if menu.Route != "" || len(childrenList) > 0 {
					data = append(data, &adminsModel.AdminMenuInfo{
						Id:       menu.ID,
						Name:     menu.Name,
						Route:    menu.Route,
						Children: childrenList,
						Data:     menu.Data,
					})
				}
			}
		}
	}
	return data
}
