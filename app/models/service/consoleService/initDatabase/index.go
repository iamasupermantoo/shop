package initDatabase

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/chatsModel"
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/consoleService"
	"gofiber/app/models/service/consoleService/initDatabase/datas"
	"gofiber/app/module/database"
	"gofiber/utils"
	"reflect"
)

// Database 数据库操作
type Database struct {
	tableList   []Table
	FilterTable []string
}

// GetTable 获取表
func (_Database *Database) GetTable(name string) Table {
	for _, table := range _Database.tableList {
		currentTableName := database.Db.NamingStrategy.TableName(reflect.TypeOf(table.Model).Name())
		if currentTableName == name {
			table.Name = currentTableName
			return table
		}
	}
	return Table{}
}

// AddTable 添加表
func (_Database *Database) AddTable(table ...Table) *Database {
	_Database.tableList = append(_Database.tableList, table...)
	return _Database
}

// ExecInit 执行初始化
func (_Database *Database) ExecInit() {
	for _, table := range _Database.tableList {
		// 过滤不需要初始化的表
		if utils.ArrayStringIndexOf(_Database.FilterTable, table.Name) > -1 {
			continue
		}

		// 先执行删除
		_ = database.Db.Migrator().DropTable(table.Model)

		// 创建表
		_ = database.Db.Set("gorm:table_options", "COMMENT='"+table.Comment+"'").AutoMigrate(table.Model)

		// 添加数据
		for _, datum := range table.Data {
			database.Db.Create(datum)
		}
	}

	// 初始化产品信息
	datas.InitProduct("./app/models/service/consoleService/initDatabase/datas/sql")

	// 同步商户数据
	merchantList := make([]*adminsModel.AdminUser, 0)
	database.Db.Model(&adminsModel.AdminUser{}).Where("parent_id = ?", adminsModel.SuperAdminId).Find(&merchantList)

	// 重置商户数据
	for _, merchant := range merchantList {
		_ = consoleService.NewMerchant(merchant.ID, append(_Database.FilterTable, "product", "category")).RunRest()
	}
}

// InitTables 初始化所有表
func (_Database *Database) InitTables() *Database {
	// 初始化数据库表
	authItem, authChild := datas.InitAdminAuth()
	_Database.AddTable(
		// 权限分配
		Table{Model: &adminsModel.AuthAssignment{}, Data: []interface{}{[]*adminsModel.AuthAssignment{
			{Name: adminsModel.AuthRoleSuperManage, AdminId: adminsModel.SuperAdminId},
			{Name: adminsModel.AuthRoleMerchantManage, AdminId: adminsModel.MerchantAdminId},
			{Name: adminsModel.AuthRoleAgentManage, AdminId: adminsModel.AgentAdminId},
		}}, Comment: "权限分配表", Name: "auth_assignment"},

		// 权限目录表
		Table{Model: adminsModel.AuthItem{}, Data: []interface{}{authItem}, Comment: "权限目录表", Name: "auth_item"},

		// 权限目录对应表
		Table{Model: adminsModel.AuthChild{}, Data: []interface{}{authChild}, Comment: "权限目录对应表", Name: "auth_child"},

		// 管理用户
		Table{Model: adminsModel.AdminUser{}, Data: []interface{}{datas.InitAdminUser()}, Comment: "管理用户表", Name: "admin_user"},

		// 管理日志表
		Table{Model: adminsModel.AdminLogs{}, Data: nil, Comment: "管理日志表", Name: "admin_logs"},

		//管理菜单表
		Table{Model: adminsModel.AdminMenu{}, Data: []interface{}{datas.InitAdminMenu()}, Comment: "管理菜单表", Name: "admin_menu"},

		// 管理设置表
		Table{Model: adminsModel.AdminSetting{}, Data: []interface{}{datas.InitAdminSetting()}, Comment: "管理设置表", Name: "admin_setting"},

		// 文章管理表
		Table{Model: systemsModel.Article{}, Data: []interface{}{datas.InitArticle()}, Comment: "文章内容表", Name: "article"},

		// 消息通知表
		Table{Model: systemsModel.Notify{}, Data: nil, Comment: "消息通知表", Name: "notify"},

		// 国家表
		Table{Model: systemsModel.Country{}, Data: []interface{}{datas.InitCountry()}, Comment: "系统国家表", Name: "country"},

		// 语言表
		Table{Model: systemsModel.Lang{}, Data: []interface{}{datas.InitLang()}, Comment: "系统语言表", Name: "lang"},

		// 用户等级表
		Table{Model: systemsModel.Level{}, Data: []interface{}{datas.InitLevel()}, Comment: "系统等级表", Name: "level"},

		// 系统翻译表
		Table{Model: systemsModel.Translate{}, Data: []interface{}{datas.InitTranslate()}, Comment: "系统翻译表", Name: "translate"},

		// 前台菜单表
		Table{Model: systemsModel.Menu{}, Data: []interface{}{datas.InitMenu()}, Comment: "前台菜单表", Name: "menu"},

		// 渠道表
		Table{Model: usersModel.Channel{}, Data: nil, Comment: "渠道表", Name: "channel"},

		// 用户表
		Table{Model: usersModel.User{}, Data: []interface{}{[]*usersModel.User{
			{AdminId: adminsModel.MerchantAdminId, UserName: "ceshi", Password: utils.PasswordEncrypt("abc123"), SecurityKey: utils.PasswordEncrypt("abc123"), Money: 100000},
		}}, Comment: "用户表", Name: "user"},

		// 用户设置表
		Table{Model: usersModel.Setting{}, Data: []interface{}{datas.InitUserSetting()}, Comment: "用户设置表", Name: "setting"},

		// 用户访问表
		Table{Model: usersModel.Access{}, Data: nil, Comment: "用户访问表", Name: "access"},

		// 用户邀请表
		Table{Model: usersModel.Invite{}, Data: nil, Comment: "用户邀请表", Name: "invite"},

		// 用户实名
		Table{Model: usersModel.UserAuth{}, Data: nil, Comment: "用户实名表", Name: "user_auth"},

		// 用户等级表
		Table{Model: usersModel.UserLevel{}, Data: nil, Comment: "用户等级表", Name: "user_level"},

		// 钱包资产表
		Table{Model: walletsModel.WalletAssets{}, Data: []interface{}{datas.InitWalletAssets()}, Comment: "钱包资产表", Name: "wallet_assets"},

		// 钱包支付表
		Table{Model: walletsModel.WalletPayment{}, Data: []interface{}{datas.InitWalletPayment()}, Comment: "钱包支付表", Name: "wallet_payment"},

		// 钱包用户账户表
		Table{Model: walletsModel.WalletUserAccount{}, Data: nil, Comment: "钱包用户账户表", Name: "wallet_user_account"},

		// 钱包用户资产表
		Table{Model: walletsModel.WalletUserAssets{}, Data: nil, Comment: "钱包用户资产表", Name: "wallet_user_assets"},

		// 钱包用户账单表
		Table{Model: walletsModel.WalletUserBill{}, Data: nil, Comment: "钱包用户账单表", Name: "wallet_user_bill"},

		// 钱包用户转换表
		Table{Model: walletsModel.WalletUserConvert{}, Data: nil, Comment: "钱包用户转换表", Name: "wallet_user_convert"},

		// 钱包用户订单表
		Table{Model: walletsModel.WalletUserOrder{}, Data: nil, Comment: "钱包用户订单表", Name: "wallet_user_order"},

		// 钱包用户转移表
		Table{Model: walletsModel.WalletUserTransfer{}, Data: nil, Comment: "钱包用户转移表", Name: "wallet_user_transfer"},

		// 产品分类表
		Table{Model: productsModel.Category{}, Data: []interface{}{datas.InitProductCategory()}, Comment: "产品分类表", Name: "category"},

		// 产品商品表
		Table{Model: productsModel.Product{}, Data: nil, Comment: "产品商品表", Name: "product"},

		// 产品订单表
		Table{Model: productsModel.ProductOrder{}, Data: nil, Comment: "产品订单表", Name: "product_order"},

		// 产品浏览记录
		Table{Model: productsModel.ProductBrowsing{}, Data: nil, Comment: "商品浏览记录", Name: "product_browsing"},

		// 产品属性名称表
		Table{Model: productsModel.ProductAttrsKey{}, Data: nil, Comment: "产品属性名称表", Name: "product_attrs_key"},

		// 产品属性键值表
		Table{Model: productsModel.ProductAttrsVal{}, Data: nil, Comment: "产品属性键值表", Name: "product_attrs_val"},

		// 产品属性SKU表
		Table{Model: productsModel.ProductAttrsSku{}, Data: nil, Comment: "产品属性SKU表", Name: "product_attrs_sku"},

		// 产品店铺订单表
		Table{Model: shopsModel.ProductStoreOrder{}, Data: nil, Comment: "产品店铺订单表", Name: "product_store_order"},

		// 用户购物地址表
		Table{Model: shopsModel.ShippingAddress{}, Data: []interface{}{datas.InitShippingAddress()}, Comment: "用户购物地址表", Name: "shipping_address"},

		// 用户店铺表
		Table{Model: shopsModel.Store{}, Data: nil, Comment: "用户店铺表", Name: "store"},

		// 店铺购物车
		Table{Model: shopsModel.StoreCart{}, Data: nil, Comment: "店铺购物车", Name: "store_cart"},

		// 店铺商品评论
		Table{Model: shopsModel.StoreComment{}, Data: nil, Comment: "店铺商品评论", Name: "store_comment"},

		// 店铺商品关注
		Table{Model: shopsModel.StoreFollow{}, Data: nil, Comment: "店铺商品关注", Name: "store_follow"},

		// 店铺产品售后
		Table{Model: shopsModel.StoreRefund{}, Data: nil, Comment: "店铺产品售后", Name: "store_refund"},

		// 店铺入驻申请
		Table{Model: shopsModel.StoreSettled{}, Data: nil, Comment: "店铺入驻申请", Name: "store_settled"},

		// 会话聊天
		Table{Model: chatsModel.Conversation{}, Data: nil, Comment: "聊天会话", Name: "conversation"},

		// 会话消息
		Table{Model: chatsModel.Messages{}, Data: nil, Comment: "会话消息", Name: "messages"},
	)

	return _Database
}
