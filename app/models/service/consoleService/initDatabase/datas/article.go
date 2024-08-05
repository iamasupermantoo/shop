package datas

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/systemsModel"
)

func InitArticle() []*systemsModel.Article {
	return []*systemsModel.Article{
		{AdminId: adminsModel.SuperAdminId, Type: systemsModel.ArticleTypeHelpers, Name: "helperLabel1", Content: "helperContent1"},
		{AdminId: adminsModel.SuperAdminId, Type: systemsModel.ArticleTypeHelpers, Name: "helperLabel2", Content: "helperContent2"},
		{AdminId: adminsModel.SuperAdminId, Type: systemsModel.ArticleTypeHelpers, Name: "helperLabel3", Content: "helperContent3"},

		{AdminId: adminsModel.SuperAdminId, Type: systemsModel.ArticleTypeAbout, Name: "aboutUs", Content: "aboutUsContent"},
		{AdminId: adminsModel.SuperAdminId, Type: systemsModel.ArticleTypeAbout, Name: "serviceAgreement", Content: "serviceAgreementContent"},
		{AdminId: adminsModel.SuperAdminId, Type: systemsModel.ArticleTypeAbout, Name: "privacyPolicy", Content: "privacyPolicyContent"},
		{AdminId: adminsModel.SuperAdminId, Type: systemsModel.ArticleTypeAbout, Name: "disclaimer", Content: "disclaimerContent"},

		{AdminId: adminsModel.SuperAdminId, Type: systemsModel.ArticleTypeProduct, Name: "productLabel1", Content: "productContent1"},
		{AdminId: adminsModel.SuperAdminId, Type: systemsModel.ArticleTypeProduct, Name: "productLabel2", Content: "productContent2"},
		{AdminId: adminsModel.SuperAdminId, Type: systemsModel.ArticleTypeProduct, Name: "productLabel3", Content: "productContent3"},

		{AdminId: adminsModel.SuperAdminId, Type: systemsModel.ArticleTypeService, Image: "/assets/icon/address.png", Name: "607-699 14th Pl NE, Washington, DC 20002 USA", Link: "https://maps.app.goo.gl/U2gNxExtu16Tbasb8"},
		{AdminId: adminsModel.SuperAdminId, Type: systemsModel.ArticleTypeService, Image: "/assets/icon/telephone.png", Name: "+12023883500", Link: "tel:12023883500"},
		{AdminId: adminsModel.SuperAdminId, Type: systemsModel.ArticleTypeService, Image: "/assets/icon/email.png", Name: "muiprosperyls15@gmail.com", Link: "mailto:muiprosperyls15@gmail.com"},

		{AdminId: adminsModel.SuperAdminId, Type: systemsModel.ArticleTypeSocial, Image: "/assets/icon/telegram.png", Name: "Telegram", Link: "https://t.me/bajie119"},
		{AdminId: adminsModel.SuperAdminId, Type: systemsModel.ArticleTypeSocial, Image: "/assets/icon/facebook.png", Name: "Facebook", Link: "https://t.me/bajie119"},
		{AdminId: adminsModel.SuperAdminId, Type: systemsModel.ArticleTypeSocial, Image: "/assets/icon/whatsapp.png", Name: "Whatsapp", Link: "https://t.me/bajie119"},
		{AdminId: adminsModel.SuperAdminId, Type: systemsModel.ArticleTypeSocial, Image: "/assets/icon/line.png", Name: "Line", Link: "https://t.me/bajie119"},
		{AdminId: adminsModel.SuperAdminId, Type: systemsModel.ArticleTypeSocial, Image: "/assets/icon/twitter.png", Name: "Twitter", Link: "https://t.me/bajie119"},

		// 内容中心
		{AdminId: adminsModel.SuperAdminId, Type: systemsModel.ArticleTypeDefault, Name: "article1", Content: "articleContent1"},
		{AdminId: adminsModel.SuperAdminId, Type: systemsModel.ArticleTypeDefault, Name: "article2", Content: "articleContent2"},
		{AdminId: adminsModel.SuperAdminId, Type: systemsModel.ArticleTypeDefault, Name: "article3", Content: "articleContent3"},
	}
}
