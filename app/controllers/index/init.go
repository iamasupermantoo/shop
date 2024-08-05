package index

import (
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/models/service/systemsService"
	"gofiber/app/module/context"
	"strconv"
	"time"
)

// Init 初始化数据
func Init(ctx *context.CustomCtx, params *context.NoRequestBody) error {
	// 如果当前域名解析失败, 那么返回当前上传的域名, 并且查看是否管理过期
	nowTime := time.Now().Unix()
	expireTime := adminsService.NewAdminUser(ctx.Rds, ctx.AdminSettingId).GetRedisExpiration()
	if ctx.AdminSettingId == 0 || expireTime < nowTime {
		panic(ctx.OriginHost)
	}

	settingService := adminsService.NewAdminSetting(ctx.Rds, ctx.AdminSettingId)
	amountRateStr := settingService.GetRedisAdminSettingField("amountRate")
	amountRate, _ := strconv.ParseFloat(amountRateStr, 64)

	// 选中默认语言
	currentLangList := systemsService.NewSystemLang(ctx.Rds, ctx.AdminSettingId).GetRedisAdminLangList()
	if ctx.Lang == "" && len(currentLangList) > 0 {
		ctx.Lang = currentLangList[0].Symbol
	}

	// 获取商户显示模版
	adminService := adminsService.NewAdminUser(ctx.Rds, ctx.AdminSettingId)
	adminData := adminService.GetRedisAdminData()
	sysMenu := systemsService.NewSystemMenu(ctx.Rds, ctx.AdminSettingId)
	data := &initData{
		Config: &initConfig{
			Name:     settingService.GetRedisAdminSettingField("siteName"),
			Logo:     settingService.GetRedisAdminSettingField("siteLogo"),
			Template: adminData.Template,
			Rate:     amountRate,
		},
		Setting: &initSetting{
			Basic:    settingService.CheckBoxToMaps("basicTemplate"),
			Auth:     settingService.CheckBoxToMaps("authTemplate"),
			Register: settingService.CheckBoxToMaps("registerTemplate"),
			Login:    settingService.CheckBoxToMaps("loginTemplate"),
			Wallets:  settingService.CheckBoxToMaps("walletsTemplate"),
			Home:     settingService.CheckBoxToMaps("homeTemplate"),
			Settled:  settingService.CheckBoxToMaps("settledTemplate"),
		},
		CountryList: systemsService.NewSystemCountry(ctx.Rds, ctx.AdminSettingId).GetRedisAdminCountryList(),
		LangList:    currentLangList,
		Translate:   systemsService.NewSystemTranslate(ctx.Rds, ctx.AdminSettingId).GetRedisAdminTranslateLangList(ctx.Lang),
		MenuList: &initMenuList{
			Navigation: sysMenu.GetRedisSystemMenuList(systemsModel.MenuTypeNavigation),
			Setting:    sysMenu.GetRedisSystemMenuList(systemsModel.MenuTypeSetting),
			More:       sysMenu.GetRedisSystemMenuList(systemsModel.MenuTypeMore),
			Merchant:   sysMenu.GetRedisSystemMenuList(systemsModel.MenuTypeStore),
		},
	}

	return ctx.SuccessJson(data)
}

type initData struct {
	Config      *initConfig                         `json:"config"`
	Setting     *initSetting                        `json:"setting"`
	CountryList []*systemsModel.SystemCountryInfo   `json:"countryList"`
	LangList    []*systemsModel.SystemLangInfo      `json:"langList"`
	Translate   []*systemsModel.SystemTranslateInfo `json:"translate"`
	MenuList    *initMenuList                       `json:"menuList"`
}

type initConfig struct {
	Name     string  `json:"name"`     //	项目名称
	Logo     string  `json:"logo"`     //	项目Logo
	Template string  `json:"template"` //	前台模版
	Rate     float64 `json:"rate"`     // 美金汇率
}

type initSetting struct {
	Basic    map[string]bool `json:"basic"`
	Auth     map[string]bool `json:"auth"`
	Register map[string]bool `json:"register"`
	Login    map[string]bool `json:"login"`
	Wallets  map[string]bool `json:"wallets"`
	Home     map[string]bool `json:"home"`
	Settled  map[string]bool `json:"settled"`
}

type initMenuList struct {
	Navigation []*systemsModel.SystemMenuInfo `json:"navigation"`
	Setting    []*systemsModel.SystemMenuInfo `json:"setting"`
	More       []*systemsModel.SystemMenuInfo `json:"more"`
	Merchant   []*systemsModel.SystemMenuInfo `json:"merchant"`
}
