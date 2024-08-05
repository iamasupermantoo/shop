package datas

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/types"
)

// InitAdminMenu 初始化菜单
func InitAdminMenu() []*adminsModel.AdminMenu {
	return []*adminsModel.AdminMenu{
		//	用户管理
		{GormModel: types.GormModel{ID: 100}, Status: 10, ParentId: 0, Name: "用户管理", Route: "", Sort: 1, Data: &adminsModel.AdminMenuData{Icon: "manage_accounts"}},
		{GormModel: types.GormModel{ID: 101}, Status: 10, ParentId: 100, Name: "用户列表", Route: "/users/user/index", Sort: 1, Data: &adminsModel.AdminMenuData{Icon: "people", Tmp: "/views", ConfURL: "/auth/users/user/views"}},
		{GormModel: types.GormModel{ID: 102}, Status: 10, ParentId: 100, Name: "会员列表", Route: "/users/level/index", Sort: 2, Data: &adminsModel.AdminMenuData{Icon: "diamond", Tmp: "/views", ConfURL: "/auth/users/level/views"}},
		{GormModel: types.GormModel{ID: 103}, Status: 10, ParentId: 100, Name: "认证管理", Route: "/users/auth/index", Sort: 3, Data: &adminsModel.AdminMenuData{Icon: "badge", Tmp: "/views", ConfURL: "/auth/users/auth/views"}},
		{GormModel: types.GormModel{ID: 105}, Status: 10, ParentId: 100, Name: "访问记录", Route: "/users/access/index", Sort: 5, Data: &adminsModel.AdminMenuData{Icon: "dvr", Tmp: "/views", ConfURL: "/auth/users/access/views"}},
		{GormModel: types.GormModel{ID: 106}, Status: 10, ParentId: 100, Name: "用户邀请", Route: "/users/invite/index", Sort: 6, Data: &adminsModel.AdminMenuData{Icon: "share", Tmp: "/views", ConfURL: "/auth/users/invite/views"}},
		{GormModel: types.GormModel{ID: 107}, Status: 10, ParentId: 100, Name: "渠道管理", Route: "/users/channel/index", Sort: 7, Data: &adminsModel.AdminMenuData{Icon: "person_pin", Tmp: "/views", ConfURL: "/auth/users/channel/views"}},

		//	产品管理
		{GormModel: types.GormModel{ID: 110}, Status: 10, ParentId: 0, Name: "商品管理", Route: "", Sort: 2, Data: &adminsModel.AdminMenuData{Icon: "sym_o_storefront"}},
		{GormModel: types.GormModel{ID: 112}, Status: 10, ParentId: 110, Name: "商品分类", Route: "/products/category/index", Sort: 1, Data: &adminsModel.AdminMenuData{Icon: "sym_o_category", Tmp: "/views", ConfURL: "/auth/products/category/views"}},
		{GormModel: types.GormModel{ID: 113}, Status: 10, ParentId: 110, Name: "商品列表", Route: "/products/product/index", Sort: 2, Data: &adminsModel.AdminMenuData{Icon: "sym_o_store", Tmp: "/views", ConfURL: "/auth/products/product/views"}},

		//	店铺管理
		{GormModel: types.GormModel{ID: 160}, Status: 10, ParentId: 0, Name: "店铺管理", Route: "", Sort: 3, Data: &adminsModel.AdminMenuData{Icon: "sym_o_local_mall"}},
		{GormModel: types.GormModel{ID: 163}, Status: 10, ParentId: 160, Name: "入驻申请", Route: "/shops/settled/index", Sort: 1, Data: &adminsModel.AdminMenuData{Icon: "sym_o_add_business", Tmp: "/views", ConfURL: "/auth/shops/settled/views"}},
		{GormModel: types.GormModel{ID: 161}, Status: 10, ParentId: 160, Name: "店铺列表", Route: "/shops/store/index", Sort: 2, Data: &adminsModel.AdminMenuData{Icon: "sym_o_store", Tmp: "/views", ConfURL: "/auth/shops/store/views"}},
		{GormModel: types.GormModel{ID: 165}, Status: 10, ParentId: 160, Name: "购物车", Route: "/shops/cart/index", Sort: 3, Data: &adminsModel.AdminMenuData{Icon: "sym_o_shopping_cart", Tmp: "/views", ConfURL: "/auth/shops/cart/views"}},
		{GormModel: types.GormModel{ID: 162}, Status: 10, ParentId: 160, Name: "店铺订单", Route: "/shops/order/index", Sort: 4, Data: &adminsModel.AdminMenuData{Icon: "sym_o_local_shipping", Tmp: "/views", ConfURL: "/auth/shops/order/views"}},
		{GormModel: types.GormModel{ID: 111}, Status: 10, ParentId: 160, Name: "商品订单", Route: "/products/order/index", Sort: 5, Data: &adminsModel.AdminMenuData{Icon: "currency_bitcoin", Tmp: "/views", ConfURL: "/auth/products/order/views"}},
		{GormModel: types.GormModel{ID: 166}, Status: 10, ParentId: 160, Name: "收货地址", Route: "/shops/address/index", Sort: 6, Data: &adminsModel.AdminMenuData{Icon: "sym_o_person_pin_circle", Tmp: "/views", ConfURL: "/auth/shops/address/views"}},
		{GormModel: types.GormModel{ID: 167}, Status: 10, ParentId: 160, Name: "关注收藏", Route: "/shops/follow/index", Sort: 7, Data: &adminsModel.AdminMenuData{Icon: "sym_o_follow_the_signs", Tmp: "/views", ConfURL: "/auth/shops/follow/views"}},
		{GormModel: types.GormModel{ID: 168}, Status: 10, ParentId: 160, Name: "商品评论", Route: "/shops/comment/index", Sort: 8, Data: &adminsModel.AdminMenuData{Icon: "sym_o_speaker_notes", Tmp: "/views", ConfURL: "/auth/shops/comment/views"}},
		{GormModel: types.GormModel{ID: 164}, Status: 10, ParentId: 160, Name: "售后管理", Route: "/shops/refund/index", Sort: 9, Data: &adminsModel.AdminMenuData{Icon: "sym_o_shopping_bag", Tmp: "/views", ConfURL: "/auth/shops/refund/views"}},

		// 余额管理
		{GormModel: types.GormModel{ID: 120}, Status: 10, ParentId: 0, Name: "余额管理", Route: "", Sort: 2, Data: &adminsModel.AdminMenuData{Icon: "wallet"}},
		{GormModel: types.GormModel{ID: 121}, Status: 10, ParentId: 120, Name: "充值订单", Route: "/wallets/order/deposit/balance", Sort: 1, Data: &adminsModel.AdminMenuData{Icon: "assured_workload", Tmp: "/views", ConfURL: "/auth/wallets/order/views?type=1"}},
		{GormModel: types.GormModel{ID: 122}, Status: 10, ParentId: 120, Name: "提现订单", Route: "/wallets/order/withdraw/balance", Sort: 2, Data: &adminsModel.AdminMenuData{Icon: "credit_score", Tmp: "/views", ConfURL: "/auth/wallets/order/views?type=11"}},

		// 资产管理
		{GormModel: types.GormModel{ID: 130}, Status: 10, ParentId: 0, Name: "资产管理", Route: "", Sort: 3, Data: &adminsModel.AdminMenuData{Icon: "card_giftcard"}},
		{GormModel: types.GormModel{ID: 131}, Status: 10, ParentId: 130, Name: "充值订单", Route: "/wallets/order/deposit/assets", Sort: 1, Data: &adminsModel.AdminMenuData{Icon: "assured_workload", Tmp: "/views", ConfURL: "/auth/wallets/order/views?type=2"}},
		{GormModel: types.GormModel{ID: 132}, Status: 10, ParentId: 130, Name: "提现订单", Route: "/wallets/order/withdraw/assets", Sort: 2, Data: &adminsModel.AdminMenuData{Icon: "credit_score", Tmp: "/views", ConfURL: "/auth/wallets/order/views?type=12"}},
		{GormModel: types.GormModel{ID: 133}, Status: 10, ParentId: 130, Name: "用户资产", Route: "/users/assets/index", Sort: 3, Data: &adminsModel.AdminMenuData{Icon: "account_balance_wallet", Tmp: "/views", ConfURL: "/auth/users/assets/views"}},
		{GormModel: types.GormModel{ID: 134}, Status: 10, ParentId: 130, Name: "资产列表", Route: "/wallets/assets/index", Sort: 4, Data: &adminsModel.AdminMenuData{Icon: "payments", Tmp: "/views", ConfURL: "/auth/wallets/assets/views"}},

		// 钱包管理
		{GormModel: types.GormModel{ID: 140}, Status: 10, ParentId: 0, Name: "财务管理", Route: "", Sort: 4, Data: &adminsModel.AdminMenuData{Icon: "assured_workload"}},
		{GormModel: types.GormModel{ID: 141}, Status: 10, ParentId: 140, Name: "转移记录", Route: "/wallets/transfer/index", Sort: 1, Data: &adminsModel.AdminMenuData{Icon: "swap_horiz", Tmp: "/views", ConfURL: "/auth/wallets/transfer/views"}},
		{GormModel: types.GormModel{ID: 142}, Status: 10, ParentId: 140, Name: "转换记录", Route: "/wallets/convert/index", Sort: 2, Data: &adminsModel.AdminMenuData{Icon: "currency_exchange", Tmp: "/views", ConfURL: "/auth/wallets/convert/views"}},
		{GormModel: types.GormModel{ID: 143}, Status: 10, ParentId: 140, Name: "账单明细", Route: "/wallets/bill/index", Sort: 3, Data: &adminsModel.AdminMenuData{Icon: "assignment", Tmp: "/views", ConfURL: "/auth/wallets/bill/views"}},
		{GormModel: types.GormModel{ID: 144}, Status: 10, ParentId: 140, Name: "提现账户", Route: "/users/account/index", Sort: 4, Data: &adminsModel.AdminMenuData{Icon: "contact_mail", Tmp: "/views", ConfURL: "/auth/users/account/views"}},
		{GormModel: types.GormModel{ID: 145}, Status: 10, ParentId: 140, Name: "支付管理", Route: "/wallets/payment/index", Sort: 5, Data: &adminsModel.AdminMenuData{Icon: "sell", Tmp: "/views", ConfURL: "/auth/wallets/payment/views"}},

		//	后台管理
		{GormModel: types.GormModel{ID: 1}, Status: 10, ParentId: 0, Name: "后台管理", Route: "", Sort: 90, Data: &adminsModel.AdminMenuData{Icon: "manage_accounts"}},
		{GormModel: types.GormModel{ID: 2}, Status: 10, ParentId: 1, Name: "管理列表", Route: "/admins/manage/index", Sort: 1, Data: &adminsModel.AdminMenuData{Icon: "manage_accounts", Tmp: "/views", ConfURL: "/auth/admins/manage/views"}},
		{GormModel: types.GormModel{ID: 3}, Status: 10, ParentId: 1, Name: "参数配置", Route: "/admins/setting/index", Sort: 2, Data: &adminsModel.AdminMenuData{Icon: "settings_ethernet", Tmp: "/setting", ConfURL: "/auth/admins/setting/views"}},
		{GormModel: types.GormModel{ID: 4}, Status: 10, ParentId: 1, Name: "角色配置", Route: "/admins/role/index", Sort: 3, Data: &adminsModel.AdminMenuData{Icon: "6_ft_apart", Tmp: "/views", ConfURL: "/auth/admins/role/views"}},
		{GormModel: types.GormModel{ID: 5}, Status: 10, ParentId: 1, Name: "菜单管理", Route: "/admins/menu/index", Sort: 4, Data: &adminsModel.AdminMenuData{Icon: "menu_open", Tmp: "/views", ConfURL: "/auth/admins/menu/views"}},
		{GormModel: types.GormModel{ID: 6}, Status: 10, ParentId: 1, Name: "操作日志", Route: "/admins/logs/index", Sort: 5, Data: &adminsModel.AdminMenuData{Icon: "event_note", Tmp: "/views", ConfURL: "/auth/admins/logs/views"}},

		//	系统配置
		{GormModel: types.GormModel{ID: 40}, Status: 10, ParentId: 0, Name: "系统配置", Route: "", Sort: 92, Data: &adminsModel.AdminMenuData{Icon: "widgets"}},
		{GormModel: types.GormModel{ID: 41}, Status: 10, ParentId: 40, Name: "等级配置", Route: "/systems/level/index", Sort: 1, Data: &adminsModel.AdminMenuData{Icon: "diamond", Tmp: "/views", ConfURL: "/auth/systems/level/views"}},
		{GormModel: types.GormModel{ID: 42}, Status: 10, ParentId: 40, Name: "国家配置", Route: "/systems/country/index", Sort: 2, Data: &adminsModel.AdminMenuData{Icon: "flag", Tmp: "/views", ConfURL: "/auth/systems/country/views"}},
		{GormModel: types.GormModel{ID: 43}, Status: 10, ParentId: 40, Name: "语言配置", Route: "/systems/lang/index", Sort: 3, Data: &adminsModel.AdminMenuData{Icon: "language", Tmp: "/views", ConfURL: "/auth/systems/lang/views"}},
		{GormModel: types.GormModel{ID: 44}, Status: 10, ParentId: 40, Name: "翻译配置", Route: "/systems/translate/index", Sort: 4, Data: &adminsModel.AdminMenuData{Icon: "translate", Tmp: "/views", ConfURL: "/auth/systems/translate/views"}},
		{GormModel: types.GormModel{ID: 45}, Status: 10, ParentId: 40, Name: "通知配置", Route: "/systems/notify/index", Sort: 5, Data: &adminsModel.AdminMenuData{Icon: "notifications", Tmp: "/views", ConfURL: "/auth/systems/notify/views"}},
		{GormModel: types.GormModel{ID: 46}, Status: 10, ParentId: 40, Name: "文章配置", Route: "/systems/article/index", Sort: 6, Data: &adminsModel.AdminMenuData{Icon: "article", Tmp: "/views", ConfURL: "/auth/systems/article/views"}},
		{GormModel: types.GormModel{ID: 47}, Status: 10, ParentId: 40, Name: "前台菜单", Route: "/systems/menu/index", Sort: 6, Data: &adminsModel.AdminMenuData{Icon: "menu_book", Tmp: "/views", ConfURL: "/auth/systems/menu/views"}},
	}
}
