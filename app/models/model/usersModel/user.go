package usersModel

import (
	"gofiber/app/models/model/types"
	"time"
)

const (
	// UserStatusActive 激活
	UserStatusActive = 10

	// UserStatusDisable 冻结
	UserStatusDisable = -1

	// UserTypeVirtual 虚拟用户
	UserTypeVirtual = -1

	// UserTypeDefault 默认用户
	UserTypeDefault = 1

	// UserTypeChannel 渠道用户
	UserTypeChannel = 2

	// UserTypeLevel 会员用户
	UserTypeLevel = 10

	// UserSexDefault 未知性别
	UserSexDefault = 0

	// UserSexMale 性别男
	UserSexMale = 1

	// UserSexFemale 性别女
	UserSexFemale = 2
)

// User 用户表
type User struct {
	types.GormModel
	AdminId     uint      `gorm:"type:int unsigned not null;default:1;comment:管理ID" json:"adminId"`
	ChannelId   uint      `gorm:"type:int unsigned not null;comment:渠道ID" json:"channelId"`
	ParentId    uint      `gorm:"type:int unsigned not null;comment:父级ID" json:"parentId"`
	CountryId   uint      `gorm:"type:int unsigned not null;comment:国家ID" json:"countryId"`
	UserName    string    `gorm:"uniqueIndex;type:varchar(60) not null;comment:用户名" json:"userName"`
	NickName    string    `gorm:"type:varchar(60) not null;comment:昵称" json:"nickName"`
	Email       string    `gorm:"type:varchar(60);default:null;comment:邮箱" json:"email"`
	Telephone   string    `gorm:"type:varchar(50);default:null;comment:手机号码" json:"telephone"`
	Avatar      string    `gorm:"type:varchar(120) not null;comment:头像" json:"avatar"`
	Score       int       `gorm:"type:tinyint not null;default:100;comment:信用分" json:"score"`
	Sex         int       `gorm:"type:tinyint not null;comment:性别0未知 1男 2女" json:"sex"`
	Birthday    time.Time `gorm:"type:datetime(3);autoCreateTime;comment:生日" json:"birthday"`
	Password    string    `gorm:"type:varchar(120) not null;comment:密码" json:"-"`
	SecurityKey string    `gorm:"type:varchar(120) not null;comment:密钥" json:"-"`
	Money       float64   `gorm:"type:decimal(12,2) not null;comment:金额" json:"money"`
	Type        int       `gorm:"type:tinyint not null;default:1;comment:类型 -1虚拟用户 1默认用户 10会员用户" json:"type"`
	Status      int       `gorm:"type:smallint not null;default:10;comment:状态 -1冻结 10激活" json:"status"`
	Data        string    `gorm:"type:text;comment:数据" json:"data"`
	Desc        string    `gorm:"type:text;comment:详情" json:"desc"`
}

type UserInfo struct {
	ID       uint   `json:"id"`       // ID
	AdminId  uint   `json:"adminId"`  // 管理ID
	Avatar   string `json:"avatar"`   //头像
	UserName string `json:"userName"` // 用户名
	NickName string `json:"nickName"` // 昵称
	Email    string `json:"email"`    // 邮箱
}

func (*UserInfo) TableName() string {
	return "user"
}
