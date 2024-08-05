package websocket

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/module/socket"
)

type AdminNotify struct {
	rdsConn      redis.Conn      //	Redis 对象
	adminId      uint            //	管理ID
	audioSetting map[string]bool // 配置设置
}

// NewAdminNotify 管理通知
func NewAdminNotify(rdsConn redis.Conn, adminId uint) *AdminNotify {
	adminSettingId := adminsService.NewAdminUser(rdsConn, adminId).GetRedisAdminSettingId(adminId)
	audioSetting := adminsService.NewAdminSetting(rdsConn, adminSettingId).CheckBoxToMaps("audioSetting")

	return &AdminNotify{rdsConn: rdsConn, adminId: adminId, audioSetting: audioSetting}
}

// Deposit 发送充值提示音
func (_AdminNotify *AdminNotify) Deposit(userInfo usersModel.User, paymentInfo walletsModel.WalletPayment, money float64) {
	notifyLabel := fmt.Sprintf("用户【%s】提交了 %s充值, 金额为【%v】,还未处理请及时处理～", paymentInfo.Name, userInfo.UserName, money)
	if _, ok := _AdminNotify.audioSetting["balanceDeposit"]; ok && _AdminNotify.audioSetting["balanceDeposit"] && paymentInfo.Mode == walletsModel.WalletPaymentModeDeposit {
		AdminWebSocket.RedisUserPublish(_AdminNotify.rdsConn, socket.MessageOperateAudio, _AdminNotify.adminId, map[string]interface{}{
			"label": notifyLabel, "source": "/assets/mp3/deposit.mp3", "type": "positive", "position": "top-right", "timeout": 3000,
		})
	}

	if _, ok := _AdminNotify.audioSetting["assetsDeposit"]; ok && _AdminNotify.audioSetting["assetsDeposit"] && paymentInfo.Mode == walletsModel.WalletPaymentModeAssetsDeposit {
		AdminWebSocket.RedisUserPublish(_AdminNotify.rdsConn, socket.MessageOperateAudio, _AdminNotify.adminId, map[string]string{
			"label": notifyLabel, "audio": "/assets/mp3/deposit.mp3", "color": "", "icon": "",
		})
	}
}

// Withdraw 发送提现提示音
func (_AdminNotify *AdminNotify) Withdraw(userInfo usersModel.User, paymentInfo walletsModel.WalletPayment, money float64) {
	notifyLabel := fmt.Sprintf("用户【%s】提交了 %s充值, 金额为【%v】,还未处理请及时处理～", paymentInfo.Name, userInfo.UserName, money)
	if _, ok := _AdminNotify.audioSetting["balanceWithdraw"]; ok && _AdminNotify.audioSetting["balanceWithdraw"] && paymentInfo.Mode == walletsModel.WalletPaymentModeWithdraw {
		AdminWebSocket.RedisUserPublish(_AdminNotify.rdsConn, socket.MessageOperateAudio, _AdminNotify.adminId, map[string]string{
			"label": notifyLabel, "audio": "/assets/mp3/withdraw.mp3", "color": "", "icon": "",
		})
	}

	if _, ok := _AdminNotify.audioSetting["assetsWithdraw"]; ok && _AdminNotify.audioSetting["assetsWithdraw"] && paymentInfo.Mode == walletsModel.WalletPaymentModeAssetsWithdraw {
		AdminWebSocket.RedisUserPublish(_AdminNotify.rdsConn, socket.MessageOperateAudio, _AdminNotify.adminId, map[string]string{
			"label": notifyLabel, "audio": "/assets/mp3/withdraw.mp3", "color": "", "icon": "",
		})
	}
}

// CreateOrder 发送创建订单提示音
func (_AdminNotify *AdminNotify) CreateOrder(label string) {
	if _, ok := _AdminNotify.audioSetting["createOrder"]; ok && _AdminNotify.audioSetting["createOrder"] {
		AdminWebSocket.RedisUserPublish(_AdminNotify.rdsConn, socket.MessageOperateAudio, _AdminNotify.adminId, map[string]string{
			"label": label, "audio": "/assets/mp3/tip.mp3", "color": "", "icon": "",
		})
	}
}
