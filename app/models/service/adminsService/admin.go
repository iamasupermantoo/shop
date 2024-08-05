package adminsService

import (
	"errors"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gomodule/redigo/redis"
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/module/database"
	"gofiber/app/module/views"
	"gofiber/utils"
	"strings"
)

const (
	// RedisAdminChildrenIds 缓存管理下级Ids
	RedisAdminChildrenIds = "RedisAdminChildrenIds"

	// RedisAdminDomainToAdminId 缓存管理域名
	RedisAdminDomainToAdminId = "RedisAdminDomainToAdminId"

	// RedisAdminToSettingAdminId 缓存配置的管理ID
	RedisAdminToSettingAdminId = "RedisAdminToSettingAdminId"

	// RedisAdminExpiration 缓存管理过期时间
	RedisAdminExpiration = "RedisAdminExpiration"

	// RedisAdminMerchantData 缓存管理商户数据配置
	RedisAdminMerchantData = "RedisAdminMerchantData"
)

type AdminUser struct {
	rdsConn redis.Conn
	adminId uint
}

func NewAdminUser(rdsConn redis.Conn, adminId uint) *AdminUser {
	return &AdminUser{rdsConn: rdsConn, adminId: adminId}
}

// RestAdminId 重置管理ID
func (_AdminUser *AdminUser) RestAdminId(adminId uint) *AdminUser {
	_AdminUser.adminId = adminId
	return _AdminUser
}

// SetAdminId 设置管理ID
func (_AdminUser *AdminUser) SetAdminId(adminId uint) *AdminUser {
	_AdminUser.adminId = adminId
	return _AdminUser
}

// GetRedisAdminData 获取缓存管理数据
func (_AdminUser *AdminUser) GetRedisAdminData() *adminsModel.AdminData {
	dataBytes, err := redis.Bytes(_AdminUser.rdsConn.Do("HGET", _AdminUser, _AdminUser.adminId))
	if err != nil {
		adminInfo := &adminsModel.AdminUser{}
		result := database.Db.Model(adminInfo).Where("id = ?", _AdminUser.adminId).Find(adminInfo)
		if result.Error != nil {
			return &adminsModel.AdminData{Template: adminsModel.AdminDefaultTemplate, AgentNums: adminsModel.AdminDefaultAgentNums}
		}

		// 设置当前管理数据缓存
		adminDataBytes, _ := json.Marshal(adminInfo.Data)
		_, _ = _AdminUser.rdsConn.Do("HSET", RedisAdminMerchantData, _AdminUser.adminId, adminDataBytes)
		return adminInfo.Data
	}

	adminData := &adminsModel.AdminData{}
	_ = json.Unmarshal(dataBytes, &adminData)
	return adminData
}

// GetRedisDomainAdminId 获取域名对应的管理ID
func (_AdminUser *AdminUser) GetRedisDomainAdminId(domain string) uint {
	adminId, err := redis.Uint64(_AdminUser.rdsConn.Do("HGET", RedisAdminDomainToAdminId, domain))
	if err != nil || adminId == 0 {
		adminInfo := &adminsModel.AdminUser{}
		result := database.Db.Model(adminInfo).Where("FIND_IN_SET(?, domains)", domain).Find(&adminInfo)
		if result.Error != nil {
			return 0
		}
		_, _ = _AdminUser.rdsConn.Do("HSET", RedisAdminDomainToAdminId, domain, adminInfo.ID)
		return adminInfo.ID
	}

	return uint(adminId)
}

// GetRedisAdminSettingId 获取管理设置ID
func (_AdminUser *AdminUser) GetRedisAdminSettingId(adminId uint) uint {
	var adminSettingId uint
	rdsAdminSettingId, err := redis.Uint64(_AdminUser.rdsConn.Do("HGET", RedisAdminToSettingAdminId, adminId))
	if err != nil {
		adminList := make([]*adminsModel.AdminUser, 0)
		database.Db.Model(&adminsModel.AdminUser{}).Find(&adminList)
		adminSettingId = _AdminUser.recursiveParent(adminId, adminList)

		if adminSettingId > 0 {
			_, _ = _AdminUser.rdsConn.Do("HSET", RedisAdminToSettingAdminId, adminId, adminSettingId)
		}
		return adminSettingId
	}
	return uint(rdsAdminSettingId)
}

// GetRedisChildrenIds 获取子集Ids
func (_AdminUser *AdminUser) GetRedisChildrenIds() []uint {
	dataBytes, err := redis.Bytes(_AdminUser.rdsConn.Do("HGET", RedisAdminChildrenIds, _AdminUser.adminId))
	adminIds := make([]uint, 0)
	if err != nil {
		adminList := make([]*adminsModel.AdminUser, 0)
		database.Db.Model(&adminsModel.AdminUser{}).Find(&adminList)

		// 递归查询
		adminIds = _AdminUser.recursiveChildren(_AdminUser.adminId, adminList)
		if len(adminIds) > 0 {
			dataBytes, _ = json.Marshal(adminIds)
			_, _ = _AdminUser.rdsConn.Do("HSET", RedisAdminChildrenIds, _AdminUser.adminId, dataBytes)
		}
	}
	_ = json.Unmarshal(dataBytes, &adminIds)

	// 加上自身
	adminIds = append(adminIds, _AdminUser.adminId)
	return adminIds
}

// GetAdminChildrenOptions 获取管理下级Options
func (_AdminUser *AdminUser) GetAdminChildrenOptions(adminIds []uint) []*views.InputOptions {
	adminList := make([]*adminsModel.AdminUser, 0)
	database.Db.Model(&adminsModel.AdminUser{}).Where("id IN ?", adminIds).Find(&adminList)

	data := make([]*views.InputOptions, 0)
	for _, adminInfo := range adminList {
		data = append(data, &views.InputOptions{Label: adminInfo.UserName, Value: adminInfo.ID})
	}
	return data
}

// UpdateDomains 更新域名
func (_AdminUser *AdminUser) UpdateDomains(oldDomains, newDomains string) error {
	domainList := strings.Split(newDomains, ",")
	for _, domain := range domainList {
		domainAdminInfo := &adminsModel.AdminUser{}
		database.Db.Model(domainAdminInfo).Where("domains LIKE ?", "%"+domain+"%").Find(domainAdminInfo)
		if domainAdminInfo.ID > 0 && domainAdminInfo.ID != _AdminUser.adminId {
			return errors.New("当前域名已被绑定 ===> " + domain)
		}

		_AdminUser.DelRedisDomainAdminId(domain)
	}

	// 删除之前域名绑定的管理ID
	currentDomainList := strings.Split(oldDomains, ",")
	for _, domain := range currentDomainList {
		_AdminUser.DelRedisDomainAdminId(domain)
	}
	return nil
}

// GetRedisExpiration 获取管理过期时间
func (_AdminUser *AdminUser) GetRedisExpiration() int64 {
	expireTime, _ := redis.Int64(_AdminUser.rdsConn.Do("HGET", RedisAdminExpiration, _AdminUser.adminId))
	if expireTime == 0 {
		adminInfo := &adminsModel.AdminUser{}
		database.Db.Model(adminInfo).Where("id = ?", _AdminUser.adminId).Find(adminInfo)
		expireTime = adminInfo.ExpiredAt.Unix()
		_, _ = _AdminUser.rdsConn.Do("HSET", RedisAdminExpiration, _AdminUser.adminId, expireTime)
	}
	return expireTime
}

// VerifyWhitelist 验证白名单
func (_AdminUser *AdminUser) VerifyWhitelist(ctx *fiber.Ctx) bool {
	adminSettingId := _AdminUser.GetRedisAdminSettingId(_AdminUser.adminId)
	merchantData := _AdminUser.SetAdminId(adminSettingId).GetRedisAdminData()
	if merchantData.Whitelist == "" {
		return true
	}

	// 判断是不是设置了白名单
	whitelistList := strings.Split(merchantData.Whitelist, ",")
	currentIP := utils.GetClientIP(ctx)
	return utils.ArrayStringIndexOf(whitelistList, currentIP) > -1
}

// recursiveChildren 递归查询子集
func (_AdminUser *AdminUser) recursiveChildren(adminId uint, adminList []*adminsModel.AdminUser) []uint {
	ids := make([]uint, 0)
	if adminId == 0 {
		return ids
	}

	for _, adminInfo := range adminList {
		if adminInfo.ParentId == adminId {
			ids = append(ids, adminInfo.ID)
			ids = append(ids, _AdminUser.recursiveChildren(adminInfo.ID, adminList)...)
		}
	}
	return ids
}

// recursiveParent 递归获取父级ID
func (_AdminUser *AdminUser) recursiveParent(adminId uint, adminList []*adminsModel.AdminUser) uint {
	var adminSettingId uint
	for _, adminInfo := range adminList {
		if adminInfo.ID == adminId {
			if adminInfo.ParentId == adminsModel.SuperAdminId || adminInfo.ParentId == 0 {
				adminSettingId = adminInfo.ID
			} else {
				adminSettingId = _AdminUser.recursiveParent(adminInfo.ParentId, adminList)
			}
			break
		}

	}
	return adminSettingId
}

// DelRedisChildrenIds 删除管理下级Ids缓存
func (_AdminUser *AdminUser) DelRedisChildrenIds() {
	_, _ = _AdminUser.rdsConn.Do("HDEL", RedisAdminChildrenIds, _AdminUser.adminId)
}

// DelRedisExpiration 删除管理过期时间
func (_AdminUser *AdminUser) DelRedisExpiration() {
	_, _ = _AdminUser.rdsConn.Do("HDEL", RedisAdminExpiration, _AdminUser.adminId)
}

// DelRedisAdminData 删除管理数据缓存
func (_AdminUser *AdminUser) DelRedisAdminData() {
	_, _ = _AdminUser.rdsConn.Do("HDEL", RedisAdminMerchantData, _AdminUser.adminId)
}

// DelRedisDomainAdminId 删除缓存对应的管理ID
func (_AdminUser *AdminUser) DelRedisDomainAdminId(domain string) {
	_, _ = _AdminUser.rdsConn.Do("HDEL", RedisAdminDomainToAdminId, domain)
}

func (_AdminUser *AdminUser) DelRedisAdminSettingId() {
	_, _ = _AdminUser.rdsConn.Do("HDEL", RedisAdminToSettingAdminId, _AdminUser.adminId)
}
