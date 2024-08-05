package commonService

import (
	"crypto/rsa"
	"errors"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gomodule/redigo/redis"
	"gofiber/app/models/model/adminsModel"
	"gofiber/utils"
	"os"
	"strconv"
	"time"
)

const (
	// TokenExpired Token过期时间
	TokenExpired = 86400 * 30

	// RedisTokenName Token名称
	RedisTokenName = "_Token"

	RsaFilePath = "./rsa.json"
)

// TokenParams Token参数
type TokenParams struct {
	Name    string `json:"name"`    //	设备名称
	Token   string `json:"token"`   //	Token
	Expired int64  `json:"expired"` //	过期时间
}

type ServiceToken struct {
	rdsConn redis.Conn
}

func NewServiceToken(rdsConn redis.Conn) *ServiceToken {
	return &ServiceToken{rdsConn: rdsConn}
}

// GenerateHomeToken 生成前端Token
func (_ServiceToken *ServiceToken) GenerateHomeToken(adminId, userId uint) string {
	return _ServiceToken.generateToken(adminsModel.ServiceHomeName, adminId, userId, TokenExpired)
}

// GenerateAdminToken 生成后端Token
func (_ServiceToken *ServiceToken) GenerateAdminToken(adminId uint) string {
	return _ServiceToken.generateToken(adminsModel.ServiceAdminRouteName, adminId, 0, TokenExpired)
}

// VerifyToken 验证Token
func (_ServiceToken *ServiceToken) VerifyToken(serviceName string, tokenStr string) (uint, uint) {
	if tokenStr == "" {
		return 0, 0
	}

	privateKey := _ServiceToken.GetFileServiceRsaPrivate(serviceName)
	publicKey := privateKey.Public()

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("invalid or expired JWT")
		}
		return publicKey, nil
	})
	if err != nil {
		return 0, 0
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return uint(claims["adminId"].(float64)), uint(claims["userId"].(float64))
	}
	return 0, 0
}

// VerifyRedisToken 验证缓存中的Token
func (_ServiceToken *ServiceToken) VerifyRedisToken(adminId, userId uint, tokenStr string) bool {
	userTokenName := strconv.FormatInt(int64(adminId), 10) + "-" + strconv.FormatInt(int64(userId), 10)

	tokenParamsBytes, err := redis.Bytes(_ServiceToken.rdsConn.Do("HGET", RedisTokenName, userTokenName))
	if err != nil {
		return false
	}

	tokenParams := TokenParams{}
	_ = json.Unmarshal(tokenParamsBytes, &tokenParams)
	return tokenParams.Token == tokenStr
}

// generateToken 生成Token
func (_ServiceToken *ServiceToken) generateToken(serviceName string, adminId, userId uint, expired int) string {
	privateKey := _ServiceToken.GetFileServiceRsaPrivate(serviceName)
	claims := jwt.MapClaims{
		"adminId": adminId,
		"userId":  userId,
		"exp":     time.Now().Add(time.Second * time.Duration(expired)).Unix(),
	}

	//	生成Token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenStr, err := token.SignedString(privateKey)
	if err != nil {
		panic(err)
	}

	//	设置缓存Token
	_ServiceToken.setRedisTokenParams(adminId, userId, tokenStr, int64(expired))
	return tokenStr
}

// setRedisTokenParams 设置Token缓存
func (_ServiceToken *ServiceToken) setRedisTokenParams(adminId, userId uint, tokenStr string, expired int64) {
	//	唯一TokenKey
	userTokenName := strconv.FormatInt(int64(adminId), 10) + "-" + strconv.FormatInt(int64(userId), 10)
	tokenParams := &TokenParams{
		Name:    "",
		Token:   tokenStr,
		Expired: expired,
	}

	tokenParamsBytes, _ := json.Marshal(tokenParams)
	_, _ = _ServiceToken.rdsConn.Do("HSET", RedisTokenName, userTokenName, tokenParamsBytes)
}

// GetClaims 获取Claims
func (_ServiceToken *ServiceToken) GetClaims(ctx *fiber.Ctx) (uint, uint) {
	user := ctx.Locals("token").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if claims == nil {
		return 0, 0
	}

	return uint(claims["adminId"].(float64)), uint(claims["userId"].(float64))
}

// GetFileServiceRsaPrivate 获取文件rsa
func (_ServiceToken *ServiceToken) GetFileServiceRsaPrivate(serviceName string) *rsa.PrivateKey {
	serviceRsaList := _ServiceToken.GetServiceRsa()
	privateKeyStr := ""
	switch serviceName {
	case adminsModel.ServiceHomeName:
		privateKeyStr = serviceRsaList.Home.PriKey
	default:
		privateKeyStr = serviceRsaList.Admin.PriKey
	}
	//	解析私钥
	privateKey, err := utils.ParsePKCS1PrivateKey(privateKeyStr)
	if err != nil {
		panic(err)
	}
	return privateKey
}

// GetServiceRsa 获取服务Rsa文件
func (_ServiceToken *ServiceToken) GetServiceRsa() *adminsModel.AdminSettingServiceRsa {
	if !utils.PathExists(RsaFilePath) {
		// 生成RSA 权限, 保存到根目录
		admPriKey, admPubKey := utils.MarshalPKCS1PrivateKey()
		homePriKey, homePubKey := utils.MarshalPKCS1PrivateKey()
		rsaSettingList := &adminsModel.AdminSettingServiceRsa{
			Admin: &adminsModel.AdminSettingRsaInfo{
				PriKey: string(admPriKey),
				PubKey: string(admPubKey),
			},
			Home: &adminsModel.AdminSettingRsaInfo{
				PriKey: string(homePriKey),
				PubKey: string(homePubKey),
			},
		}
		if !utils.PathExists(RsaFilePath) {
			rsaBytes, _ := json.Marshal(rsaSettingList)
			_ = utils.FileWrite(RsaFilePath, rsaBytes)
		}
		return rsaSettingList
	}

	// 如果文件存在, 那么使用当前文件rsa内容
	rsaText, err := os.ReadFile(RsaFilePath)
	if err != nil {
		panic(err)
	}
	serviceRsaList := &adminsModel.AdminSettingServiceRsa{}
	_ = json.Unmarshal(rsaText, &serviceRsaList)
	if serviceRsaList.Admin == nil || serviceRsaList.Home == nil || serviceRsaList.Admin.PriKey == "" || serviceRsaList.Home.PriKey == "" {
		panic("Can't find RSA for JWT setup")
	}
	return serviceRsaList
}
