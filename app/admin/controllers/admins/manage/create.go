package manage

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/models/service/consoleService"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/utils"
	"gorm.io/gorm"
	"time"
)

// CreateParams 新增管理参数
type CreateParams struct {
	UserName string `validate:"required,alphanum" json:"userName"` //	用户名
	Password string `validate:"required" json:"password"`          //	密码
	Role     string `validate:"required" json:"role"`              //	角色
}

// Create 新增管理
func Create(ctx *context.CustomCtx, params *CreateParams) error {
	// 如果是商户管理, 那么设置管理数据
	var merchantData *adminsModel.AdminData
	if ctx.AdminId == adminsModel.SuperAdminId {
		merchantData = adminsModel.NewMerchantData()
	}
	adminInfo := &adminsModel.AdminUser{
		ParentId:    ctx.AdminId,
		UserName:    params.UserName,
		Password:    utils.PasswordEncrypt(params.Password),
		SecurityKey: utils.PasswordEncrypt(params.Password),
		ExpiredAt:   time.Now().Add(30 * 24 * time.Hour),
		Data:        merchantData,
	}

	err := database.Db.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(adminInfo)
		if result.Error != nil {
			return ctx.ErrorJson(result.Error.Error())
		}
		authAssignmentInfo := adminsModel.AuthAssignment{AdminId: adminInfo.ID, Name: params.Role}
		result = tx.Create(&authAssignmentInfo)
		if result.Error != nil {
			return ctx.ErrorJson(result.Error.Error())
		}
		return nil
	})
	if err != nil {
		return err
	}

	// 重置商户设置
	if ctx.AdminId == adminsModel.SuperAdminId {
		_ = consoleService.NewMerchant(adminInfo.ID, []string{"product", "category"}).RunRest()
	}

	// 删除当前管理下级缓存
	adminsService.NewAdminUser(ctx.Rds, ctx.AdminId).DelRedisChildrenIds()
	return ctx.SuccessJsonOK()
}
