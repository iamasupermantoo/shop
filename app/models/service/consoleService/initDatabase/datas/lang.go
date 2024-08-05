package datas

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/systemsModel"
)

// InitLang 初始化语言
func InitLang() []*systemsModel.Lang {
	return []*systemsModel.Lang{
		{AdminId: adminsModel.SuperAdminId, Name: "简体中文", Alias: "简体中文", Symbol: "zh-CN", Icon: "/assets/country/china.png", Sort: 1, Status: 10, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "繁体中文", Alias: "繁體中文", Symbol: "zh-TW", Icon: "/assets/country/taiwan.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "英格兰语", Alias: "English", Symbol: "en-US", Icon: "/assets/country/usa.png", Sort: 99, Status: 10, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "阿拉伯语", Alias: "عربي", Symbol: "ar-AE", Icon: "/assets/country/united_arab_emirates.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "白俄罗斯语", Alias: "беларускі", Symbol: "be-BY", Icon: "/assets/country/belarus.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "保加利亚语", Alias: "български", Symbol: "bg-BG", Icon: "/assets/country/bulgaria.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "捷克语", Alias: "čeština", Symbol: "cs-CZ", Icon: "/assets/country/czech.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "丹麦语", Alias: "dansk", Symbol: "da-DK", Icon: "/assets/country/denmark.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "德语", Alias: "Deutsch", Symbol: "de-DE", Icon: "/assets/country/germany.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "希腊语", Alias: "Ελληνικά", Symbol: "el-GR", Icon: "/assets/country/greece.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "西班牙语", Alias: "español", Symbol: "es-ES", Icon: "/assets/country/spain.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "爱沙尼亚语", Alias: "eestikeel", Symbol: "et-EE", Icon: "/assets/country/estonia.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "芬兰语", Alias: "Suomalainen", Symbol: "fi-FI", Icon: "/assets/country/finland.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "法语", Alias: "Français", Symbol: "fr-FR", Icon: "/assets/country/france.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "克罗地亚语", Alias: "Hrvatski", Symbol: "hr-HR", Icon: "/assets/country/croatia.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "匈牙利语", Alias: "Magyar", Symbol: "hu-HU", Icon: "/assets/country/hungary.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "冰岛语", Alias: "íslenskur", Symbol: "is-IS", Icon: "/assets/country/iceland.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "意大利语", Alias: "italiano", Symbol: "it-IT", Icon: "/assets/country/italy.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "日语", Alias: "日本", Symbol: "ja-JP", Icon: "/assets/country/japan.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "马来语", Alias: "Melayu", Symbol: "ms-MY", Icon: "/assets/country/malaysia.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "越南语", Alias: "TiếngViệt", Symbol: "vi-VN", Icon: "/assets/country/vietnam.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "韩语", Alias: "한국인", Symbol: "ko-KR", Icon: "/assets/country/north_korea.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "立陶宛语", Alias: "lietuvių", Symbol: "lt-LT", Icon: "/assets/country/lithuania.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "马其顿语", Alias: "македонски", Symbol: "mk-MK", Icon: "/assets/country/macedonia.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "荷兰语", Alias: "Nederlands", Symbol: "nl-NL", Icon: "/assets/country/netherlands.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "挪威语", Alias: "norsk", Symbol: "no-NO", Icon: "/assets/country/norway.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "波兰语", Alias: "Polski", Symbol: "pl-PL", Icon: "/assets/country/poland.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "葡萄牙语", Alias: "Português", Symbol: "pt-PT", Icon: "/assets/country/portugal.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "罗马尼亚语", Alias: "Română", Symbol: "ro-RO", Icon: "/assets/country/romania.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "俄语", Alias: "Русский", Symbol: "ru-RU", Icon: "/assets/country/russia.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "克罗地亚语", Alias: "Hrvatski", Symbol: "sh-YU", Icon: "/assets/country/croatia.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "斯洛伐克语", Alias: "slovenský", Symbol: "sk-SK", Icon: "/assets/country/slovakia.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "斯洛文尼亚语", Alias: "Slovenščina", Symbol: "sl-SI", Icon: "/assets/country/slovenia.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "阿尔巴尼亚语", Alias: "shqiptare", Symbol: "sq-AL", Icon: "/assets/country/albania.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "瑞典语", Alias: "svenska", Symbol: "sv-SE", Icon: "/assets/country/sweden.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "泰语", Alias: "แบบไทย", Symbol: "th-TH", Icon: "/assets/country/thailand.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "土耳其语", Alias: "Türkçe", Symbol: "tr-TR", Icon: "/assets/country/turkey.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "乌克兰语", Alias: "українська", Symbol: "uk-UA", Icon: "/assets/country/ukraine.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "塞尔维亚语", Alias: "Српски", Symbol: "sr-YU", Icon: "/assets/country/serbia.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "希伯来语", Alias: "עִברִית", Symbol: "iw-IL", Icon: "/assets/country/israel.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "印地语", Alias: "हिंदी", Symbol: "hi-IN", Icon: "/assets/country/india.png", Sort: 99, Status: -1, Data: ""},
		{AdminId: adminsModel.SuperAdminId, Name: "印尼语", Alias: "Indonesia", Symbol: "id-ID", Icon: "/assets/country/indonesia.png", Sort: 99, Status: -1, Data: ""},
	}
}
