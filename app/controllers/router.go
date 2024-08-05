package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gofiber/app/controllers/chats/conversation"
	"gofiber/app/controllers/chats/messages"
	"gofiber/app/controllers/index"
	"gofiber/app/controllers/stores/address"
	"gofiber/app/controllers/stores/browsing"
	"gofiber/app/controllers/stores/comment"
	"gofiber/app/controllers/stores/follow"
	storeComment "gofiber/app/controllers/stores/merchants/comment"
	storeOrder "gofiber/app/controllers/stores/merchants/order"
	"gofiber/app/controllers/stores/merchants/refund"
	"gofiber/app/controllers/stores/merchants/store"
	"gofiber/app/controllers/stores/merchants/wholesale"
	userOrder "gofiber/app/controllers/stores/order"
	"gofiber/app/controllers/stores/product"
	"gofiber/app/controllers/stores/refund"
	"gofiber/app/controllers/stores/settled"
	"gofiber/app/controllers/stores/shippingCart"
	userStore "gofiber/app/controllers/stores/store"
	"gofiber/app/controllers/systems"
	"gofiber/app/controllers/systems/article"
	"gofiber/app/controllers/users/auth"
	"gofiber/app/controllers/users/channel"
	"gofiber/app/controllers/users/level"
	"gofiber/app/controllers/users/team"
	"gofiber/app/controllers/users/user"
	"gofiber/app/controllers/wallets/account"
	"gofiber/app/controllers/wallets/assets"
	"gofiber/app/controllers/wallets/bill"
	"gofiber/app/controllers/wallets/convert"
	"gofiber/app/controllers/wallets/order"
	"gofiber/app/controllers/wallets/transfer"
	"gofiber/app/middleware"
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/service/commonService"
	"gofiber/app/models/service/usersService"
	"gofiber/app/module/cache"
	"gofiber/app/module/context"
	"gofiber/app/websocket"
	"gofiber/module/socket/handler"
)

// InitWebRouter 初始化前台App
func InitWebRouter(app *fiber.App) {
	authRouter := middlewareAuthRouter(app)

	router := app.Group("/")
	router.Get("/ws", handler.NewSocketConn(websocket.HomeWebSocket)).Name("Websocket")
	// 初始化一些websocket数据

	router.Get("/", context.WrapCustomHandler[context.NoRequestBody](index.Index)).Name("接口信息")
	router.Get("/captcha/create", index.NewCaptcha).Name("创建验证码")
	router.Get("/captcha/:id/:width-:height", index.Captcha).Name("显示验证码")
	router.Get("/init", context.WrapCustomHandler[context.NoRequestBody](index.Init)).Name("初始化数据")
	router.Post("/online", context.WrapCustomHandler[context.NoRequestBody](index.Online)).Name("请求客服链接")
	router.Post("/translate", context.WrapCustomHandler[context.NoRequestBody](systems.Translate)).Name("获取语言包")
	router.Post("/download", context.WrapCustomHandler[context.NoRequestBody](index.Download)).Name("下载信息")
	router.Post("/footer", context.WrapCustomHandler[context.NoRequestBody](index.Footer)).Name("底部信息")
	router.Post("/article/index", context.WrapCustomHandler[context.NoRequestBody](article.Index)).Name("文章列表")
	router.Post("/article/details", context.WrapCustomHandler[article.DetailsParams](article.Details)).Name("文章详情")
	router.Post("/access", context.WrapCustomHandler[index.AccessParams](index.Access)).Name("前台访问记录")
	router.Post("/login", context.WrapCustomHandler[user.LoginParams](user.Login)).Name("用户登录")
	router.Post("/register", context.WrapCustomHandler[user.RegisterParams](user.Register)).Name("用户注册")
	router.Post("/home", context.WrapCustomHandler[context.NoRequestBody](index.Home)).Name("首页信息")
	authRouter.Post("/upload", context.WrapCustomHandler[context.NoRequestBody](index.Upload)).Name("上传文件")

	router.Post("/product/index", context.WrapCustomHandler[product.IndexParams](product.Index)).Name("商品列表")
	router.Post("/product/details", context.WrapCustomHandler[product.DetailParams](product.Details)).Name("商品详情")
	router.Post("/product/category", context.WrapCustomHandler[context.NoRequestBody](product.Category)).Name("商品分类")

	// 聊天模块
	authRouter.Post("/chats/conversation/index", context.WrapCustomHandler[conversation.IndexParams](conversation.Index)).Name("会话列表")
	authRouter.Post("/chats/conversation/details", context.WrapCustomHandler[conversation.DetailsParams](conversation.Details)).Name("会话详情")

	// 聊天消息
	authRouter.Post("/chats/messages/index", context.WrapCustomHandler[messages.IndexParams](messages.Index)).Name("消息列表")
	authRouter.Post("/chats/messages/cancel", context.WrapCustomHandler[messages.CancelParams](messages.Cancel)).Name("消息撤回")
	authRouter.Post("/chats/messages/read", context.WrapCustomHandler[messages.ReadParams](messages.Read)).Name("消息已读")
	authRouter.Post("/chats/messages/send", context.WrapCustomHandler[messages.SendParams](messages.Send)).Name("发送消息")

	// 用户渠道
	authRouter.Post("/users/channel/login", context.WrapCustomHandler[channel.ApproveLoginParams](channel.ApproveLogin)).Name("授权用户登录")
	router.Post("/users/channel/approve", context.WrapCustomHandler[usersService.ApproveLoginParams](channel.Approve)).Name("授权登录处理")
	router.Post("/users/channel/withdraw", context.WrapCustomHandler[usersService.ApproveDeposit](channel.Withdraw)).Name("授权提现处理")

	// 用户模块
	authRouter.Post("/users/info", context.WrapCustomHandler[context.NoRequestBody](user.Info)).Name("用户信息")
	authRouter.Post("/users/invite", context.WrapCustomHandler[context.NoRequestBody](user.Invite)).Name("邀请信息")
	authRouter.Post("/users/update", context.WrapCustomHandler[user.UpdateParams](user.Update)).Name("更新用户信息")
	authRouter.Post("/users/update/password", context.WrapCustomHandler[user.UpdatePasswordParams](user.UpdatePassword)).Name("更新用户密码")

	authRouter.Post("/users/auth/create", context.WrapCustomHandler[auth.CreateParams](auth.Create)).Name("申请认证")
	authRouter.Post("/users/auth/info", context.WrapCustomHandler[context.NoRequestBody](auth.Info)).Name("认证信息")
	authRouter.Post("/users/level/index", context.WrapCustomHandler[context.NoRequestBody](level.Index)).Name("等级列表")
	authRouter.Post("/users/level/create", context.WrapCustomHandler[level.CreateParams](level.Create)).Name("购买等级")

	authRouter.Post("/users/team/index", context.WrapCustomHandler[context.NoRequestBody](team.Index)).Name("我的团队")
	authRouter.Post("/users/team/details", context.WrapCustomHandler[context.NoRequestBody](team.Details)).Name("团队收益")

	// 提现账户
	authRouter.Post("/wallets/account/index", context.WrapCustomHandler[account.IndexParams](account.Index)).Name("提现账户列表")
	authRouter.Post("/wallets/account/info", context.WrapCustomHandler[account.InfoParams](account.Info)).Name("提现账户信息")
	authRouter.Post("/wallets/account/payment", context.WrapCustomHandler[account.PaymentParams](account.Payment)).Name("提现账户类型")
	authRouter.Post("/wallets/account/create", context.WrapCustomHandler[account.CreateParams](account.Create)).Name("提现账户新增")
	authRouter.Post("/wallets/account/update", context.WrapCustomHandler[account.UpdateParams](account.Update)).Name("提现账户更新")
	authRouter.Post("/wallets/account/delete", context.WrapCustomHandler[account.DeleteParams](account.Delete)).Name("提现账户删除")

	// 我的资产
	authRouter.Post("/wallets/assets/assets", context.WrapCustomHandler[context.NoRequestBody](assets.Assets)).Name("钱包资产列表")
	authRouter.Post("/wallets/assets/index", context.WrapCustomHandler[context.NoRequestBody](assets.Index)).Name("用户资产列表")
	authRouter.Post("/wallets/assets/details", context.WrapCustomHandler[assets.DetailsParams](assets.Details)).Name("用户资产详情")

	// 钱包订单
	authRouter.Post("/wallets/order/deposit", context.WrapCustomHandler[order.DepositParams](order.Deposit)).Name("钱包订单充值")
	authRouter.Post("/wallets/order/withdraw", context.WrapCustomHandler[order.WithdrawParams](order.Withdraw)).Name("钱包订单提现")
	authRouter.Post("/wallets/order/index", context.WrapCustomHandler[order.IndexParams](order.Index)).Name("钱包订单列表")

	// 钱包账单
	authRouter.Post("/wallets/bill/index", context.WrapCustomHandler[bill.IndexParams](bill.Index)).Name("钱包账单列表")
	authRouter.Post("/wallets/bill/options", context.WrapCustomHandler[bill.OptionsParams](bill.Options)).Name("钱包账单类型")

	// 资金转移
	authRouter.Post("/wallets/transfer/create", context.WrapCustomHandler[transfer.CreateParams](transfer.Create)).Name("资金转移申请")
	authRouter.Post("/wallets/transfer/index", context.WrapCustomHandler[transfer.IndexParams](transfer.Index)).Name("资金转移列表")

	// 资金转换
	authRouter.Post("/wallets/convert/create", context.WrapCustomHandler[convert.CreateParams](convert.Create)).Name("资金转换申请")
	authRouter.Post("/wallets/convert/index", context.WrapCustomHandler[convert.IndexParams](convert.Index)).Name("资金转换列表")
	authRouter.Post("/wallets/convert/info", context.WrapCustomHandler[convert.InfoParams](convert.Info)).Name("资金转换信息")

	// 用户绑定收货地址
	authRouter.Post("/stores/address/index", context.WrapCustomHandler[context.NoRequestBody](address.Index)).Name("收货地址列表")
	authRouter.Post("/stores/address/info", context.WrapCustomHandler[address.InfoParams](address.Info)).Name("收货地址信息")
	authRouter.Post("/stores/address/create", context.WrapCustomHandler[address.CreateParams](address.Create)).Name("新增收货地址")
	authRouter.Post("/stores/address/delete", context.WrapCustomHandler[address.DeleteParams](address.Delete)).Name("删除收货地址")
	authRouter.Post("/stores/address/update", context.WrapCustomHandler[address.UpdateParams](address.Update)).Name("更新收货地址")

	// 产品关注收藏
	authRouter.Post("/stores/follow/index", context.WrapCustomHandler[follow.IndexParams](follow.Index)).Name("收藏关注列表")
	authRouter.Post("/stores/follow/create", context.WrapCustomHandler[follow.CreateParams](follow.Create)).Name("收藏关注接口")

	// 产品浏览记录
	authRouter.Post("/stores/browsing/create", context.WrapCustomHandler[browsing.CreateParams](browsing.Create)).Name("浏览记录记录")
	authRouter.Post("/stores/browsing/index", context.WrapCustomHandler[browsing.IndexParams](browsing.Index)).Name("浏览记录列表")

	// 用户购物车
	authRouter.Post("/stores/cart/index", context.WrapCustomHandler[context.NoRequestBody](shippingCart.Index)).Name("产品购物车列表")
	authRouter.Post("/stores/cart/create", context.WrapCustomHandler[shippingCart.CreateParams](shippingCart.Create)).Name("添加购物车")
	authRouter.Post("/stores/cart/delete", context.WrapCustomHandler[shippingCart.DeleteParams](shippingCart.Delete)).Name("删除购物车")
	authRouter.Post("/stores/cart/update", context.WrapCustomHandler[shippingCart.UpdateParams](shippingCart.Update)).Name("更新购物车")

	// 用户订单
	authRouter.Post("/stores/order/index", context.WrapCustomHandler[userOrder.IndexParams](userOrder.Index)).Name("订单列表")
	authRouter.Post("/stores/order/details", context.WrapCustomHandler[userOrder.DetailsParams](userOrder.Details)).Name("订单详情")
	authRouter.Post("/stores/order/create", context.WrapCustomHandler[userOrder.CreateParams](userOrder.Create)).Name("创建订单")
	authRouter.Post("/stores/order/delete", context.WrapCustomHandler[userOrder.DeleteParams](userOrder.Delete)).Name("订单删除")
	authRouter.Post("/stores/order/complete", context.WrapCustomHandler[userOrder.CompleteParams](userOrder.Complete)).Name("订单完成")
	authRouter.Post("/stores/order/cancel", context.WrapCustomHandler[userOrder.CancelParams](userOrder.Cancel)).Name("订单取消")
	//authRouter.Post("/stores/order/pay", context.WrapCustomHandler[userOrder.PayParams](userOrder.Pay)).Name("订单支付")
	authRouter.Post("/stores/order/submitIndex", context.WrapCustomHandler[userOrder.SubmitIndexParams](userOrder.SubmitIndex)).Name("提交订单页面信息")

	// 商家订单
	authRouter.Post("/stores/merchant/order/index", context.WrapCustomHandler[storeOrder.IndexParams](storeOrder.Index)).Name("商家订单列表")
	authRouter.Post("/stores/merchant/order/details", context.WrapCustomHandler[storeOrder.DetailsParams](storeOrder.Details)).Name("商家订单详情")
	authRouter.Post("/stores/merchant/order/shipping", context.WrapCustomHandler[storeOrder.ShippingParams](storeOrder.Shipping)).Name("商家订单发货")

	// 订单售后
	authRouter.Post("/stores/refund/create", context.WrapCustomHandler[refund.CreateParams](refund.Create)).Name("售后申请")
	authRouter.Post("/stores/refund/details", context.WrapCustomHandler[refund.DetailsParams](refund.Details)).Name("售后详情")
	authRouter.Post("/stores/refund/index", context.WrapCustomHandler[refund.IndexParams](refund.Index)).Name("售后列表")

	// 商家售后
	authRouter.Post("/stores/merchant/refund/index", context.WrapCustomHandler[storeRefund.IndexParams](storeRefund.Index)).Name("商家售后列表")
	authRouter.Post("/stores/merchant/refund/details", context.WrapCustomHandler[storeRefund.DetailsParams](storeRefund.Details)).Name("商家售后详情")
	authRouter.Post("/stores/merchant/refund/agree", context.WrapCustomHandler[storeRefund.AgreeParams](storeRefund.Agree)).Name("商家同意售后")
	authRouter.Post("/stores/merchant/refund/refuse", context.WrapCustomHandler[storeRefund.RefuseParams](storeRefund.Refuse)).Name("商家拒绝售后")

	// 订单评论
	authRouter.Post("/stores/comment/create", context.WrapCustomHandler[comment.CreateParams](comment.Create)).Name("申请评论")
	authRouter.Post("/stores/comment/info", context.WrapCustomHandler[comment.InfoParams](comment.Info)).Name("评论信息")
	authRouter.Post("/stores/comment/index", context.WrapCustomHandler[comment.IndexParams](comment.Index)).Name("评论列表")
	authRouter.Post("/stores/merchant/comment/index", context.WrapCustomHandler[storeComment.IndexParams](storeComment.Index)).Name("商家评论列表")
	authRouter.Post("/stores/merchant/product/comment/index", context.WrapCustomHandler[storeComment.ProductParams](storeComment.Product)).Name("商家评论列表")

	// 店铺入驻管理
	authRouter.Post("/stores/settled/info", context.WrapCustomHandler[context.NoRequestBody](settled.Info)).Name("商家入驻信息")
	authRouter.Post("/stores/settled/create", context.WrapCustomHandler[settled.CreateParams](settled.Create)).Name("商家入驻申请")

	// 店铺详细信息
	authRouter.Post("/stores/merchant/store/home", context.WrapCustomHandler[context.NoRequestBody](store.Home)).Name("商家店铺首页")
	authRouter.Post("/stores/merchant/store/info", context.WrapCustomHandler[context.NoRequestBody](store.Info)).Name("商家店铺信息")
	authRouter.Post("/stores/merchant/store/update", context.WrapCustomHandler[store.UpdateParams](store.Update)).Name("商家店铺更新")
	router.Post("/stores/store/details", context.WrapCustomHandler[userStore.DetailsParams](userStore.Details)).Name("店铺详情")
	router.Post("/stores/store/index", context.WrapCustomHandler[userStore.IndexParams](userStore.Index)).Name("店铺列表")

	// 批发产品
	authRouter.Post("/stores/wholesale/index", context.WrapCustomHandler[wholesale.IndexParams](wholesale.Index)).Name("批发中心列表")
	authRouter.Post("/stores/wholesale/shelves", context.WrapCustomHandler[wholesale.ShelvesParams](wholesale.Shelves)).Name("批发商品上架")
	authRouter.Post("/stores/wholesale/status", context.WrapCustomHandler[wholesale.StatusParams](wholesale.Status)).Name("商家商品上架下架")
	authRouter.Post("/stores/wholesale/details", context.WrapCustomHandler[wholesale.DetailsParams](wholesale.Details)).Name("商家商品详情")
	authRouter.Post("/stores/wholesale/update", context.WrapCustomHandler[wholesale.UpdateParams](wholesale.Update)).Name("商家商品更新")
}

func middlewareAuthRouter(app *fiber.App) fiber.Router {
	priKey := commonService.NewServiceToken(nil).GetFileServiceRsaPrivate(adminsModel.ServiceHomeName)
	return app.Group("/auth", middleware.InitJwtMiddleware(priKey, successHandler))
}

func successHandler(ctx *fiber.Ctx) error {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()
	commService := commonService.NewServiceToken(rdsConn)

	adminId, userId := commService.GetClaims(ctx)
	if userId > 0 {
		tokenStr := context.NewCustomCtx(ctx).GetToken()
		if commService.VerifyRedisToken(adminId, userId, tokenStr) {
			return ctx.Next()
		}
	}

	return ctx.Status(fiber.StatusUnauthorized).SendString("Invalid or expired JWT")
}
