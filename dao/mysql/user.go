package mysql

import (
	"fmt"
	"server/model"
)

func GetUserById(id int64) (model.User, error) {
	var user model.User
	if err := db.First(&user, id).Error; err != nil {
		return user, fmt.Errorf("failed to find user: %v", err)
	}
	return user, nil
}

func UpdateUser(user *model.User) error {
	if err := db.Save(user).Error; err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}
	return nil
}
