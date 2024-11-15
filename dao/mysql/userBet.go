package mysql

import (
	"fmt"
	"server/model"
)

func CreateBetRecord(bet *model.UserBet) error {
	if err := db.Create(&bet).Error; err != nil {
		return fmt.Errorf("failed to place bet: %v", err)
	}
	return nil
}

func GetUnSettledRecordsByBetListId(id int64) ([]model.UserBet, error) {
	var bets []model.UserBet
	if err := db.Joins("JOIN db_items ON user_bet.bet_item_id = db_items.id").
		Where("db_items.bet_list_id = ? AND user_bet.is_settled = ?", id, 0).
		Find(&bets).Error; err != nil {
		return bets, fmt.Errorf("failed to find bets: %v", err)
	}
	return bets, nil
}

func UpdateBetRecord(bet model.UserBet) error {
	if err := db.Save(bet).Error; err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}
	return nil
}
