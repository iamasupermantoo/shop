package websocket

import (
	"github.com/goccy/go-json"
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/commonService"
	"gofiber/app/module/cache"
	"gofiber/app/module/database"
	"gofiber/module/socket"
)

var AdminWebSocket *socket.Socket

// InitAdminWebSocket 初始化后台WebSocket
func InitAdminWebSocket() {
	AdminWebSocket = socket.NewSocket(adminsModel.ServiceAdminRouteName)
	AdminWebSocket.SetEventMessage(OnAdminWebSocketMessageFunc)
	AdminWebSocket.SetEventClose(OnAdminWebSocketCloseFunc)
	AdminWebSocket.SetEventOpen(OnAdminWebSocketOpenFunc)
}

// OnAdminWebSocketMessageFunc 消息事件
func OnAdminWebSocketMessageFunc(socketInterface *socket.Socket, uuidStr string, msg []byte) error {
	dataMsg := &socket.ClientMessage{}
	err := json.Unmarshal(msg, &dataMsg)
	if err != nil {
		return err
	}
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	switch dataMsg.Op {
	case socket.MessageOperateSubscribe:
		// 订阅通道
		subscribe := &socket.SubscribeInfo{}
		dataBytes, _ := json.Marshal(dataMsg.Data)
		err = json.Unmarshal(dataBytes, &subscribe)
		if err != nil {
			return err
		}

		if subscribe.Channel != "" && len(subscribe.Args) > 0 {
			socketInterface.Subscribe(rdsConn, uuidStr, subscribe.Channel, subscribe.Args)
			socketInterface.RedisPublish(socket.MessageOperateSubscribe, uuidStr, &socket.ClientMessage{
				Op:   socket.MessageOperateSubscribe,
				Data: "ok",
			})
		}

	case socket.MessageOperateUnSubscribe:
		// 取消订阅通道
		subscribe := &socket.SubscribeInfo{}
		dataBytes, _ := json.Marshal(dataMsg.Data)
		err = json.Unmarshal(dataBytes, &subscribe)
		if err != nil {
			return err
		}

		if subscribe.Channel != "" && len(subscribe.Args) > 0 {
			socketInterface.UnSubscribe(rdsConn, uuidStr, subscribe.Channel, subscribe.Args)
			socketInterface.RedisPublish(socket.MessageOperateUnSubscribe, uuidStr, &socket.ClientMessage{
				Op:   socket.MessageOperateUnSubscribe,
				Data: "ok",
			})
		}
	case socket.MessageOperateBindUser:
		adminId, userId := commonService.NewServiceToken(rdsConn).VerifyToken(adminsModel.ServiceAdminRouteName, dataMsg.Data.(string))
		if adminId > 0 {
			// 设置当前用户ID
			socketInterface.ConnMaps.SetConnUserId(uuidStr, adminId)

			// 绑定用户ID 后台绑定的是管理ID
			socketInterface.ConnMaps.RedisSetConnInfo(rdsConn, adminId, &socket.UserConnInfo{
				UUID: uuidStr, Key: socketInterface.GetSocketKey(), UserId: userId, AdminId: adminId,
				Origin: "", Device: "", IP: "",
			})

			socketInterface.RedisPublish(socket.MessageOperateBindUser, uuidStr, &socket.ClientMessage{
				Op:   socket.MessageOperateBindUser,
				Data: "ok",
			})

			// 上线通知
			socketInterface.EventOpen(socketInterface, uuidStr)
		}
	default:
	}

	return nil
}

// OnAdminWebSocketCloseFunc 关闭事件方法
func OnAdminWebSocketCloseFunc(socketInterface *socket.Socket, uuidStr string) {

}

// OnAdminWebSocketOpenFunc 打开事件方法
func OnAdminWebSocketOpenFunc(socketInterface *socket.Socket, uuidStr string) {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	connInfo := socketInterface.ConnMaps.GetConn(uuidStr)
	notifyService := NewAdminNotify(rdsConn, connInfo.UserId)

	// 初始化余额充值
	depositList := make([]*walletUserOrder, 0)
	database.Db.Model(&walletsModel.WalletUserOrder{}).Preload("UserInfo").Preload("PaymentInfo").
		Where("admin_id = ?", connInfo.UserId).Where("type IN ?", []int{walletsModel.WalletUserOrderTypeDeposit, walletsModel.WalletUserOrderTypeAssetsDeposit}).Where("status = ?", walletsModel.WalletUserOrderStatusActive).Find(&depositList)
	for _, orderInfo := range depositList {
		notifyService.Deposit(orderInfo.UserInfo, orderInfo.PaymentInfo, orderInfo.Money)
	}

	// 初始化余额提现
	withdrawList := make([]*walletUserOrder, 0)
	database.Db.Model(&walletsModel.WalletUserOrder{}).Preload("UserInfo").Preload("PaymentInfo").
		Where("admin_id = ?", connInfo.UserId).Where("type IN ?", []int{walletsModel.WalletUserOrderTypeWithdraw, walletsModel.WalletUserOrderTypeAssetsWithdraw}).Where("status = ?", walletsModel.WalletUserOrderStatusActive).Find(&withdrawList)
	for _, orderInfo := range withdrawList {
		notifyService.Withdraw(orderInfo.UserInfo, orderInfo.PaymentInfo, orderInfo.Money)
	}
}

type walletUserOrder struct {
	walletsModel.WalletUserOrder
	UserInfo    usersModel.User            `gorm:"foreignKey:UserId;" json:"userInfo"`
	PaymentInfo walletsModel.WalletPayment `gorm:"foreignKey:SourceId" json:"paymentInfo"`
}
