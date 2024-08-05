package cache

// Publish 发送订阅消息
func (_SubScribe *SubScribe) Publish(name string, data interface{}) {
	rdsConn := Rds.Get()
	defer rdsConn.Close()

	_, _ = rdsConn.Do("PUBLISH", name, data)
}
