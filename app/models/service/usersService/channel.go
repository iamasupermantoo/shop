package usersService

import (
	"errors"
	"github.com/goccy/go-json"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/module/context"
	"gofiber/utils"
	"net/url"
)

type UserChannel struct {
	requestURL  string              //	请求地址
	channelInfo *usersModel.Channel //	渠道信息
}

// NewUserChannel 创建渠道对象
func NewUserChannel(channelInfo *usersModel.Channel) *UserChannel {
	requestURL := ""
	channelURL, err := url.Parse(channelInfo.Route)
	if err == nil {
		requestURL = channelURL.Scheme + "://" + channelURL.Hostname() + channelURL.Path
		if channelURL.Port() != "" {
			requestURL += ":" + channelURL.Port()
		}
	}
	return &UserChannel{channelInfo: channelInfo, requestURL: requestURL}
}

// ApproveLogin 授权登录
func (_UserChannel *UserChannel) ApproveLogin(params *ApproveLoginParams) (string, error) {
	params.Sign = utils.StructSign(params, _UserChannel.channelInfo.Pass)
	resBytes, err := utils.NewClient().SetHeaders(map[string]string{"Content-Type": "application/json"}).
		Request("POST", _UserChannel.requestURL+"/users/channel/approve", params)
	if err != nil {
		return "", err
	}
	resData := &context.RespJson{}
	_ = json.Unmarshal(resBytes, &resData)
	if resData.Code != 0 {
		return "", errors.New(resData.Msg)
	}
	return resData.Data.(string), nil
}

// ApproveDeposit 授权充值
func (_UserChannel *UserChannel) ApproveDeposit(params *ApproveDeposit) error {
	params.Sign = utils.StructSign(params, _UserChannel.channelInfo.Pass)

	resBytes, err := utils.NewClient().SetHeaders(map[string]string{"Content-Type": "application/json"}).
		Request("POST", _UserChannel.requestURL+"/users/channel/withdraw", params)
	if err != nil {
		return err
	}
	resData := &context.RespJson{}
	_ = json.Unmarshal(resBytes, &resData)
	if resData.Code != 0 {
		return errors.New(resData.Msg)
	}
	return nil
}

type ApproveDeposit struct {
	Symbol string  `json:"symbol"` //	渠道标识
	User   string  `json:"user"`   //	用户名
	Money  float64 `json:"money"`  //	金额
	Time   int64   `json:"time"`   //	时间
	Sign   string  `json:"sign"`   //	加密
}

type ApproveLoginParams struct {
	Symbol string `json:"symbol"` //	渠道标识
	User   string `json:"user"`   //	用户名
	Pass   string `json:"pass"`   //	密码
	Time   int64  `json:"time"`   //	时间
	Sign   string `json:"sign"`   //	加密
}
