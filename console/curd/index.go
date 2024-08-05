package curd

var indexTmp = `package {package}

import (
	"github.com/gofiber/fiber/v2"
	"{modelsPath}/{modelsPackage}"
	"gofiber/app/models/cacheRds"
	"gofiber/app/models/common"
	"gofiber/app/module/cache"
	"gofiber/app/module/database"
	"gofiber/utils"
)

type IndexParams struct {{modelsFieldStruct}` + "\t" + `Pagination *utils.Pagination      //	分页
}

// Index 管理列表
func Index(ctx *fiber.Ctx) error {
	params := ctx.Locals("params").(*IndexParams)
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	// 	只允许查询下级管理
	adminId, _ := utils.GetContextClaims(ctx)
	adminChildIds := cacheRds.RedisFindAdminChildrenIds(rdsConn, adminId)
	adminChildIds = append(adminChildIds, adminId)
	data := &common.IndexData{Items: make([]*{modelsPackage}.{modelsStruct}, 0)}

	//	过滤参数
	filterParams := utils.NewFilterParams()

	database.Db.Model(&{modelsPackage}.{modelsStruct}{}).Where("admin_id IN ?", adminChildIds).
		Scopes(filterParams.Scopes()).
		Count(&data.Count).
		Scopes(utils.Paginate(params.Pagination)).
		Find(&data.Items)

	return ctx.JSON(utils.SuccessJson(data))
}`
