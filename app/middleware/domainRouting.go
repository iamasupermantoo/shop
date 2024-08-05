package middleware

import (
	"github.com/gofiber/fiber/v2"
	"gofiber/app/models/model/adminsModel"
	"strings"
)

type Service struct {
	Fiber *fiber.App
}

// CurrentServiceList 当前服务列表
var CurrentServiceList = map[string]*Service{}

// InitService 初始化运行服务
func InitService(serviceList map[string]*Service) {
	CurrentServiceList = serviceList
}

// InitDomainRoutingMiddleware 初始化域名路由中间件
func InitDomainRoutingMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if strings.Index(c.Path(), adminsModel.ServiceAdminRouteName) == 1 {
			c.Locals("ServiceName", adminsModel.ServiceAdminRouteName)
			CurrentServiceList[adminsModel.ServiceAdminRouteName].Fiber.Handler()(c.Context())
		} else {
			c.Locals("ServiceName", adminsModel.ServiceHomeName)
			CurrentServiceList[adminsModel.ServiceHomeName].Fiber.Handler()(c.Context())
		}
		return nil
	}
}
