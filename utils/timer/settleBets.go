package timer

import (
	"fmt"
	"server/service"
	"time"
)

func SettleBets() {
	// 启动定时器，每分钟轮询一次
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if err := service.CheckAndSettleBets(); err != nil {
				fmt.Println("Failed to check and settle bets:", err)
			}
			fmt.Println("预测已结算")
		}
	}
}
