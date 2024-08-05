package product

import (
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"time"
)

// CreateCommentParams 新增参数
type CreateCommentParams struct {
	ID        uint      `json:"id" validate:"required"`  // 店铺产品Id
	UserId    int       `json:"userId" validate:"gt=0"`  // 用户Id
	Comment   string    `json:"comment"`                 // 评论内容
	Rating    float64   `json:"rating" validate:"lte=5"` // 评分
	Image     []string  `json:"image"`                   // 图片
	CreatedAt time.Time `json:"createdAt"`               // 评论时间
}

// CreateComment 新增接口
func CreateComment(ctx *context.CustomCtx, params *CreateCommentParams) error {
	userInfo := usersModel.User{}
	result := database.Db.Where("id = ?", params.UserId).Where("type = ?", usersModel.UserTypeVirtual).Find(&userInfo)
	if result.RowsAffected == 0 {
		return ctx.ErrorJson("找不到用户信息")
	}

	var storeId uint
	if err := database.Db.Model(&productsModel.Product{}).Select("store_id").Where("id = ?", params.ID).Scan(&storeId).Error; err != nil {
		return ctx.ErrorJson("找不到产品信息")
	}

	if params.CreatedAt.IsZero() {
		params.CreatedAt = time.Now()
	}

	database.Db.Create(&shopsModel.StoreComment{
		AdminId:   userInfo.AdminId,
		UserId:    userInfo.ID,
		StoreId:   storeId,
		ProductId: params.ID,
		Name:      params.Comment,
		Rating:    params.Rating,
		Status:    shopsModel.StoreCommentsStatusComplete,
		Images:    params.Image,
	})
	return ctx.SuccessJsonOK()
}
