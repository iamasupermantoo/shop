package product

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/types"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

// CreateParams 新增参数
type CreateParams struct {
	CategoryId uint              `json:"categoryId"` //  类目ID
	AssetsId   uint              `json:"assetsId"`   //  资产ID
	Name       string            `json:"name"`       //  标题
	Images     types.GormStrings `json:"images"`     //  图标
	Money      float64           `json:"money"`      //  金额
	Type       int               `json:"type"`       //  类型1默认类型
}

// Create 新增接口
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	categoryInfo := &productsModel.Category{}
	result := database.Db.Model(categoryInfo).Where("id = ?", params.CategoryId).Find(categoryInfo)
	if result.Error != nil || categoryInfo.ID == 0 {
		return ctx.ErrorJson("找不到分类信息")
	}

	result = database.Db.Create(&productsModel.Product{
		AdminId:    categoryInfo.AdminId,
		CategoryId: categoryInfo.ID,
		AssetsId:   params.AssetsId,
		Name:       params.Name,
		Images:     params.Images,
		Money:      params.Money,
		Type:       params.Type,
	})

	return ctx.IsErrorJson(result.Error)
}
