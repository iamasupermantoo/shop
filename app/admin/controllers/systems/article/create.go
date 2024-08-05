package article

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gorm.io/gorm"
	"strconv"
)

// CreateParams 新增参数
type CreateParams struct {
	Image   string `json:"image"`                       // 图片
	Name    string `validate:"required" json:"name"`    // 标题
	Content string `validate:"required" json:"content"` // 内容
	Type    int    `validate:"required" json:"type"`    // 1基础文章
}

// Create 新增接口
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	return ctx.IsErrorJson(database.Db.Transaction(func(tx *gorm.DB) error {
		createInfo := &systemsModel.Article{
			AdminId: ctx.AdminId,
			Image:   params.Image,
			Name:    params.Name,
			Content: params.Content,
			Type:    params.Type,
		}

		result := tx.Create(createInfo)
		if result.Error != nil {
			return ctx.ErrorJson("添加失败, 原因 => " + result.Error.Error())
		}

		// 添加多语言
		result = tx.Create(&[]*systemsModel.Translate{
			{AdminId: ctx.AdminId, Lang: "zh-CN", Name: "文章标题" + strconv.Itoa(int(createInfo.ID)), Field: params.Name},
			{AdminId: ctx.AdminId, Lang: "zh-CN", Name: "文章内容" + strconv.Itoa(int(createInfo.ID)), Field: params.Content},
		})
		if result.Error != nil {
			return ctx.ErrorJson("添加失败, 原因 => " + result.Error.Error())
		}
		return nil
	}))
}
