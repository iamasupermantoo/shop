package utils

import (
	"encoding/binary"
	"github.com/gofiber/fiber/v2"
	"net"
)

// GetClientIP 获取客户IP地址
func GetClientIP(ctx *fiber.Ctx) string {
	xForwarded := ctx.IPs()
	if len(xForwarded) > 0 {
		return xForwarded[0]
	}

	return ctx.IP()
}

// Ip4ToInt IP4转uint32
func Ip4ToInt(addr net.IP) uint32 {
	if len(addr) == 16 {
		return binary.BigEndian.Uint32(addr[12:16])
	}
	return binary.BigEndian.Uint32(addr)
}

// IntToIp4 整型转IP
func IntToIp4(n uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, n)
	return ip
}
