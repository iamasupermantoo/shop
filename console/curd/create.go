package curd

var createTmp = `package {package}

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"{modelsPath}/{modelsPackage}"
	"gofiber/app/module/database"
	"gofiber/utils"
)

// CreateParams 新增参数
type CreateParams struct {{modelsFieldStruct}}

// Create 新增接口
func Create(ctx *fiber.Ctx) error {
	//adminId, _ := utils.GetContextClaims(ctx)
	params := ctx.Locals("params").(*CreateParams)

	createInfo := &{modelsPackage}.{modelsStruct}{{modelsFieldParams}}

	result := database.Db.Create(createInfo)
	if result.Error != nil {
		return ctx.JSON(utils.ErrorJson(errors.New("添加失败, 原因 => " + result.Error.Error())))
	}

	return ctx.JSON(utils.SuccessJson("ok"))
}`
