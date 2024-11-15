package mysql

import "fmt"

type Player struct {
	Nickname string `json:"nickname" gorm:"nickname"`
	Photo    string `json:"photo" gorm:"photo"`
}

func GetPlayersByTeams(teamId int64) ([]Player, error) {
	var players []Player
	err := db.Table("players").
		Select("nickname, photo").
		Where("teamId =?", teamId).
		Scan(&players).Error
	fmt.Println(players)
	return players, err
}
