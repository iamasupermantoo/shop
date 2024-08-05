package comment

import (
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
)

type CreateParams struct {
	ID     uint     `json:"id" validate:"required"`     //	评论ID
	Name   string   `json:"name" validate:"required"`   //	评论内容
	Images []string `json:"images"`                     // 评论图片
	Rating float64  `json:"rating" validate:"required"` //	评分
}

// Create 评论产品
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	commentInfo := &shopsModel.StoreComment{}
	result := database.Db.Model(commentInfo).Where("id = ?", params.ID).Where("user_id = ?", ctx.UserId).Find(commentInfo)
	if result.Error != nil || commentInfo.ID == 0 || commentInfo.Status != shopsModel.StoreCommentsStatusPending {
		return ctx.ErrorJsonTranslate("findError")
	}

	database.Db.Model(&shopsModel.StoreComment{}).Where("id = ?", commentInfo.ID).Updates(&shopsModel.StoreComment{
		Name:   params.Name,
		Images: params.Images,
		Rating: params.Rating,
		Status: shopsModel.StoreCommentsStatusComplete,
	})

	return ctx.SuccessJsonOK()
}
