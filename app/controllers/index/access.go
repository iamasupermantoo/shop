package index

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/utils"
)

type AccessParams struct {
	Type     int    `validate:"required" json:"type"` // 记录类型
	SourceId uint   `json:"sourceId"`                 // 来源ID
	Name     string `json:"name"`                     // 记录名称
}

// Access 访问记录
func Access(ctx *context.CustomCtx, params *AccessParams) error {
	adminId := ctx.AdminId
	if adminId > 0 {
		database.Db.Create(&usersModel.Access{
			Name:     params.Name,
			AdminId:  adminId,
			SourceId: params.SourceId,
			UserId:   ctx.UserId,
			Type:     params.Type,
			Ip:       utils.GetClientIP(ctx.Ctx),
			Headers:  utils.JsonMarshal(ctx.GetReqHeaders()),
			Route:    ctx.GetReqHeaders()["Referer"][0],
		})
	}
	return ctx.SuccessJsonOK()
}
