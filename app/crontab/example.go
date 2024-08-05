package crontab

import "time"

func example(second time.Duration) {
	ch := time.NewTicker(second)

	for {
		<-ch.C
	}
}
