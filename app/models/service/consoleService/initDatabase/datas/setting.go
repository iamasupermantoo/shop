package datas

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/views"
	"gofiber/utils"
)

func InitUserSetting() []*usersModel.Setting {
	return []*usersModel.Setting{
		{Name: "站点信息", Type: views.InputTypeJson, Field: "siteInfo", Value: utils.JsonMarshal(&adminsModel.AdminSettingSiteInfo{Introduce: "", Notice: ""}),
			Data: utils.JsonMarshal([][]*views.InputAttrsViews{{
				{Label: "站点信息", Field: "introduce", Type: views.InputTypeText},
				{Label: "站点公告", Field: "notice", Type: views.InputTypeText},
			}}),
		},
	}
}
