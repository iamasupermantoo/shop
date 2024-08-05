package user

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

type homeUserInfo struct {
	usersModel.User
	AuthInfo  usersModel.UserAuth  `json:"authInfo" gorm:"foreignKey:UserId"`
	LevelInfo usersModel.UserLevel `json:"levelInfo" gorm:"foreignKey:UserId"`
	StoreInfo shopsModel.Store     `json:"storeInfo" gorm:"foreignKey:UserId"`
}

func (homeUserInfo) TableName() string {
	return "user"
}

// Info 管理信息
func Info(ctx *context.CustomCtx, params *context.NoRequestBody) error {
	userInfo := &homeUserInfo{}
	database.Db.Model(userInfo).Where("id = ?", ctx.UserId).Preload("AuthInfo").Preload("LevelInfo").Preload("StoreInfo").Find(userInfo)
	return ctx.SuccessJson(userInfo)
}
