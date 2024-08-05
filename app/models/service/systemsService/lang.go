package systemsService

import (
	"github.com/goccy/go-json"
	"github.com/gomodule/redigo/redis"
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/module/database"
	"gofiber/app/module/views"
	"strconv"
)

const (
	RedisSystemLangList = "RedisSystemLangList"
)

type SystemLang struct {
	rdsConn redis.Conn
	adminId uint
}

func NewSystemLang(rdsConn redis.Conn, adminId uint) *SystemLang {
	return &SystemLang{rdsConn: rdsConn, adminId: adminId}
}

// GetRedisAdminLangList 获取管理语言列表
func (_SystemLang *SystemLang) GetRedisAdminLangList() []*systemsModel.SystemLangInfo {
	dataBytes, err := redis.Bytes(_SystemLang.rdsConn.Do("HGET", RedisSystemLangList, _SystemLang.adminId))
	currentLangList := make([]*systemsModel.SystemLangInfo, 0)

	if err != nil {
		database.Db.Model(&systemsModel.Lang{}).Where("admin_id = ?", _SystemLang.adminId).
			Where("status = ?", systemsModel.LangStatusActive).Order("sort ASC").
			Find(&currentLangList)
		if len(currentLangList) > 0 {
			dataBytes, _ = json.Marshal(currentLangList)
			_, _ = _SystemLang.rdsConn.Do("HSET", RedisSystemLangList, _SystemLang.adminId, dataBytes)
		}
		return currentLangList
	}

	_ = json.Unmarshal(dataBytes, &currentLangList)
	return currentLangList
}

// GetAdminOptions 获取管理Options
func (_SystemLang *SystemLang) GetAdminOptions(adminIds []uint) []*views.InputOptions {
	langList := make([]*systemsModel.Lang, 0)
	database.Db.Model(&systemsModel.Lang{}).Where("admin_id in ?", adminIds).Where("status = ?", systemsModel.LangStatusActive).Find(&langList)

	data := make([]*views.InputOptions, 0)
	for _, lang := range langList {
		data = append(data, &views.InputOptions{
			Label: lang.Symbol + "[" + lang.Name + "]" + " - [管理ID:" + strconv.Itoa(int(lang.AdminId)) + "]",
			Value: lang.Symbol,
		})
	}
	return data
}

// DelRedisAdminLangList 删除管理语言列表缓存
func (_SystemLang *SystemLang) DelRedisAdminLangList() {
	_, _ = _SystemLang.rdsConn.Do("HDEL", RedisSystemLangList, _SystemLang.adminId)
}
