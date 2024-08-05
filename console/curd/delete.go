package curd

var deleteTmp = `package {package}

import (
	"github.com/gofiber/fiber/v2"
	"{modelsPath}/{modelsPackage}"
	"gofiber/app/models/cacheRds"
	"gofiber/app/models/common"
	"gofiber/app/module/cache"
	"gofiber/app/module/database"
	"gofiber/utils"
)

// Delete 删除接口
func Delete(ctx *fiber.Ctx) error {
	params := ctx.Locals("params").(*common.DeleteParams)
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	// 	只允许删除当前管理和下级管理
	adminId, _ := utils.GetContextClaims(ctx)
	adminChildIds := cacheRds.RedisFindAdminChildrenIds(rdsConn, adminId)
	adminChildIds = append(adminChildIds, adminId)

	result := database.Db.Where("id IN ?", params.Ids).
		Where("admin_id IN ?", adminChildIds).Delete(&{modelsPackage}.{modelsStruct}{})

	return ctx.JSON(utils.IsErrorJson(result.Error))
}`
