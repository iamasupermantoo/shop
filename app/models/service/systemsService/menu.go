package systemsService

import (
	"github.com/goccy/go-json"
	"github.com/gomodule/redigo/redis"
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/module/database"
	"strconv"
)

const (
	RedisSystemMenuList = "RedisSystemMenuList"
)

type SystemMenu struct {
	rdsConn redis.Conn
	adminId uint
}

func NewSystemMenu(rdsConn redis.Conn, adminId uint) *SystemMenu {
	return &SystemMenu{rdsConn: rdsConn, adminId: adminId}
}

// GetRedisSystemMenuList 获取前端菜单
func (_SystemMenu *SystemMenu) GetRedisSystemMenuList(menuType int) []*systemsModel.SystemMenuInfo {
	redisName := RedisSystemMenuList + strconv.Itoa(int(_SystemMenu.adminId))
	menuListBytes, err := redis.Bytes(_SystemMenu.rdsConn.Do("HGET", redisName))
	menuList := make([]*systemsModel.Menu, 0)
	if err != nil {
		database.Db.Model(&systemsModel.Menu{}).Where("admin_id = ?", _SystemMenu.adminId).
			Where("status = ?", systemsModel.MenuStatusActive).Order("sort asc").Find(&menuList)

		if len(menuList) > 0 {
			menuListBytes, _ = json.Marshal(menuList)
			_, _ = _SystemMenu.rdsConn.Do("HSET", redisName, menuType, menuListBytes)
		}
	}

	_ = json.Unmarshal(menuListBytes, &menuList)
	return _SystemMenu.systemMenuChildren(0, menuType, menuList)
}

// DelRedisSystemMenuList 删除用户菜单缓存
func (_SystemMenu *SystemMenu) DelRedisSystemMenuList() {
	_, _ = _SystemMenu.rdsConn.Do("DEL", RedisSystemMenuList+strconv.Itoa(int(_SystemMenu.adminId)))
}

func (_SystemMenu *SystemMenu) systemMenuChildren(menuId uint, menuType int, menuList []*systemsModel.Menu) []*systemsModel.SystemMenuInfo {
	var data []*systemsModel.SystemMenuInfo
	for _, menu := range menuList {
		if menu.ParentId == menuId && menu.Type == menuType {
			data = append(data, &systemsModel.SystemMenuInfo{
				Name:       menu.Name,
				Route:      menu.Route,
				Icon:       menu.Icon,
				ActiveIcon: menu.ActiveIcon,
				IsDesktop:  menu.IsDesktop,
				IsMobile:   menu.IsMobile,
				Children:   _SystemMenu.systemMenuChildren(menu.ID, menuType, menuList),
			})
		}
	}
	return data
}
