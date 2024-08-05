package category

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// UpdateParams 更新参数
type UpdateParams struct {
	ID       uint   `gorm:"-" validate:"required" json:"id"` //	ID
	ParentId uint   `json:"parentId"`                        //  分类上级ID
	Type     int    `json:"type"`                            //  类型1默认类型
	Name     string `json:"name"`                            //  标题
	Icon     string `json:"icon"`                            //  封面
	Sort     int    `json:"sort"`                            // 排序
	Status   int    `json:"status"`                          //  状态-1禁用 10启用
}

// Update 更新接口
func Update(ctx *context.CustomCtx, params *UpdateParams) error {
	result := database.Db.Model(&productsModel.Category{}).
		Where("id = ?", params.ID).Where("admin_id IN ?", ctx.GetAdminChildIds()).
		Updates(params)
	return ctx.IsErrorJson(result.Error)
}
