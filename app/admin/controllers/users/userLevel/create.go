package userLevel

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"time"
)

// CreateParams 新增参数
type CreateParams struct {
	UserName string `json:"userName"` // 用户账户
	Symbol   uint   `json:"symbol"`   // 等级ID
}

// Create 新增接口
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	adminChildIds := ctx.GetAdminChildIds()
	userInfo := &usersModel.User{}
	result := database.Db.Model(userInfo).Where("user_name = ?", params.UserName).Where("admin_id IN ?", adminChildIds).Find(userInfo)
	if result.Error != nil || userInfo.ID == 0 {
		return ctx.ErrorJson("找不到当前用户账户")
	}

	levelInfo := &systemsModel.Level{}
	result = database.Db.Model(levelInfo).Where("symbol = ?", params.Symbol).
		Where("admin_id IN ?", adminsService.NewAdminUser(ctx.Rds, userInfo.AdminId).GetRedisChildrenIds()).
		Find(levelInfo)
	if result.Error != nil || levelInfo.ID == 0 {
		return ctx.ErrorJson("找不到当前等级信息")
	}

	// 如果当前用户已存在等级
	currentLevel := &usersModel.UserLevel{}
	database.Db.Model(currentLevel).Where("admin_id = ?", userInfo.AdminId).Where("user_id = ?", userInfo.ID).Find(currentLevel)
	if currentLevel.ID > 0 {
		return ctx.ErrorJson("当前用户已存在用户等级")
	}

	if levelInfo.Days == -1 {
		levelInfo.Days = 365
	}
	createInfo := &usersModel.UserLevel{
		AdminId:   userInfo.AdminId,
		UserId:    userInfo.ID,
		Name:      levelInfo.Name,
		Symbol:    levelInfo.Symbol,
		Icon:      levelInfo.Icon,
		Money:     levelInfo.Money,
		ExpiredAt: time.Now().Add(time.Duration(levelInfo.Days) * 24 * time.Hour),
	}
	result = database.Db.Create(createInfo)
	if result.Error != nil {
		return ctx.ErrorJson("添加失败, 原因 => " + result.Error.Error())
	}

	return ctx.SuccessJsonOK()
}
