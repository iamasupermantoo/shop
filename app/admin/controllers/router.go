package controllers

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"gofiber/app/admin/controllers/admins/logs"
	"gofiber/app/admin/controllers/admins/manage"
	"gofiber/app/admin/controllers/admins/menu"
	"gofiber/app/admin/controllers/admins/role"
	"gofiber/app/admin/controllers/admins/setting"
	"gofiber/app/admin/controllers/index"
	"gofiber/app/admin/controllers/products/category"
	"gofiber/app/admin/controllers/products/order"
	"gofiber/app/admin/controllers/products/product"
	"gofiber/app/admin/controllers/products/productBrowsing"
	"gofiber/app/admin/controllers/shops/address"
	"gofiber/app/admin/controllers/shops/cart"
	"gofiber/app/admin/controllers/shops/comment"
	"gofiber/app/admin/controllers/shops/follow"
	shopsOrder "gofiber/app/admin/controllers/shops/order"
	"gofiber/app/admin/controllers/shops/refund"
	"gofiber/app/admin/controllers/shops/settled"
	"gofiber/app/admin/controllers/shops/store"
	"gofiber/app/admin/controllers/systems/article"
	"gofiber/app/admin/controllers/systems/country"
	"gofiber/app/admin/controllers/systems/lang"
	"gofiber/app/admin/controllers/systems/level"
	homeMenu "gofiber/app/admin/controllers/systems/menu"
	"gofiber/app/admin/controllers/systems/notify"
	"gofiber/app/admin/controllers/systems/translate"
	"gofiber/app/admin/controllers/users/access"
	"gofiber/app/admin/controllers/users/channel"
	"gofiber/app/admin/controllers/users/invite"
	usersSetting "gofiber/app/admin/controllers/users/setting"
	"gofiber/app/admin/controllers/users/user"
	"gofiber/app/admin/controllers/users/userAuth"
	"gofiber/app/admin/controllers/users/userLevel"
	"gofiber/app/admin/controllers/wallets/assets"
	"gofiber/app/admin/controllers/wallets/payment"
	"gofiber/app/admin/controllers/wallets/userAccount"
	"gofiber/app/admin/controllers/wallets/userAssets"
	"gofiber/app/admin/controllers/wallets/userBill"
	"gofiber/app/admin/controllers/wallets/userConvert"
	"gofiber/app/admin/controllers/wallets/userOrder"
	"gofiber/app/admin/controllers/wallets/userTransfer"
	"gofiber/app/middleware"
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/models/service/commonService"
	"gofiber/app/module/cache"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/app/websocket"
	"gofiber/module/socket/handler"
	"gofiber/utils"
	"strings"
	"time"
)

// adminRouterList 后端路由列表
var adminRouterList = map[string]string{}

// 过滤掉写日志掉路由
var adminFilterRouterList = []string{"audio", "index", "views", "init", "info", "balance", "assets", "upload"}

// InitAdminRouter 初始化后端路由
func InitAdminRouter(app *fiber.App) {
	authRouter := middlewareAuthRouter(app)

	router := app.Group("/" + adminsModel.ServiceAdminRouteName)
	router.Get("/ws", handler.NewSocketConn(websocket.AdminWebSocket)).Name("Websocket")
	// 初始化websocket一些数据

	router.Get("/captcha/create", index.NewCaptcha).Name("创建验证码")
	router.Get("/captcha/:id/:width-:height", index.Captcha).Name("显示验证码")
	router.Post("/login", context.WrapCustomHandler[index.LoginParams](index.Login)).Name("管理员登录")

	authRouter.Post("/upload", context.WrapCustomHandler[context.NoRequestBody](index.Upload)).Name("上传文件")
	authRouter.Post("/index", context.WrapCustomHandler[context.NoRequestBody](index.Index)).Name("首页信息")
	authRouter.Post("/update", context.WrapCustomHandler[index.UpdateParams](index.Update)).Name("管理员更新")
	authRouter.Post("/update/password", context.WrapCustomHandler[index.UpdatePasswdParams](index.UpdatePasswd)).Name("管理员更新密码")

	// 管理列表
	authRouter.Post("/admins/manage/index", context.WrapCustomHandler[manage.IndexParams](manage.Index)).Name("管理列表")
	authRouter.Post("/admins/manage/views", context.WrapCustomHandler[context.NoRequestBody](manage.Views)).Name("管理视图")
	authRouter.Post("/admins/manage/create", context.WrapCustomHandler[manage.CreateParams](manage.Create)).Name("添加管理")
	authRouter.Post("/admins/manage/delete", context.WrapCustomHandler[context.DeleteParams](manage.Delete)).Name("删除管理")
	authRouter.Post("/admins/manage/update", context.WrapCustomHandler[manage.UpdateParams](manage.Update)).Name("更新管理")
	authRouter.Post("/admins/manage/setting", context.WrapCustomHandler[manage.SettingParams](manage.Setting)).Name("管理配置")
	authRouter.Post("/admins/manage/setting/sync", context.WrapCustomHandler[context.PrimaryKeyParams](manage.SettingSync)).Name("管理同步配置")
	authRouter.Post("/admins/manage/setting/reset", context.WrapCustomHandler[context.PrimaryKeyParams](manage.SettingReset)).Name("管理重置配置")
	authRouter.Post("/admins/manage/product/init", context.WrapCustomHandler[context.PrimaryKeyParams](manage.InitProduct)).Name("初始化店铺产品")

	// 管理日志
	authRouter.Post("/admins/logs/index", context.WrapCustomHandler[logs.IndexParams](logs.Index)).Name("管理日志列表")
	authRouter.Post("/admins/logs/views", context.WrapCustomHandler[context.NoRequestBody](logs.Views)).Name("管理日志视图")

	// 管理菜单
	authRouter.Post("/admins/menu/index", context.WrapCustomHandler[menu.IndexParams](menu.Index)).Name("管理菜单列表")
	authRouter.Post("/admins/menu/update", context.WrapCustomHandler[menu.UpdateParams](menu.Update)).Name("管理菜单更新")
	authRouter.Post("/admins/menu/create", context.WrapCustomHandler[menu.CreateParams](menu.Create)).Name("管理菜单新增")
	authRouter.Post("/admins/menu/delete", context.WrapCustomHandler[context.DeleteParams](menu.Delete)).Name("管理菜单删除")
	authRouter.Post("/admins/menu/setting", context.WrapCustomHandler[menu.SettingParams](menu.Setting)).Name("管理菜单配置")
	authRouter.Post("/admins/menu/views", context.WrapCustomHandler[context.NoRequestBody](menu.Views)).Name("管理菜单视图")

	// 管理角色
	authRouter.Post("/admins/role/index", context.WrapCustomHandler[role.IndexParams](role.Index)).Name("管理角色列表")
	authRouter.Post("/admins/role/update", context.WrapCustomHandler[role.UpdateParams](role.Update)).Name("管理角色更新")
	authRouter.Post("/admins/role/delete", context.WrapCustomHandler[context.DeleteParams](role.Delete)).Name("管理角色删除")
	authRouter.Post("/admins/role/create", context.WrapCustomHandler[role.CreateParams](role.Create)).Name("管理角色新增")
	authRouter.Post("/admins/role/setting", context.WrapCustomHandler[role.SettingParams](role.Setting)).Name("管理角色配置")
	authRouter.Post("/admins/role/auth", context.WrapCustomHandler[role.AuthParams](role.Auth)).Name("管理权限添加")
	authRouter.Post("/admins/role/views", context.WrapCustomHandler[context.NoRequestBody](role.Views)).Name("管理角色视图")

	// 管理配置
	authRouter.Post("/admins/setting/index", context.WrapCustomHandler[setting.IndexParams](setting.Index)).Name("管理配置列表")
	authRouter.Post("/admins/setting/create", context.WrapCustomHandler[setting.CreateParams](setting.Create)).Name("管理配置新增")
	authRouter.Post("/admins/setting/update", context.WrapCustomHandler[setting.UpdateParams](setting.Update)).Name("管理配置更新")
	authRouter.Post("/admins/setting/delete", context.WrapCustomHandler[context.DeleteParams](setting.Delete)).Name("管理配置删除")
	authRouter.Post("/admins/setting/views", context.WrapCustomHandler[context.NoRequestBody](setting.Views)).Name("管理配置视图")

	// 钱包资产
	authRouter.Post("/wallets/assets/index", context.WrapCustomHandler[assets.IndexParams](assets.Index)).Name("钱包资产列表")
	authRouter.Post("/wallets/assets/update", context.WrapCustomHandler[assets.UpdateParams](assets.Update)).Name("钱包资产更新")
	authRouter.Post("/wallets/assets/delete", context.WrapCustomHandler[context.DeleteParams](assets.Delete)).Name("钱包资产删除")
	authRouter.Post("/wallets/assets/create", context.WrapCustomHandler[assets.CreateParams](assets.Create)).Name("钱包资产新增")
	authRouter.Post("/wallets/assets/views", context.WrapCustomHandler[context.NoRequestBody](assets.Views)).Name("钱包资产视图")

	// 钱包支付
	authRouter.Post("/wallets/payment/index", context.WrapCustomHandler[payment.IndexParams](payment.Index)).Name("钱包支付列表")
	authRouter.Post("/wallets/payment/update", context.WrapCustomHandler[payment.UpdateParams](payment.Update)).Name("钱包支付更新")
	authRouter.Post("/wallets/payment/delete", context.WrapCustomHandler[context.DeleteParams](payment.Delete)).Name("钱包支付删除")
	authRouter.Post("/wallets/payment/create", context.WrapCustomHandler[payment.CreateParams](payment.Create)).Name("钱包支付新增")
	authRouter.Post("/wallets/payment/data", context.WrapCustomHandler[payment.DataParams](payment.Data)).Name("钱包支付数据")
	authRouter.Post("/wallets/payment/views", context.WrapCustomHandler[context.NoRequestBody](payment.Views)).Name("钱包支付视图")

	// 用户提现账户
	authRouter.Post("/users/account/index", context.WrapCustomHandler[userAccount.IndexParams](userAccount.Index)).Name("用户提现账户列表")
	authRouter.Post("/users/account/update", context.WrapCustomHandler[userAccount.UpdateParams](userAccount.Update)).Name("用户提现账户更新")
	authRouter.Post("/users/account/delete", context.WrapCustomHandler[context.DeleteParams](userAccount.Delete)).Name("用户提现账户删除")
	authRouter.Post("/users/account/create", context.WrapCustomHandler[userAccount.CreateParams](userAccount.Create)).Name("用户提现账户新增")
	authRouter.Post("/users/account/views", context.WrapCustomHandler[context.NoRequestBody](userAccount.Views)).Name("用户提现账户视图")

	// 用户资产列表
	authRouter.Post("/users/assets/index", context.WrapCustomHandler[userAssets.IndexParams](userAssets.Index)).Name("用户资产列表")
	authRouter.Post("/users/assets/update", context.WrapCustomHandler[userAssets.UpdateParams](userAssets.Update)).Name("用户资产更新")
	authRouter.Post("/users/assets/money", context.WrapCustomHandler[userAssets.MoneyParams](userAssets.Money)).Name("用户资产加减款")
	authRouter.Post("/users/assets/views", context.WrapCustomHandler[context.NoRequestBody](userAssets.Views)).Name("用户资产视图")

	// 用户账单
	authRouter.Post("/wallets/bill/index", context.WrapCustomHandler[userBill.IndexParams](userBill.Index)).Name("钱包账单列表")
	authRouter.Post("/wallets/bill/delete", context.WrapCustomHandler[context.DeleteParams](userBill.Delete)).Name("钱包账单删除")
	authRouter.Post("/wallets/bill/views", context.WrapCustomHandler[context.NoRequestBody](userBill.Views)).Name("钱包账单视图")

	// 钱包余额资产转换
	authRouter.Post("/wallets/convert/index", context.WrapCustomHandler[userConvert.IndexParams](userConvert.Index)).Name("钱包转换列表")
	authRouter.Post("/wallets/convert/delete", context.WrapCustomHandler[context.DeleteParams](userConvert.Delete)).Name("钱包转移删除")
	authRouter.Post("/wallets/convert/views", context.WrapCustomHandler[context.NoRequestBody](userConvert.Views)).Name("钱包转换视图")

	// 钱包余额资产转移
	authRouter.Post("/wallets/transfer/index", context.WrapCustomHandler[userTransfer.IndexParams](userTransfer.Index)).Name("钱包转移列表")
	authRouter.Post("/wallets/transfer/delete", context.WrapCustomHandler[context.DeleteParams](userTransfer.Delete)).Name("钱包转移删除")
	authRouter.Post("/wallets/transfer/views", context.WrapCustomHandler[context.NoRequestBody](userTransfer.Views)).Name("钱包转移视图")

	// 钱包订单管理
	authRouter.Post("/wallets/order/deposit/balance", middleware.PresetParams(&userOrder.PresetParams{Type: walletsModel.WalletUserOrderTypeDeposit}), context.WrapCustomHandler[userOrder.IndexParams](userOrder.Index)).Name("余额充值订单")
	authRouter.Post("/wallets/order/deposit/assets", middleware.PresetParams(&userOrder.PresetParams{Type: walletsModel.WalletUserOrderTypeAssetsDeposit}), context.WrapCustomHandler[userOrder.IndexParams](userOrder.Index)).Name("资产充值订单")
	authRouter.Post("/wallets/order/withdraw/balance", middleware.PresetParams(&userOrder.PresetParams{Type: walletsModel.WalletUserOrderTypeWithdraw}), context.WrapCustomHandler[userOrder.IndexParams](userOrder.Index)).Name("余额提现订单")
	authRouter.Post("/wallets/order/withdraw/assets", middleware.PresetParams(&userOrder.PresetParams{Type: walletsModel.WalletUserOrderTypeAssetsWithdraw}), context.WrapCustomHandler[userOrder.IndexParams](userOrder.Index)).Name("资产提现订单")
	authRouter.Post("/wallets/order/delete", context.WrapCustomHandler[context.DeleteParams](userOrder.Delete)).Name("钱包订单删除")
	authRouter.Post("/wallets/order/agree", context.WrapCustomHandler[userOrder.AgreeParams](userOrder.Agree)).Name("钱包订单同意")
	authRouter.Post("/wallets/order/refuse", context.WrapCustomHandler[userOrder.RefuseParams](userOrder.Refuse)).Name("钱包订单拒绝")
	authRouter.Post("/wallets/order/views", context.WrapCustomHandler[context.NoRequestBody](userOrder.Views)).Name("钱包订单视图")

	//	系统国家模块
	authRouter.Post("/systems/country/index", context.WrapCustomHandler[country.IndexParams](country.Index)).Name("国家列表")
	authRouter.Post("/systems/country/create", context.WrapCustomHandler[country.CreateParams](country.Create)).Name("新增国家")
	authRouter.Post("/systems/country/delete", context.WrapCustomHandler[context.DeleteParams](country.Delete)).Name("删除国家")
	authRouter.Post("/systems/country/update", context.WrapCustomHandler[country.UpdateParams](country.Update)).Name("更新国家")
	authRouter.Post("/systems/country/views", context.WrapCustomHandler[context.NoRequestBody](country.Views)).Name("国家视图")

	//	系统语言模块
	authRouter.Post("/systems/lang/index", context.WrapCustomHandler[lang.IndexParams](lang.Index)).Name("语言列表")
	authRouter.Post("/systems/lang/create", context.WrapCustomHandler[lang.CreateParams](lang.Create)).Name("新增语言")
	authRouter.Post("/systems/lang/delete", context.WrapCustomHandler[context.DeleteParams](lang.Delete)).Name("删除语言")
	authRouter.Post("/systems/lang/update", context.WrapCustomHandler[lang.UpdateParams](lang.Update)).Name("更新语言")
	authRouter.Post("/systems/lang/views", context.WrapCustomHandler[context.NoRequestBody](lang.Views)).Name("语言视图")

	//	系统等级模块
	authRouter.Post("/systems/level/index", context.WrapCustomHandler[level.IndexParams](level.Index)).Name("等级列表")
	authRouter.Post("/systems/level/create", context.WrapCustomHandler[level.CreateParams](level.Create)).Name("新增等级")
	authRouter.Post("/systems/level/update", context.WrapCustomHandler[level.UpdateParams](level.Update)).Name("更新等级")
	authRouter.Post("/systems/level/views", context.WrapCustomHandler[context.NoRequestBody](level.Views)).Name("等级视图")

	//	系统翻译模块
	authRouter.Post("/systems/translate/index", context.WrapCustomHandler[translate.IndexParams](translate.Index)).Name("翻译列表")
	authRouter.Post("/systems/translate/create", context.WrapCustomHandler[translate.CreateParams](translate.Create)).Name("新增翻译")
	authRouter.Post("/systems/translate/delete", context.WrapCustomHandler[context.DeleteParams](translate.Delete)).Name("删除翻译")
	authRouter.Post("/systems/translate/update", context.WrapCustomHandler[translate.UpdateParams](translate.Update)).Name("更新翻译")
	authRouter.Post("/systems/translate/lang", context.WrapCustomHandler[translate.LangParams](translate.Lang)).Name("翻译语言翻译")
	authRouter.Post("/systems/translate/views", context.WrapCustomHandler[context.NoRequestBody](translate.Views)).Name("翻译视图")

	//	系统通知模块
	authRouter.Post("/systems/notify/index", context.WrapCustomHandler[notify.IndexParams](notify.Index)).Name("通知列表")
	authRouter.Post("/systems/notify/create", context.WrapCustomHandler[notify.CreateParams](notify.Create)).Name("新增通知")
	authRouter.Post("/systems/notify/delete", context.WrapCustomHandler[context.DeleteParams](notify.Delete)).Name("删除通知")
	authRouter.Post("/systems/notify/update", context.WrapCustomHandler[notify.UpdateParams](notify.Update)).Name("更新通知")
	authRouter.Post("/systems/notify/views", context.WrapCustomHandler[context.NoRequestBody](notify.Views)).Name("通知视图")

	// 前台菜单模块
	authRouter.Post("/systems/menu/index", context.WrapCustomHandler[homeMenu.IndexParams](homeMenu.Index)).Name("前台菜单列表")
	authRouter.Post("/systems/menu/create", context.WrapCustomHandler[homeMenu.CreateParams](homeMenu.Create)).Name("前台菜单新增")
	authRouter.Post("/systems/menu/delete", context.WrapCustomHandler[context.DeleteParams](homeMenu.Delete)).Name("前台菜单删除")
	authRouter.Post("/systems/menu/update", context.WrapCustomHandler[homeMenu.UpdateParams](homeMenu.Update)).Name("前台菜单更新")
	authRouter.Post("/systems/menu/views", context.WrapCustomHandler[context.NoRequestBody](homeMenu.Views)).Name("前台菜单视图")

	// 文章管理
	authRouter.Post("/systems/article/index", context.WrapCustomHandler[article.IndexParams](article.Index)).Name("文章列表")
	authRouter.Post("/systems/article/create", context.WrapCustomHandler[article.CreateParams](article.Create)).Name("文章新增")
	authRouter.Post("/systems/article/update", context.WrapCustomHandler[article.UpdateParams](article.Update)).Name("文章更新")
	authRouter.Post("/systems/article/delete", context.WrapCustomHandler[context.DeleteParams](article.Delete)).Name("文章删除")
	authRouter.Post("/systems/article/views", context.WrapCustomHandler[context.NoRequestBody](article.Views)).Name("文章视图")

	// 用户表
	authRouter.Post("/users/user/index", context.WrapCustomHandler[user.IndexParams](user.Index)).Name("用户列表")
	authRouter.Post("/users/user/delete", context.WrapCustomHandler[context.DeleteParams](user.Delete)).Name("用户删除")
	authRouter.Post("/users/user/update", context.WrapCustomHandler[user.UpdateParams](user.Update)).Name("用户更新")
	authRouter.Post("/users/user/create", context.WrapCustomHandler[user.CreateParams](user.Create)).Name("用户新增")
	authRouter.Post("/users/user/virtual/create", context.WrapCustomHandler[user.CreateVirtualParams](user.CreateVirtual)).Name("虚拟用户新增")
	authRouter.Post("/users/user/money", context.WrapCustomHandler[user.MoneyParams](user.Money)).Name("用户资金变更")
	authRouter.Post("/users/user/views", context.WrapCustomHandler[context.NoRequestBody](user.Views)).Name("用户视图")

	// 渠道表
	authRouter.Post("/users/channel/index", context.WrapCustomHandler[channel.IndexParams](channel.Index)).Name("渠道列表")
	authRouter.Post("/users/channel/create", context.WrapCustomHandler[channel.CreateParams](channel.Create)).Name("新增渠道")
	authRouter.Post("/users/channel/delete", context.WrapCustomHandler[context.DeleteParams](channel.Delete)).Name("删除渠道")
	authRouter.Post("/users/channel/update", context.WrapCustomHandler[channel.UpdateParams](channel.Update)).Name("更新渠道")
	authRouter.Post("/users/channel/views", context.WrapCustomHandler[context.NoRequestBody](channel.Views)).Name("渠道视图")

	// 用户设置表
	authRouter.Post("/users/setting/update", context.WrapCustomHandler[usersSetting.UpdateParams](usersSetting.Update)).Name("用户更新")

	// 用户访问
	authRouter.Post("/users/access/index", context.WrapCustomHandler[access.IndexParams](access.Index)).Name("用户访问列表")
	authRouter.Post("/users/access/delete", context.WrapCustomHandler[context.DeleteParams](access.Delete)).Name("用户访问删除")
	authRouter.Post("/users/access/views", context.WrapCustomHandler[context.NoRequestBody](access.Views)).Name("用户访问视图")

	// 用户邀请
	authRouter.Post("/users/invite/index", context.WrapCustomHandler[invite.IndexParams](invite.Index)).Name("用户邀请列表")
	authRouter.Post("/users/invite/delete", context.WrapCustomHandler[context.DeleteParams](invite.Delete)).Name("用户邀请删除")
	authRouter.Post("/users/invite/create", context.WrapCustomHandler[invite.CreateParams](invite.Create)).Name("用户邀请新增")
	authRouter.Post("/users/invite/update", context.WrapCustomHandler[invite.UpdateParams](invite.Update)).Name("用户邀请更新")
	authRouter.Post("/users/invite/views", context.WrapCustomHandler[context.NoRequestBody](invite.Views)).Name("用户邀请视图")

	// 用户认证
	authRouter.Post("/users/auth/index", context.WrapCustomHandler[userAuth.IndexParams](userAuth.Index)).Name("用户认证列表")
	authRouter.Post("/users/auth/delete", context.WrapCustomHandler[context.DeleteParams](userAuth.Delete)).Name("用户认证删除")
	authRouter.Post("/users/auth/create", context.WrapCustomHandler[userAuth.CreateParams](userAuth.Create)).Name("用户认证新增")
	authRouter.Post("/users/auth/update", context.WrapCustomHandler[userAuth.UpdateParams](userAuth.Update)).Name("用户认证更新")
	authRouter.Post("/users/auth/agree", context.WrapCustomHandler[userAuth.AgreeParams](userAuth.Agree)).Name("用户认证同意")
	authRouter.Post("/users/auth/refuse", context.WrapCustomHandler[userAuth.RefuseParams](userAuth.Refuse)).Name("用户认证拒绝")
	authRouter.Post("/users/auth/views", context.WrapCustomHandler[context.NoRequestBody](userAuth.Views)).Name("用户认证视图")

	// 用户等级
	authRouter.Post("/users/level/index", context.WrapCustomHandler[userLevel.IndexParams](userLevel.Index)).Name("用户等级列表")
	authRouter.Post("/users/level/delete", context.WrapCustomHandler[context.DeleteParams](userLevel.Delete)).Name("用户等级删除")
	authRouter.Post("/users/level/update", context.WrapCustomHandler[userLevel.UpdateParams](userLevel.Update)).Name("用户等级更新")
	authRouter.Post("/users/level/create", context.WrapCustomHandler[userLevel.CreateParams](userLevel.Create)).Name("用户等级新增")
	authRouter.Post("/users/level/views", context.WrapCustomHandler[context.NoRequestBody](userLevel.Views)).Name("用户等级视图")

	// 产品分类
	authRouter.Post("/products/category/index", context.WrapCustomHandler[category.IndexParams](category.Index)).Name("产品分类列表")
	authRouter.Post("/products/category/delete", context.WrapCustomHandler[context.DeleteParams](category.Delete)).Name("产品分类删除")
	authRouter.Post("/products/category/update", context.WrapCustomHandler[category.UpdateParams](category.Update)).Name("产品分类更新")
	authRouter.Post("/products/category/crawling", context.WrapCustomHandler[category.CrawlingParams](category.Crawling)).Name("爬取产品分类数据")
	authRouter.Post("/products/category/create", context.WrapCustomHandler[category.CreateParams](category.Create)).Name("产品分类新增")
	authRouter.Post("/products/category/views", context.WrapCustomHandler[context.NoRequestBody](category.Views)).Name("产品分类视图")

	// 产品管理
	authRouter.Post("/products/product/index", context.WrapCustomHandler[product.IndexParams](product.Index)).Name("产品列表")
	authRouter.Post("/products/product/delete", context.WrapCustomHandler[context.DeleteParams](product.Delete)).Name("产品删除")
	authRouter.Post("/products/product/update", context.WrapCustomHandler[product.UpdateParams](product.Update)).Name("产品更新")
	authRouter.Post("/products/product/create", context.WrapCustomHandler[product.CreateParams](product.Create)).Name("产品新增")
	authRouter.Post("/products/product/comment/create", context.WrapCustomHandler[product.CreateCommentParams](product.CreateComment)).Name("创建产品评论")
	authRouter.Post("/products/product/crawling", context.WrapCustomHandler[product.CrawlingParams](product.Crawling)).Name("获取产品")
	authRouter.Post("/products/product/update/attrs", context.WrapCustomHandler[product.AttrsUpdateParams](product.UpdateAttrs)).Name("更新产品属性")
	authRouter.Post("/products/product/update/sku", context.WrapCustomHandler[product.UpdateSkuParams](product.UpdateSku)).Name("更新产品属性Sku")
	authRouter.Post("/products/product/views", context.WrapCustomHandler[context.NoRequestBody](product.Views)).Name("产品视图")

	// 产品订单
	authRouter.Post("/products/order/index", context.WrapCustomHandler[order.IndexParams](order.Index)).Name("产品订单列表")
	authRouter.Post("/products/order/delete", context.WrapCustomHandler[context.DeleteParams](order.Delete)).Name("产品订单删除")
	authRouter.Post("/products/order/update", context.WrapCustomHandler[order.UpdateParams](order.Update)).Name("产品订单更新")
	//authRouter.Post("/products/order/create", context.WrapCustomHandler[order.CreateParams](order.Create)).Name("产品订单新增")
	authRouter.Post("/products/order/views", context.WrapCustomHandler[context.NoRequestBody](order.Views)).Name("产品订单视图")

	// 购物车地址管理
	authRouter.Post("/shops/address/index", context.WrapCustomHandler[address.IndexParams](address.Index)).Name("购物车地址列表")
	authRouter.Post("/shops/address/create", context.WrapCustomHandler[address.CreateParams](address.Create)).Name("新增购物车地址")
	authRouter.Post("/shops/address/delete", context.WrapCustomHandler[context.DeleteParams](address.Delete)).Name("删除购物车地址")
	authRouter.Post("/shops/address/update", context.WrapCustomHandler[address.UpdateParams](address.Update)).Name("更新购物车地址")
	authRouter.Post("/shops/address/views", context.WrapCustomHandler[context.NoRequestBody](address.Views)).Name("购物车地址视图")

	// 购物车管理
	authRouter.Post("/shops/cart/index", context.WrapCustomHandler[cart.IndexParams](cart.Index)).Name("购物车列表")
	authRouter.Post("/shops/cart/delete", context.WrapCustomHandler[context.DeleteParams](cart.Delete)).Name("删除购物车")
	authRouter.Post("/shops/cart/update", context.WrapCustomHandler[cart.UpdateParams](cart.Update)).Name("更新购物车")
	authRouter.Post("/shops/cart/views", context.WrapCustomHandler[context.NoRequestBody](cart.Views)).Name("购物车视图")

	// 店铺管理
	authRouter.Post("/shops/store/index", context.WrapCustomHandler[store.IndexParams](store.Index)).Name("店铺列表")
	authRouter.Post("/shops/store/delete", context.WrapCustomHandler[context.DeleteParams](store.Delete)).Name("删除店铺")
	authRouter.Post("/shops/store/update", context.WrapCustomHandler[store.UpdateParams](store.Update)).Name("更新店铺")
	authRouter.Post("/shops/store/order/create", context.WrapCustomHandler[store.CreateOrderParams](store.CreateOrder)).Name("创建店铺订单")
	authRouter.Post("/shops/store/views", context.WrapCustomHandler[context.NoRequestBody](store.Views)).Name("店铺视图")

	// 店铺订单
	authRouter.Post("/shops/order/index", context.WrapCustomHandler[shopsOrder.IndexParams](shopsOrder.Index)).Name("店铺订单列表")
	authRouter.Post("/shops/order/complete", context.WrapCustomHandler[shopsOrder.CompleteParams](shopsOrder.Complete)).Name("店铺订单收货")
	authRouter.Post("/shops/order/shipping", context.WrapCustomHandler[shopsOrder.ShippingParams](shopsOrder.Shipping)).Name("店铺订单发货")
	authRouter.Post("/shops/order/views", context.WrapCustomHandler[context.NoRequestBody](shopsOrder.Views)).Name("店铺订单视图")

	// 店铺浏览记录
	authRouter.Post("/products/browsing/index", context.WrapCustomHandler[productBrowsing.IndexParams](productBrowsing.Index)).Name("店铺浏览记录列表")
	authRouter.Post("/products/browsing/views", context.WrapCustomHandler[context.NoRequestBody](productBrowsing.Views)).Name("店铺浏览记录视图")

	// 店铺入驻管理
	authRouter.Post("/shops/settled/index", context.WrapCustomHandler[settled.IndexParams](settled.Index)).Name("店铺入驻列表")
	authRouter.Post("/shops/settled/create", context.WrapCustomHandler[settled.CreateParams](settled.Create)).Name("新增店铺入驻")
	authRouter.Post("/shops/settled/delete", context.WrapCustomHandler[context.DeleteParams](settled.Delete)).Name("删除店铺入驻")
	authRouter.Post("/shops/settled/update", context.WrapCustomHandler[settled.UpdateParams](settled.Update)).Name("更新店铺入驻")
	authRouter.Post("/shops/settled/status", context.WrapCustomHandler[settled.StatusParams](settled.Status)).Name("店铺入驻审核")
	authRouter.Post("/shops/settled/views", context.WrapCustomHandler[context.NoRequestBody](settled.Views)).Name("店铺入驻视图")

	// 店铺商品关注收藏管理
	authRouter.Post("/shops/follow/index", context.WrapCustomHandler[follow.IndexParams](follow.Index)).Name("店铺商品关注收藏列表")
	authRouter.Post("/shops/follow/delete", context.WrapCustomHandler[context.DeleteParams](follow.Delete)).Name("删除店铺商品关注收藏")
	authRouter.Post("/shops/follow/update", context.WrapCustomHandler[follow.UpdateParams](follow.Update)).Name("更新店铺商品关注收藏")
	authRouter.Post("/shops/follow/views", context.WrapCustomHandler[context.NoRequestBody](follow.Views)).Name("店铺商品关注收藏视图")

	// 店铺商品评论管理
	authRouter.Post("/shops/comment/index", context.WrapCustomHandler[comment.IndexParams](comment.Index)).Name("商品评论列表")
	authRouter.Post("/shops/comment/delete", context.WrapCustomHandler[context.DeleteParams](comment.Delete)).Name("删除商品评论")
	authRouter.Post("/shops/comment/update", context.WrapCustomHandler[comment.UpdateParams](comment.Update)).Name("更新商品评论")
	authRouter.Post("/shops/comment/views", context.WrapCustomHandler[context.NoRequestBody](comment.Views)).Name("商品评论视图")

	// 店铺售后管理
	authRouter.Post("/shops/refund/index", context.WrapCustomHandler[refund.IndexParams](refund.Index)).Name("店铺售后列表")
	authRouter.Post("/shops/refund/delete", context.WrapCustomHandler[context.DeleteParams](refund.Delete)).Name("删除店铺售后")
	authRouter.Post("/shops/refund/update", context.WrapCustomHandler[refund.UpdateParams](refund.Update)).Name("更新店铺售后")
	authRouter.Post("/shops/refund/views", context.WrapCustomHandler[context.NoRequestBody](refund.Views)).Name("店铺售后视图")

	// 载入后端路由
	initAdminRoutes(app)
}

func middlewareAuthRouter(app *fiber.App) fiber.Router {
	priKey := commonService.NewServiceToken(nil).GetFileServiceRsaPrivate(adminsModel.ServiceAdminRouteName)
	return app.Group("/"+adminsModel.ServiceAdminAuthRouteName, middleware.InitJwtMiddleware(priKey, successHandler))
}

// initAdminRoutes 后端控制器路由
func initAdminRoutes(app *fiber.App) {
	routerList := app.GetRoutes()
	for i := 0; i < len(routerList); i++ {
		if routerList[i].Name != "" {
			adminRouterList[routerList[i].Path] = routerList[i].Name
		}
	}
}

// successHandler 成功返回值
func successHandler(ctx *fiber.Ctx) error {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()
	commService := commonService.NewServiceToken(rdsConn)

	adminId, userId := commService.GetClaims(ctx)
	if adminId != 0 {
		// 验证是否唯一Token， 验证管理是否过期, 验证是否白名单
		tokenStr := context.NewCustomCtx(ctx).GetToken()
		nowTime := time.Now().Unix()
		expireTime := adminsService.NewAdminUser(rdsConn, adminId).GetRedisExpiration()
		if commService.VerifyRedisToken(adminId, userId, tokenStr) && expireTime > nowTime {
			// 验证权限路由是否有效
			routerList := adminsService.NewAdminAuth(rdsConn, adminId).GetRedisAdminRouterList()
			if utils.ArrayStringIndexOf(routerList, ctx.Path()) > -1 {
				currentPathList := strings.Split(ctx.Path(), "/")
				if utils.ArrayStringIndexOf(adminFilterRouterList, currentPathList[len(currentPathList)-1]) == -1 {
					headersBytes, _ := json.Marshal(ctx.GetReqHeaders())
					database.Db.Create(&adminsModel.AdminLogs{
						AdminId: adminId,
						Name:    adminRouterList[ctx.Path()],
						Ip:      utils.GetClientIP(ctx),
						Headers: string(headersBytes),
						Route:   ctx.Path(),
						Body:    string(ctx.Body()),
					})
				}
				return ctx.Next()
			}
		}
	}

	//	缓存中不存在Token
	return ctx.Status(fiber.StatusUnauthorized).SendString("Invalid or expired JWT")
}
