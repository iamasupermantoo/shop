package usersService

import (
	"github.com/gomodule/redigo/redis"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/database"
	"gofiber/app/module/views"
	"strconv"
)

const (
	RedisUserSetting = "RedisUserSetting"
)

type UserSetting struct {
	rdsConn redis.Conn
	adminId uint
	userId  uint
}

func NewUserSetting(rdsConn redis.Conn, adminId, userId uint) *UserSetting {
	return &UserSetting{rdsConn: rdsConn, adminId: adminId, userId: userId}
}

// GetRedisUserSettingField 获取用户设置字段值
func (_UserSetting *UserSetting) GetRedisUserSettingField(field string) string {
	value, _ := redis.String(_UserSetting.rdsConn.Do("HGET", RedisUserSetting+strconv.Itoa(int(_UserSetting.userId)), field))
	if value == "" {
		settingInfo := &usersModel.Setting{}
		database.Db.Where("user_id = ?", _UserSetting.userId).Where("field = ?", field).Find(settingInfo)
		_, _ = _UserSetting.rdsConn.Do("HSET", RedisUserSetting+strconv.Itoa(int(_UserSetting.userId)), field, settingInfo.Value)
		value = settingInfo.Value
	}
	return value
}

// UserSettingUpdate 用户设置更新或创建
func (_UserSetting *UserSetting) UserSettingUpdate(field string, value interface{}) error {
	settingInfo := &usersModel.Setting{}
	result := database.Db.Where("user_id = ?", _UserSetting.userId).Where("field = ?", field).Find(settingInfo)
	if settingInfo.ID == 0 {
		// 找到模版对应的值
		settingTmp := &usersModel.Setting{}
		database.Db.Model(settingTmp).Where("field = ?", field).Where("admin_id = ?", 0).Find(settingTmp)
		database.Db.Model(&usersModel.Setting{}).Create(map[string]interface{}{
			"admin_id": _UserSetting.adminId, "user_id": _UserSetting.userId, "field": field, "value": value, "name": settingTmp.Name,
			"type": settingTmp.Type,
		})
		return result.Error
	}
	result = database.Db.Model(&usersModel.Setting{}).Where("id = ?", settingInfo.ID).Update("value", value)
	return result.Error
}

// DelRedisUserSettingField 删除用户设置缓存
func (_UserSetting *UserSetting) DelRedisUserSettingField(field string) {
	_, _ = _UserSetting.rdsConn.Do("HDEL", RedisUserSetting+strconv.Itoa(int(_UserSetting.userId)), field)
}

// GetDefaultInputViews 获取用户设置显示配置
func (_UserSetting *UserSetting) GetDefaultInputViews() (map[string]interface{}, *views.InputViews) {
	data := views.NewInputViews()

	userSettingList := make([]usersModel.Setting, 0)
	database.Db.Where("admin_id = ?", _UserSetting.adminId).
		Find(&userSettingList)

	params := map[string]interface{}{}
	for _, settingInfo := range userSettingList {
		params[settingInfo.Field] = views.InputViewsStringToData(settingInfo.Type, settingInfo.Value)
		data.SetInput(settingInfo.Type, settingInfo.Name, settingInfo.Field, views.InputViewsStringToData(settingInfo.Type, settingInfo.Data))
	}

	return params, data
}
