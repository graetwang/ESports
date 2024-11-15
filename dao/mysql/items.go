package mysql

import (
	"fmt"
	"server/model"
)

func ItemList() ([]model.DbItems, error) {
	var items []model.DbItems
	if err := db.Find(&items).Error; err != nil {
		fmt.Println("Failed to fetch bets:", err)
		return nil, err
	}
	return items, nil
}

func UpdateOdds(item *model.DbItems) error {
	if err := db.Save(&item).Error; err != nil {
		fmt.Println("Failed to update odds:", err)
		return err
	}
	return nil
}

func GetItemById(id int64) (model.DbItems, error) {
	var item model.DbItems
	if err := db.First(&item, id).Error; err != nil {
		return item, fmt.Errorf("failed to find bet item: %v", err)
	}
	return item, nil
}

func UpdateBetItem(item model.DbItems) error {
	if err := db.Save(item).Error; err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}
	return nil
}
