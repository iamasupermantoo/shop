package settled

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// UpdateParams 更新参数
type UpdateParams struct {
	ID      uint   `json:"id" gorm:"-" validate:"required"` // ID
	Type    int    `json:"type"`                            //  类型 1营业执照
	Name    string `json:"name"`                            //  证件名字
	Number  string `json:"number"`                          //  证件号
	Photo1  string `json:"photo1"`                          //  证件正
	Photo2  string `json:"photo2"`                          //  证件反
	Contact string `json:"contact"`                         //  联系方式
	Data    string `json:"data"`                            //  数据
	Status  int    `json:"status"`                          //  状态  -1拒绝 10审核中 20 通过
}

// Update 更新接口
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	result := database.Db.Model(&shopsModel.StoreSettled{}).
		Where("id = ?", params.ID).Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Updates(params)
	return ctx.IsErrorJson(result.Error)
}
