package service

import (
	"fmt"
	"server/dao/mysql"
	"server/model"
	"strconv"
)

func UpdateOdds() {
	// 获取当前的投注数据
	items, err := mysql.ItemList()
	if err != nil {
		fmt.Println("Get bet items failed:", err)
		return
	}

	// 按 bet_list_id 分组
	betGroups := make(map[int64][]model.DbItems)
	for _, item := range items {
		betGroups[item.BetListId] = append(betGroups[item.BetListId], item)
	}

	// 更新每个组的赔率
	for betListId, group := range betGroups {
		var homeBetTotal, awayBetTotal float64

		// 计算主队和客队的投注总金额
		for _, item := range group {
			betAmount, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				fmt.Println("Invalid bet amount:", err)
				return
			}
			if item.BetItemId == "home" {
				homeBetTotal += betAmount
			} else if item.BetItemId == "away" {
				awayBetTotal += betAmount
			}
		}

		// 更新每个投注项的赔率
		for _, item := range group {
			betAmount, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				fmt.Println("Invalid bet amount:", err)
				continue
			}

			initOdds, err := strconv.ParseFloat(item.InitPrice, 64)
			if err != nil {
				fmt.Println("Invalid initial odds:", err)
				continue
			}

			var newOdds float64
			if item.BetItemId == "home" {
				newOdds = adjustOdds(initOdds, homeBetTotal, betAmount)
			} else if item.BetItemId == "away" {
				newOdds = adjustOdds(initOdds, awayBetTotal, betAmount)
			} else {
				continue
			}

			// 更新数据库中的赔率
			item.Odds = fmt.Sprintf("%.2f", newOdds)
			mysql.UpdateOdds(&item)

			fmt.Printf("Updated Odds: BetListId: %d, Item: %s, New Odds: %.2f\n", betListId, item.ItemName, newOdds)
		}
	}
}

func adjustOdds(currentOdds, totalBet, betAmount float64) float64 {
	// 调整算法：根据投注比例调整赔率
	if betAmount == 0 {
		return currentOdds * 2 // 如果某一方没有投注，将其赔率设为一个较高值
	}
	newOdds := currentOdds * (totalBet / (2 * betAmount))

	// 确保赔率不低于1.01
	if newOdds < 1.01 {
		newOdds = 1.01
	}
	return newOdds
}
