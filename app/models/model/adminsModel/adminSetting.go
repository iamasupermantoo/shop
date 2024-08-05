package adminsModel

import (
	"gofiber/app/models/model/types"
	"gofiber/app/module/views"
)

const (
	// AdminSettingGroupDefault 默认分类
	AdminSettingGroupDefault = -1

	// AdminSettingGroupBasic 基础分组
	AdminSettingGroupBasic = 1

	// AdminSettingGroupWallet 钱包分组
	AdminSettingGroupWallet = 2

	// AdminSettingGroupTemplate 模版配置
	AdminSettingGroupTemplate = 3
)

// AdminSetting 管理设置表
type AdminSetting struct {
	types.GormModel
	AdminId uint   `gorm:"type:int unsigned not null;comment:管理ID" json:"adminId"`
	GroupId int    `gorm:"type:int not null;comment:分组ID" json:"groupId"`
	Name    string `gorm:"type:varchar(60) not null;comment:设置名称" json:"name"`
	Type    int    `gorm:"type:tinyint not null;default:1;comment:类型" json:"type"`
	Field   string `gorm:"type:varchar(60) not null;comment:建铭" json:"field"`
	Value   string `gorm:"type:text;comment:键值" json:"value"`
	Data    string `gorm:"type:text;comment:input配置" json:"data"`
}

// AdminSettingRange 金额范围设置
type AdminSettingRange struct {
	Max float64 `json:"max"` // 最大值
	Min float64 `json:"min"` // 最小值
}

// AdminSettingRegisterAward 注册奖励设置
type AdminSettingRegisterAward struct {
	Register float64 `json:"register"` //	注册奖励
	Share    float64 `json:"share"`    // 分享者奖励
}

// AdminSettingDownload 下载地址设置
type AdminSettingDownload struct {
	Android string `json:"android"` // 安卓地址
	Ios     string `json:"apple"`   // 苹果地址
}

// AdminSettingSiteInfo 站点信息设置
type AdminSettingSiteInfo struct {
	Introduce string `json:"introduce"` // 站点说明
	Notice    string `json:"notice"`    // 弹窗公告
}

// AdminSettingWithdraw 提现设置
type AdminSettingWithdraw struct {
	Days int     `json:"days"` //	天数
	Nums int     `json:"nums"` //	次数
	Fee  float64 `json:"fee"`  //	手续费
}

type AdminSettingEarningsSetting struct {
	Options []*views.InputCheckboxOptions `json:"options"` //	选项
	Rate    float64                       `json:"rate"`    //	收益
}

// AdminSettingRsaInfo Rsa配置
type AdminSettingRsaInfo struct {
	PriKey string `json:"priKey"` //	私钥
	PubKey string `json:"pubKey"` //	公钥
}

// AdminSettingServiceRsa 管理服务器RSA
type AdminSettingServiceRsa struct {
	Admin *AdminSettingRsaInfo `json:"admin"` //	后台RSA
	Home  *AdminSettingRsaInfo `json:"home"`  //	前台RSA
}
