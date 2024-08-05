package datas

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/models/model/types"
)

// InitMenu 初始化前台菜单
func InitMenu() []*systemsModel.Menu {
	return []*systemsModel.Menu{
		// 导航菜单
		{GormModel: types.GormModel{ID: 1}, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeNavigation, Name: "tab_home", Route: "/", Sort: 1, Icon: "/assets/menu/tab_home.png", ActiveIcon: "/assets/menu/tab_home_active.png", IsDesktop: types.ModelBoolFalse, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 10}, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeNavigation, Name: "tab_category", Route: "/category", Sort: 2, Icon: "/assets/menu/tab_category.png", ActiveIcon: "/assets/menu/tab_category_active.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 20}, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeNavigation, Name: "tab_shopping", Route: "/product/shopping", Sort: 3, Icon: "/assets/menu/tab_shopping.png", ActiveIcon: "/assets/menu/tab_shopping_active.png", IsDesktop: types.ModelBoolFalse, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 30}, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeNavigation, Name: "tab_message", Route: "/message", Sort: 4, Icon: "/assets/menu/tab_message.png", ActiveIcon: "/assets/menu/tab_message_active.png", IsDesktop: types.ModelBoolFalse, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 40}, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeNavigation, Name: "newProduct", Route: "/product", Sort: 5, Icon: "", ActiveIcon: "", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolFalse},
		{GormModel: types.GormModel{ID: 41}, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeNavigation, Name: "discountProduct", Route: "/product", Sort: 6, Icon: "", ActiveIcon: "", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolFalse},
		{GormModel: types.GormModel{ID: 50}, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeNavigation, Name: "tab_user", Route: "/users", Sort: 7, Icon: "/assets/menu/tab_user.png", ActiveIcon: "/assets/menu/tab_user_active.png", IsDesktop: types.ModelBoolFalse, IsMobile: types.ModelBoolTrue},

		// 店铺管理菜单
		{GormModel: types.GormModel{ID: 90}, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeSetting, Name: "shopsManage", Route: "", Sort: 1, Icon: "", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 91}, ParentId: 90, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeSetting, Name: "merchantsSettleIn", Route: "/stores/settled", Sort: 1, Icon: "/assets/menu/settled.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 92}, ParentId: 90, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeSetting, Name: "storeService", Route: "/stores/home", Sort: 2, Icon: "/assets/menu/store.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 93}, ParentId: 90, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeSetting, Name: "shopsOrder", Route: "/stores/user/order", Sort: 3, Icon: "/assets/menu/store_order.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 94}, ParentId: 90, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeSetting, Name: "followProduct", Route: "/stores/follow?type=2", Sort: 4, Icon: "/assets/menu/follow_product.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 95}, ParentId: 90, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeSetting, Name: "followStore", Route: "/stores/follow?type=1", Sort: 5, Icon: "/assets/menu/follow_store.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 96}, ParentId: 90, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeSetting, Name: "browsingProduct", Route: "/stores/product/browsing", Sort: 6, Icon: "/assets/menu/views_product.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 97}, ParentId: 90, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeSetting, Name: "shippingAddress", Route: "/stores/address/index", Sort: 6, Icon: "/assets/menu/address.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},

		// 功能菜单 - 钱包管理
		{GormModel: types.GormModel{ID: 100}, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeSetting, Name: "walletsManage", Route: "", Sort: 1, Icon: "", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 101}, ParentId: 100, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeSetting, Name: "userWallets", Route: "/wallets", Sort: 1, Icon: "/assets/menu/wallets.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 102}, ParentId: 100, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeSetting, Name: "userAssets", Route: "/wallets/assets", Sort: 2, Icon: "/assets/menu/assets.png", IsDesktop: types.ModelBoolFalse, IsMobile: types.ModelBoolFalse},
		{GormModel: types.GormModel{ID: 103}, ParentId: 100, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeSetting, Name: "withdrawAccount", Route: "/wallets/account", Sort: 3, Icon: "/assets/menu/account.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 104}, ParentId: 100, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeSetting, Name: "walletsBill", Route: "/wallets/bill", Sort: 4, Icon: "/assets/menu/bill.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 105}, ParentId: 100, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeSetting, Name: "transfer", Route: "/wallets/transfer", Sort: 5, Icon: "/assets/menu/transfer.png", IsDesktop: types.ModelBoolFalse, IsMobile: types.ModelBoolFalse},
		{GormModel: types.GormModel{ID: 106}, ParentId: 100, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeSetting, Name: "convert", Route: "/wallets/convert", Sort: 6, Icon: "/assets/menu/convert.png", IsDesktop: types.ModelBoolFalse, IsMobile: types.ModelBoolFalse},

		// 功能菜单 - 团队管理
		{GormModel: types.GormModel{ID: 120}, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeSetting, Name: "teamManage", Route: "", Sort: 2, Icon: "", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 121}, ParentId: 120, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeSetting, Name: "userTeam", Route: "/users/team", Sort: 1, Icon: "/assets/menu/team.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 122}, ParentId: 120, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeSetting, Name: "userEarnings", Route: "/users/team/earnings", Sort: 2, Icon: "/assets/menu/earnings.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 123}, ParentId: 120, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeSetting, Name: "userInvite", Route: "/users/team/invite", Sort: 3, Icon: "/assets/menu/share.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},

		// 功能菜单 - 更多服务
		{GormModel: types.GormModel{ID: 130}, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeSetting, Name: "serviceManage", Route: "", Sort: 3, Icon: "", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 131}, ParentId: 130, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeSetting, Name: "userLevel", Route: "/users/level", Sort: 1, Icon: "/assets/menu/level.png", IsDesktop: types.ModelBoolFalse, IsMobile: types.ModelBoolFalse},
		{GormModel: types.GormModel{ID: 132}, ParentId: 130, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeSetting, Name: "userAuth", Route: "/users/auth", Sort: 2, Icon: "/assets/menu/auth.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 133}, ParentId: 130, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeSetting, Name: "helpers", Route: "/helpers", Sort: 3, Icon: "/assets/menu/helpers.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 134}, ParentId: 130, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeSetting, Name: "download", Route: "/download", Sort: 4, Icon: "/assets/menu/download.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 135}, ParentId: 130, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeSetting, Name: "settings", Route: "/settings", Sort: 5, Icon: "/assets/menu/settings.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},

		// 更多服务菜单
		{GormModel: types.GormModel{ID: 200}, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeMore, Name: "deposit", Route: "/wallets/deposit?mode=1", Sort: 1, Icon: "/assets/menu/deposit.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 201}, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeMore, Name: "assetsDeposit", Route: "/wallets/assets", Sort: 2, Icon: "/assets/menu/deposit.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolFalse},
		{GormModel: types.GormModel{ID: 202}, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeMore, Name: "withdraw", Route: "/wallets/withdraw?mode=1", Sort: 3, Icon: "/assets/menu/withdraw.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 203}, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeMore, Name: "assetsWithdraw", Route: "/wallets/assets", Sort: 4, Icon: "/assets/menu/withdraw.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolFalse},
		{GormModel: types.GormModel{ID: 204}, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeMore, Name: "transfer", Route: "/wallets/transfer", Sort: 5, Icon: "/assets/menu/transfer.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 205}, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeMore, Name: "convert", Route: "/wallets/convert", Sort: 6, Icon: "/assets/menu/convert.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 206}, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeMore, Name: "switchLang", Route: "/lang", Sort: 7, Icon: "/assets/menu/lang.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 207}, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeMore, Name: "helpers", Route: "/helpers", Sort: 8, Icon: "/assets/menu/helpers.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 208}, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeMore, Name: "download", Route: "/download", Sort: 9, Icon: "/assets/menu/download.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 209}, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeMore, Name: "userInvite", Route: "/users/team/invite", Sort: 10, Icon: "/assets/menu/share.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},

		{GormModel: types.GormModel{ID: 300}, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeStore, Name: "merchantManage", Route: "", Sort: 5, Icon: "", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 301}, ParentId: 300, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeStore, Name: "mallServices", Route: "/", Sort: 1, Icon: "", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 302}, ParentId: 300, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeStore, Name: "userWallets", Route: "/wallets", Sort: 2, Icon: "/assets/menu/wallets.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 303}, ParentId: 300, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeStore, Name: "withdrawAccount", Route: "/wallets/account", Sort: 3, Icon: "/assets/menu/account.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 304}, ParentId: 300, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeStore, Name: "walletsBill", Route: "/wallets/bill", Sort: 4, Icon: "/assets/menu/bill.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 305}, ParentId: 300, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeStore, Name: "userLevel", Route: "/users/level", Sort: 5, Icon: "/assets/menu/level.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 307}, ParentId: 300, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeStore, Name: "helpers", Route: "/helpers", Sort: 6, Icon: "/assets/menu/helpers.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 308}, ParentId: 300, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeStore, Name: "download", Route: "/download", Sort: 7, Icon: "/assets/menu/download.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
		{GormModel: types.GormModel{ID: 309}, ParentId: 300, AdminId: adminsModel.SuperAdminId, Type: systemsModel.MenuTypeStore, Name: "settings", Route: "/settings", Sort: 8, Icon: "/assets/menu/settings.png", IsDesktop: types.ModelBoolTrue, IsMobile: types.ModelBoolTrue},
	}
}
