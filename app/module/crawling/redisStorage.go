package crawling

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"gofiber/app/module/cache"
	"net/url"
	"sync"
	"time"
)

// Storage colly 存储实例
type Storage struct {
	prefix  string
	u       *url.URL
	Expires time.Duration
	mu      sync.RWMutex
}

// NewStorage 新建存储实例
func NewStorage(prefix string) *Storage {
	return &Storage{
		Expires: 5,
		prefix:  prefix,
		mu:      sync.RWMutex{},
	}
}

// Init 实现初始化redis storage 接口
func (s *Storage) Init() error {
	return nil
}

// Visited 验证
func (s *Storage) Visited(requestID uint64) error {
	conn := cache.Rds.Get()
	defer conn.Close()
	_, err := conn.Do("SETEX", s.getIDStr(requestID), "1", s.Expires.String())
	if err != nil {
		zap.L().Error(LogMsg, zap.Error(err))
	}
	return err
}

// IsVisited 否通过验证
func (s *Storage) IsVisited(requestID uint64) (bool, error) {
	conn := cache.Rds.Get()
	defer conn.Close()
	_, err := conn.Do("GET", s.getIDStr(requestID))
	if errors.Is(err, redis.ErrNil) {
		return false, nil
	} else if err != nil {
		zap.L().Error(LogMsg, zap.Error(err))
		return false, err
	}
	return true, nil
}

// SetCookies 设置 cookie
func (s *Storage) SetCookies(u *url.URL, cookies string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.u = u
	conn := cache.Rds.Get()
	defer conn.Close()
	_, err := conn.Do("SETEX", s.getCookieID(u.Host), 10, cookies)
	if err != nil {
		zap.L().Error(LogMsg, zap.Error(err))
		return
	}
}

// Cookies 获取 cookie
func (s *Storage) Cookies(u *url.URL) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	conn := cache.Rds.Get()
	defer conn.Close()
	s.u = u
	cookiesStr, err := redis.String(conn.Do("GET", s.getCookieID(u.Host)))
	if err != nil {
		zap.L().Error(LogMsg, zap.Error(err))
		time.Sleep(100 * time.Millisecond)
		cookiesStr, err = redis.String(conn.Do("GET", s.getCookieID(u.Host)))
	}
	return cookiesStr
}

// AddRequest 添加请求头信息
func (s *Storage) AddRequest(r []byte) error {
	conn := cache.Rds.Get()
	defer conn.Close()
	_, err := conn.Do("RPUSH", s.getQueueID(), string(r))
	if err != nil {
		zap.L().Error(LogMsg, zap.Error(err))
	}
	return err
}

// GetRequest 获取请求信息
func (s *Storage) GetRequest() ([]byte, error) {
	conn := cache.Rds.Get()
	defer conn.Close()
	bytes, err := redis.Bytes(conn.Do("LPOP", s.getQueueID()))
	if err != nil {
		zap.L().Error(LogMsg, zap.Error(err))
	}
	return bytes, err
}

// QueueSize 获取队列大小
func (s *Storage) QueueSize() (int, error) {
	conn := cache.Rds.Get()
	defer conn.Close()
	bytes, err := redis.Int(conn.Do("LLEN", s.getQueueID()))
	if err != nil {
		zap.L().Error(LogMsg, zap.Error(err))
	}
	return bytes, err
}

// getIDStr 获取请求Key
func (s *Storage) getIDStr(ID uint64) string {
	return fmt.Sprintf("%s:request:%d", s.getPrefix(), ID)
}

// getCookieID 获取cookie key
func (s *Storage) getCookieID(c string) string {
	return fmt.Sprintf("%s:cookie:%s", s.getPrefix(), c)
}

// getCookieID 获取队列key
func (s *Storage) getQueueID() string {
	return fmt.Sprintf("%s:queue", s.getPrefix())
}

// getPrefix 获取域名作为前缀
func (s *Storage) getPrefix() string {
	return s.prefix
}
