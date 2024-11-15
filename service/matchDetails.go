package service

import (
	"server/dao/mysql"
)

func GetMatchDetails() ([]mysql.MatchDetail, error) {
	matchDetails, err := mysql.GetMatchDetails()
	return matchDetails, err
}

func GetMatchDetailsByProject(projectId int64) ([]mysql.MatchDetail, error) {
	matchDetails, err := mysql.GetMatchDetailsByProject(projectId)
	return matchDetails, err
}

func GetMatchDetailsByGroup(groupId int64) ([]mysql.MatchDetail, error) {
	matchDetails, err := mysql.GetMatchDetailsByGroup(groupId)
	return matchDetails, err
}
