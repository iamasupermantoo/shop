package userLevel

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/models/model/types"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// UpdateParams 更新参数
type UpdateParams struct {
	ID        uint                 `gorm:"-" validate:"required" json:"id"` //	ID
	Symbol    uint                 `json:"symbol"`                          // 等级ID
	Status    int                  `json:"status"`                          // 状态 -1禁用 10开启
	Increase  float64              `json:"increase"`                        // 涨幅
	ExpiredAt types.GormTimeParams // 过期时间
}

// Update 更新接口
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	adminChildIds := ctx.GetAdminChildIds()
	// 获取当前用户等级信息
	var currentSymbol int
	database.Db.Model(&usersModel.UserLevel{}).Select("symbol").Where(params.ID).Find(&currentSymbol)
	levelInfo := &systemsModel.Level{}
	result := database.Db.Model(levelInfo).Where("symbol = ?", currentSymbol).Where("admin_id IN ?", adminChildIds).Find(levelInfo)
	if result.Error != nil || levelInfo.ID == 0 {
		return ctx.ErrorJson("找不到当前等级信息")
	}

	result = database.Db.Model(&usersModel.UserLevel{}).
		Where("id = ?", params.ID).Where("admin_id IN ?", adminChildIds).
		Updates(params)
	return ctx.IsErrorJson(result.Error)
}
