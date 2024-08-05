package adminsModel

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"gofiber/app/models/model/types"
	"gofiber/utils"
	"time"
)

const (
	// SuperAdminId 超级管理ID
	SuperAdminId = 1

	// MerchantAdminId 商户管理ID
	MerchantAdminId = 2

	// AgentAdminId 代理管理员ID
	AgentAdminId = 3

	// ServiceAdminRouteName 后端项目路由名称
	ServiceAdminRouteName = "admin"

	// ServiceAdminAuthRouteName 后端验证路由名称
	ServiceAdminAuthRouteName = "admin/auth"

	//	ServiceHomeName 前台项目名称
	ServiceHomeName = "home"

	// AdminDefaultTemplate 默认模版
	AdminDefaultTemplate = "default"

	// AdminDefaultAgentNums 管理代理数量
	AdminDefaultAgentNums = 5

	// AdminUserStatusActive 激活
	AdminUserStatusActive = 10

	// AdminUserStatusDisable 冻结
	AdminUserStatusDisable = -1
)

// AdminUser 管理表
type AdminUser struct {
	types.GormModel
	ParentId    uint       `gorm:"type:int unsigned not null;comment:上级ID" json:"parentId"`
	UserName    string     `gorm:"uniqueIndex;type:varchar(60) not null;comment:用户名" json:"userName"`
	NickName    string     `gorm:"type:varchar(60) not null;comment:昵称" json:"nickName"`
	Email       string     `gorm:"type:varchar(60) not null;comment:邮箱" json:"email"`
	Avatar      string     `gorm:"type:varchar(120) not null;comment:头像" json:"avatar"`
	Password    string     `gorm:"type:varchar(120) not null;comment:密码" json:"-"`
	SecurityKey string     `gorm:"type:varchar(120) not null;comment:密钥" json:"-"`
	Money       float64    `gorm:"type:decimal(12,2) not null;comment:金额" json:"money"`
	Domains     string     `gorm:"type:varchar(1020) not null;comment:绑定域名" json:"domains"`
	SeatLink    string     `gorm:"type:varchar(255) not null;comment:坐席链接" json:"seatLink"`
	Online      string     `gorm:"type:varchar(255) not null;comment:客服链接" json:"online"`
	Status      int        `gorm:"type:smallint not null;default:10;comment:状态 10激活 -1冻结" json:"status"`
	Data        *AdminData `gorm:"type:text;comment:数据" json:"data"`
	ExpiredAt   time.Time  `gorm:"type:datetime(3);comment:过期时间" json:"expiredAt"`
}

// AdminData 管理数据
type AdminData struct {
	Key       string `json:"key"`       //	授权Key
	Template  string `json:"template"`  //	前端模版
	AgentNums int    `json:"agentNums"` //	代理数量
	Whitelist string `json:"whitelist"` //	登录白名单
}

// NewMerchantData 创建商户数据
func NewMerchantData() *AdminData {
	return &AdminData{Key: utils.NewRandom().String(32), Template: AdminDefaultTemplate, AgentNums: AdminDefaultAgentNums, Whitelist: ""}
}

// Value 设置数据
func (_AdminData *AdminData) Value() (driver.Value, error) {
	if _AdminData == nil {
		return []byte(""), nil
	}
	return json.Marshal(_AdminData)
}

// Scan 查询数据
func (_AdminData *AdminData) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to scan GormAdminData value:", value))
	}

	if len(bytes) > 0 {
		return json.Unmarshal(bytes, _AdminData)
	}
	*_AdminData = AdminData{}
	return nil
}
