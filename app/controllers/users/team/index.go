package team

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"time"
)

// Index 我的团队
func Index(ctx *context.CustomCtx, params *context.NoRequestBody) error {
	userId := ctx.UserId
	userTeamInfo := userTeam{}
	database.Db.Model(&usersModel.User{}).Where("parent_id = ?", userId).Count(&userTeamInfo.Nums)
	database.Db.Model(&usersModel.User{}).Where("id = ?", userId).Preload("Children").Find(&userTeamInfo)

	for _, child := range userTeamInfo.Children {
		database.Db.Model(&usersModel.User{}).Where("parent_id = ?", child.ID).Count(&child.Nums)
	}

	return ctx.SuccessJson(userTeamInfo)
}

type userTeam struct {
	usersModel.User
	Nums     int64          `json:"nums" gorm:"-"` // 邀请人数
	Children []userChildren `json:"children" gorm:"foreignKey:ParentId"`
}

func (userTeam) TableName() string {
	return "user"
}

type userChildren struct {
	ID        uint      `json:"iD"`            // 用户ID
	ParentId  uint      `json:"parentId"`      // 上级ID
	Avatar    string    `json:"avatar"`        // 头像
	UserName  string    `json:"userName"`      // 用户账户
	Nums      int64     `json:"nums" gorm:"-"` // 邀请人数
	CreatedAt time.Time `json:"createdAt"`     // 时间
}

func (userChildren) TableName() string {
	return "user"
}
