package mysql

import (
	"fmt"
	"server/model"
)

func GetBetListById(id int64) (model.DbBetLists, error) {
	var betList model.DbBetLists
	if err := db.Where("list_id = ?", id).First(&betList).Error; err != nil {
		return betList, fmt.Errorf("failed to find bet list: %v", err)
	}
	return betList, nil
}

func GetBetListByIsSettled() ([]model.DbBetLists, error) {
	var betLists []model.DbBetLists
	if err := db.Find(&betLists).Error; err != nil {
		return betLists, fmt.Errorf("failed to find bet lists: %v", err)
	}
	return betLists, nil
}
