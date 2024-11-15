package service

import (
	"server/dao/mysql"
)

func GetMatchesByEventAndPhase(event, phase string) (map[string][]mysql.Match, error) {
	matches, err := mysql.GetMatchesByEventAndPhase(event, phase)
	groupedByRound := make(map[string][]mysql.Match)
	for _, match := range matches {
		groupedByRound[match.Round] = append(groupedByRound[match.Round], match)
	}
	return groupedByRound, err
}
