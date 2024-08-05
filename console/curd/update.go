package curd

var updateTmp = `package {package}

import (
	"github.com/gofiber/fiber/v2"
	"{modelsPath}/{modelsPackage}"
	"gofiber/app/models/cacheRds"
	"gofiber/app/module/cache"
	"gofiber/app/module/database"
	"gofiber/utils"
)

// UpdateParams 更新参数
type UpdateParams struct {
	ID          uint    ` + "`" + `gorm:"-" validate:"required"` + "`" + ` //	ID{modelsFieldStruct}}

// Update 更新接口
func Update(ctx *fiber.Ctx) error {
	params := ctx.Locals("params").(*UpdateParams)
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	// 	只允许更新当前管理和下级管理
	adminId, _ := utils.GetContextClaims(ctx)
	adminChildIds := cacheRds.RedisFindAdminChildrenIds(rdsConn, adminId)
	adminChildIds = append(adminChildIds, adminId)

	result := database.Db.Model(&{modelsPackage}.{modelsStruct}{}).
		Where("id = ?", params.ID).Where("admin_id IN ?", adminChildIds).
		Updates(params)
	return ctx.JSON(utils.IsErrorJson(result.Error))
}`
