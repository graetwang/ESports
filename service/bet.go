package service

import (
	"fmt"
	"server/dao/mysql"
	"server/model"
	"strconv"
)

func PlaceBet(userID int64, betItemID int64, amount float64) error {
	item, err := mysql.GetItemById(betItemID)
	if err != nil {
		return err
	}
	// 检查下注金额是否符合单项下限
	maxAmount, err := strconv.ParseFloat(item.MemberMaxBet, 64)
	if err != nil {
		return err
	}
	if amount > maxAmount {
		return fmt.Errorf("下注金额超出单项上限")
	}
	// 检查用户余额是否足够
	user, err := mysql.GetUserById(userID)
	if err != nil {
		return err
	}

	if user.LiuCoin < amount {
		return fmt.Errorf("您的余额不足")
	}

	// 创建下注记录
	bet := model.UserBet{
		UserId:    userID,
		BetItemId: betItemID,
		Amount:    amount,
	}

	err = mysql.CreateBetRecord(&bet)

	// 扣除用户余额
	user.LiuCoin -= amount
	price, err := strconv.ParseFloat(item.Price, 64)
	price += amount
	item.Price = fmt.Sprintf("%f", price)
	err = mysql.UpdateUser(&user)
	err = mysql.UpdateBetItem(item)

	return nil
}

func SettleBets(listID int64, winningBetItemID int64) error {
	// 获取所有相关的投注记录
	bets, err := mysql.GetUnSettledRecordsByBetListId(listID)
	if err != nil {
		return err
	}

	// 处理每个下注
	for _, bet := range bets {
		user, err := mysql.GetUserById(bet.UserId)
		if err != nil {
			return err
		}
		// 获取投注项
		item, err := mysql.GetItemById(bet.BetItemId)

		// 计算赔率
		odds, err := strconv.ParseFloat(item.Odds, 64)
		if err != nil {
			return fmt.Errorf("invalid odds: %v", err)
		}

		// 如果用户的投注项是赢的
		if bet.BetItemId == winningBetItemID {
			// 增加用户余额
			winAmount := bet.Amount * odds
			user.LiuCoin += winAmount
			mysql.UpdateUser(&user)
		}
		bet.IsSettled = 1
		mysql.UpdateBetRecord(bet)
	}
	return nil
}

func CheckAndSettleBets() error {
	betLists, err := mysql.GetBetListByIsSettled()
	if err != nil {
		return err
	}
	// 结算每个满足条件的投注
	for _, betList := range betLists {
		var winningBetItemID int64
		if betList.TeamAWin == "1" {
			winningBetItemID = betList.ID // 根据业务逻辑确定winningBetItemID
		} else if betList.TeamBWin == "1" {
			winningBetItemID = betList.ID // 根据业务逻辑确定winningBetItemID
		} else {
			continue
		}

		if err := SettleBets(betList.ListId, winningBetItemID); err != nil {
			return fmt.Errorf("failed to settle bets for list ID %d: %v", betList.ListId, err)
		}
	}
	return nil
}
