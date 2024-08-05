package adminsModel

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"gofiber/app/models/model/types"
)

const (
	// AdminMenuStatusActive 开启
	AdminMenuStatusActive = 10

	// AdminMenuStatusDisable 禁用
	AdminMenuStatusDisable = -1

	// AdminMenuViewsTable 菜单视图模版表格
	AdminMenuViewsTable = "/views"

	// AdminMenuViewsSetting 菜单视图模版设置
	AdminMenuViewsSetting = "/setting"
)

// AdminMenu 管理员菜单
type AdminMenu struct {
	types.GormModel
	ParentId uint           `gorm:"type:int unsigned not null;comment:父级" json:"parentId"`
	Name     string         `gorm:"type:varchar(50) not null;comment:名称" json:"name"`
	Route    string         `gorm:"type:varchar(50) not null;comment:路由" json:"route"`
	Sort     int            `gorm:"type:tinyint not null;default:99;comment:排序" json:"sort"`
	Status   int            `gorm:"type:tinyint not null;default:10;comment:状态" json:"status"`
	Data     *AdminMenuData `gorm:"type:text;comment:数据" json:"data"`
}

// AdminMenuData 管理菜单配置
type AdminMenuData struct {
	Icon    string `json:"icon"`    //	图标
	Tmp     string `json:"tmp"`     //	显示模版
	ConfURL string `json:"confURL"` //	配置地址
}

// Value 设置数据
func (_AdminMenuData *AdminMenuData) Value() (driver.Value, error) {
	if _AdminMenuData == nil {
		return json.Marshal(&AdminMenuData{})
	}
	return json.Marshal(_AdminMenuData)
}

// Scan 查询数据
func (_AdminMenuData *AdminMenuData) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to scan AdminMenuData value:", value))
	}

	if len(bytes) > 0 {
		return json.Unmarshal(bytes, _AdminMenuData)
	}
	*_AdminMenuData = AdminMenuData{}
	return nil
}

// AdminMenuInfo 菜单列表
type AdminMenuInfo struct {
	Id       uint             `json:"id"`       //	ID
	Name     string           `json:"name"`     //	名称
	Route    string           `json:"route"`    //	路由
	Data     *AdminMenuData   `json:"data"`     //	数据
	Children []*AdminMenuInfo `json:"children"` //	子集
}
