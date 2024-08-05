package crontab

import "time"

func InitCrontab() {
	// 例子定时任务
	go example(30 * time.Second)
}
