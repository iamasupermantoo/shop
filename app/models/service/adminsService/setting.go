package adminsService

import (
	"github.com/goccy/go-json"
	"github.com/gomodule/redigo/redis"
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/module/database"
	"gofiber/app/module/views"
	"strconv"
)

const (
	// RedisTokenRsa Redis Rsa私钥
	RedisTokenRsa = "TokenRsa"

	RedisAdminSettingName = "AdminSetting"
)

type AdminSetting struct {
	rdsConn redis.Conn
	adminId uint
}

type valueStruct struct {
	Label string `json:"label"`
	Value *bool  `json:"value"`
	Field string `json:"field"`
}

func NewAdminSetting(rdsConn redis.Conn, adminId uint) *AdminSetting {
	return &AdminSetting{rdsConn: rdsConn, adminId: adminId}
}

// GetCheckBox 获取CheckBox 数据
func (_AdminSetting *AdminSetting) GetCheckBox(field string) []*views.InputCheckboxOptions {
	data := make([]*views.InputCheckboxOptions, 0)
	checkBoxStr := _AdminSetting.GetRedisAdminSettingField(field)
	_ = json.Unmarshal([]byte(checkBoxStr), &data)
	return data
}

// CheckBoxToMaps CheckBox转Map
func (_AdminSetting *AdminSetting) CheckBoxToMaps(field string) map[string]bool {
	// 查询field对应的value
	var valueJsonText string
	if result := database.Db.Model(&adminsModel.AdminSetting{}).
		Select("value").
		Where("admin_id = ?", _AdminSetting.adminId).
		Where("field = ?", field).
		Find(&valueJsonText); result.RowsAffected == 0 {
		return nil
	}

	var values []valueStruct
	err := json.Unmarshal([]byte(valueJsonText), &values)
	if err != nil {
		panic(err)
	}

	// 检查valueJsonText是否包含value这个字段
	for _, valueText := range values {
		if valueText.Value == nil {
			panic("admin_setting 's value don't exist")
		}
	}

	data := map[string]bool{}
	for _, checkBox := range _AdminSetting.GetCheckBox(field) {
		data[checkBox.Field] = checkBox.Value.(bool)
	}
	return data
}

// GetRegisterAward 获取注册奖励
func (_AdminSetting *AdminSetting) GetRegisterAward() *adminsModel.AdminSettingRegisterAward {
	data := &adminsModel.AdminSettingRegisterAward{}
	dataStr := _AdminSetting.GetRedisAdminSettingField("registerAward")
	if dataStr == "" {
		return data
	}

	_ = json.Unmarshal([]byte(dataStr), data)
	return data
}

// GetDownload 获取下载地址
func (_AdminSetting *AdminSetting) GetDownload() *adminsModel.AdminSettingDownload {
	data := &adminsModel.AdminSettingDownload{}
	dataStr := _AdminSetting.GetRedisAdminSettingField("download")
	if dataStr == "" {
		return data
	}

	_ = json.Unmarshal([]byte(dataStr), data)
	return data
}

// GetSiteInfo 获取站点信息
func (_AdminSetting *AdminSetting) GetSiteInfo() *adminsModel.AdminSettingSiteInfo {
	data := &adminsModel.AdminSettingSiteInfo{}
	dataStr := _AdminSetting.GetRedisAdminSettingField("siteInfo")
	if dataStr == "" {
		return data
	}

	_ = json.Unmarshal([]byte(dataStr), data)
	return data
}

// GetBanner 获取banner列表
func (_AdminSetting *AdminSetting) GetBanner() []string {
	data := make([]string, 0)
	bannerStr := _AdminSetting.GetRedisAdminSettingField("siteBanners")
	if bannerStr == "" {
		return data
	}
	_ = json.Unmarshal([]byte(bannerStr), &data)
	return data
}

// GetWithdrawAccountNums 获取提现账户数量
func (_AdminSetting *AdminSetting) GetWithdrawAccountNums() int64 {
	numsStr := _AdminSetting.GetRedisAdminSettingField("walletAccountNums")
	if numsStr == "" {
		return 0
	}
	nums, _ := strconv.ParseInt(numsStr, 10, 64)
	return nums
}

// GetRangeMoney 获取金额范围
func (_AdminSetting *AdminSetting) GetRangeMoney(field string) *adminsModel.AdminSettingRange {
	data := &adminsModel.AdminSettingRange{}
	dataStr := _AdminSetting.GetRedisAdminSettingField(field)
	if dataStr == "" {
		return data
	}

	_ = json.Unmarshal([]byte(dataStr), data)
	return data
}

// GetWithdrawSetting 获取提现设置
func (_AdminSetting *AdminSetting) GetWithdrawSetting() *adminsModel.AdminSettingWithdraw {
	data := &adminsModel.AdminSettingWithdraw{}
	dataStr := _AdminSetting.GetRedisAdminSettingField("walletWithdrawSetting")
	if dataStr == "" {
		return data
	}

	_ = json.Unmarshal([]byte(dataStr), data)
	return data
}

// GetEarningSetting 分销设置
func (_AdminSetting *AdminSetting) GetEarningSetting() *adminsModel.AdminSettingEarningsSetting {
	data := &adminsModel.AdminSettingEarningsSetting{}
	dataStr := _AdminSetting.GetRedisAdminSettingField("earningsSetting")
	if dataStr == "" {
		return data
	}

	_ = json.Unmarshal([]byte(dataStr), data)
	return data
}

// GetRedisAdminSettingField 获取管理设置缓存数据
func (_AdminSetting *AdminSetting) GetRedisAdminSettingField(field string) string {
	redisName := RedisAdminSettingName + strconv.Itoa(int(_AdminSetting.adminId))
	value, err := redis.String(_AdminSetting.rdsConn.Do("HGET", redisName, field))
	if err != nil {
		settingInfo := &adminsModel.AdminSetting{}
		database.Db.Model(settingInfo).Select("value").Where("field = ?", field).
			Where("admin_id = ?", _AdminSetting.adminId).Find(settingInfo)
		value = settingInfo.Value
		if settingInfo.Value != "" {
			_, _ = _AdminSetting.rdsConn.Do("HSET", redisName, field, value)
		}
	}
	return value
}

// GroupOptions 获取分组Options
func (_AdminSetting *AdminSetting) GroupOptions() []*views.InputOptions {
	options := make([]*views.InputOptions, 0)
	if _AdminSetting.adminId == adminsModel.SuperAdminId {
		options = append(options, &views.InputOptions{
			Label: "默认配置",
			Value: adminsModel.AdminSettingGroupDefault,
		})
	}

	options = append(options, []*views.InputOptions{
		{Label: "基础配置", Value: adminsModel.AdminSettingGroupBasic},
		{Label: "财务配置", Value: adminsModel.AdminSettingGroupWallet},
		{Label: "模版配置", Value: adminsModel.AdminSettingGroupTemplate},
	}...)

	return options
}

// DelRedisAdminSettingField 删除管理设置单字段缓存
func (_AdminSetting *AdminSetting) DelRedisAdminSettingField(field string) {
	_, _ = _AdminSetting.rdsConn.Do("HDEL", RedisAdminSettingName+strconv.Itoa(int(_AdminSetting.adminId)), field)
}

// DelRedisAdminSetting 删除所有管理设置
func (_AdminSetting *AdminSetting) DelRedisAdminSetting() {
	_, _ = _AdminSetting.rdsConn.Do("DEL", RedisAdminSettingName+strconv.Itoa(int(_AdminSetting.adminId)))
}
