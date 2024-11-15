package timer

import (
	"fmt"
	"server/service"
	"time"
)

func OddsUpdate() {
	// 定期更新赔率
	ticker := time.NewTicker(1 * time.Minute) // 每分钟更新一次
	defer ticker.Stop()

	for range ticker.C {
		service.UpdateOdds()
		fmt.Println("赔率已更新")
	}
}
