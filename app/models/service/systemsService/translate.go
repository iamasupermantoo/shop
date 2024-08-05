package systemsService

import (
	"github.com/goccy/go-json"
	"github.com/gomodule/redigo/redis"
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/module/database"
	"strconv"
)

const (
	RedisSystemTranslateField = "RedisSystemTranslateField"
	RedisSystemTranslate      = "RedisSystemTranslate"
)

type SystemTranslate struct {
	rdsConn redis.Conn
	adminId uint
}

func NewSystemTranslate(rdsConn redis.Conn, adminId uint) *SystemTranslate {
	return &SystemTranslate{rdsConn: rdsConn, adminId: adminId}
}

// GetRedisAdminTranslateLangField 获取语言字段翻译值
func (_SystemTranslate *SystemTranslate) GetRedisAdminTranslateLangField(acceptLang string, field string) string {
	redisName := RedisSystemTranslateField + "_" + strconv.Itoa(int(_SystemTranslate.adminId)) + "_" + acceptLang
	value, _ := redis.String(_SystemTranslate.rdsConn.Do("HGET", redisName, field))
	if value == "" {
		database.Db.Model(&systemsModel.Translate{}).Select("value").
			Where("admin_id = ?", _SystemTranslate.adminId).Where("lang = ?", acceptLang).Where("field = ?", field).
			Find(&value)
		if value != "" {
			_, _ = _SystemTranslate.rdsConn.Do("HSET", redisName, field, value)
		}
	}

	// 如果值为空, 那么使用键名
	if value == "" {
		return field
	}
	return value
}

// GetRedisAdminTranslateLangList 获取管理语言列表
func (_SystemTranslate *SystemTranslate) GetRedisAdminTranslateLangList(acceptLang string) []*systemsModel.SystemTranslateInfo {
	redisName := RedisSystemTranslate + strconv.Itoa(int(_SystemTranslate.adminId))
	translateBytes, err := redis.Bytes(_SystemTranslate.rdsConn.Do("HGET", redisName, acceptLang))
	translateList := make([]*systemsModel.SystemTranslateInfo, 0)

	if err != nil {
		database.Db.Model(&systemsModel.Translate{}).Where("admin_id = ?", _SystemTranslate.adminId).
			Where("type = ?", systemsModel.TranslateTypeFrontend).Where("lang = ?", acceptLang).Find(&translateList)

		if len(translateList) > 0 {
			translateBytes, _ = json.Marshal(translateList)
			_, _ = _SystemTranslate.rdsConn.Do("HSET", redisName, acceptLang, translateBytes)
		}
		return translateList
	}

	_ = json.Unmarshal(translateBytes, &translateList)
	return translateList
}

// DelRedisAdminTranslateLangList 删除管理语言列表缓存
func (_SystemTranslate *SystemTranslate) DelRedisAdminTranslateLangList(acceptLang string) {
	_, _ = _SystemTranslate.rdsConn.Do("HDEL", RedisSystemTranslate+strconv.Itoa(int(_SystemTranslate.adminId)), acceptLang)
}

// DelRedisAdminTranslate 删除管理语言缓存
func (_SystemTranslate *SystemTranslate) DelRedisAdminTranslate() {
	_, _ = _SystemTranslate.rdsConn.Do("DEL", RedisSystemTranslate+strconv.Itoa(int(_SystemTranslate.adminId)))
}
