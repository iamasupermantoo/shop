package systemsService

import (
	"github.com/goccy/go-json"
	"github.com/gomodule/redigo/redis"
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/module/database"
)

const (
	RedisSystemCountryList = "RedisSystemCountryList"
)

type SystemCountry struct {
	rdsConn redis.Conn
	adminId uint
}

func NewSystemCountry(rdsConn redis.Conn, adminId uint) *SystemCountry {
	return &SystemCountry{rdsConn: rdsConn, adminId: adminId}
}

// GetRedisAdminCountryList 获取管理国家列表
func (_SystemCountry *SystemCountry) GetRedisAdminCountryList() []*systemsModel.SystemCountryInfo {
	countryBytes, err := redis.Bytes(_SystemCountry.rdsConn.Do("HGET", RedisSystemCountryList, _SystemCountry.adminId))
	countryList := make([]*systemsModel.SystemCountryInfo, 0)
	if err != nil {
		database.Db.Model(&systemsModel.Country{}).Where("admin_id = ?", _SystemCountry.adminId).
			Where("status = ?", systemsModel.CountryStatusActive).Order("sort ASC").
			Find(&countryList)
		if len(countryList) > 0 {
			countryBytes, _ = json.Marshal(countryList)
			_, _ = _SystemCountry.rdsConn.Do("HSET", RedisSystemCountryList, _SystemCountry.adminId, countryBytes)
		}
		return countryList
	}

	_ = json.Unmarshal(countryBytes, &countryList)
	return countryList
}

// DelRedisAdminCountryList 删除管理国家列表缓存
func (_SystemCountry *SystemCountry) DelRedisAdminCountryList() {
	_, _ = _SystemCountry.rdsConn.Do("HDEL", RedisSystemCountryList, _SystemCountry.adminId)
}
