package service

import (
	"server/dao/mysql"
)

func GetGroupStandingsByGroup(groupId int64) ([]mysql.GroupStanding, error) {
	groupStandings, err := mysql.GetGroupStandingsByGroup(groupId)
	return groupStandings, err
}
