package translate

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/utils"
)

type LangParams struct {
	Lang string `validate:"required" json:"lang"` // 语言标识
}

// Lang 语言翻译
func Lang(ctx *context.CustomCtx, params *LangParams) error {
	translateList := make([]*systemsModel.Translate, 0)
	database.Db.Model(&systemsModel.Translate{}).Where("lang = ?", "zh-CN").Where("admin_id IN ?", []uint{adminsModel.SuperAdminId, 0}).Find(&translateList)

	newTranslateList := make([]*systemsModel.Translate, 0)
	for _, currentTranslate := range translateList {
		if currentTranslate.Lang != params.Lang {
			// 如果存在字段, 那么不进行
			translateInfo := &systemsModel.Translate{}
			database.Db.Model(translateInfo).Where("lang = ?", params.Lang).Where("field = ?", currentTranslate.Field).Find(translateInfo)
			if translateInfo.ID == 0 {
				langValue, _ := utils.Translate(currentTranslate.Value, currentTranslate.Lang, params.Lang)
				newTranslateList = append(newTranslateList, &systemsModel.Translate{
					AdminId: currentTranslate.AdminId,
					Lang:    params.Lang,
					Name:    currentTranslate.Name,
					Type:    currentTranslate.Type,
					Field:   currentTranslate.Field,
					Value:   langValue,
				})
			}
		}
	}

	if len(newTranslateList) > 0 {
		database.Db.Create(newTranslateList)
	}
	return ctx.SuccessJsonOK()
}
