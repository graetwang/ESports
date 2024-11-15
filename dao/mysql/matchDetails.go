package mysql

import (
	"fmt"
	"time"
)

type MatchDetail struct {
	Event     string    `json:"event" gorm:"event"`
	Round     string    `json:"round" gorm:"round"`
	MatchTime time.Time `json:"matchTime" gorm:"matchTime"`
	Status    int64     `json:"status" gorm:"status"`
	HomeTeam  string    `json:"homeTeam" gorm:"homeTeam"`
	AwayTeam  string    `json:"awayTeam" gorm:"awayTeam"`
	HomeLogo  string    `json:"homeLogo" gorm:"homeLogo"`
	AwayLogo  string    `json:"awayLogo" gorm:"awayLogo"`
	HomeScore int64     `json:"homeScore" gorm:"homeScore"`
	AwayScore int64     `json:"awayScore" gorm:"awayScore"`
	Section   string    `json:"section" gorm:"section"`
	Phase     string    `json:"phase" gorm:"phase"`
	Group     string    `json:"group" gorm:"group"`
}

// GetMatchDetails 获取比赛详情
func GetMatchDetails() ([]MatchDetail, error) {
	var matchDetails []MatchDetail

	err := db.Debug().
		Table("schedules as s").
		Select("s.match_event AS event, m.phase AS round,s.project AS section, s.matchTime AS match_time, m.home_name AS home_team, m.home_logo, m.away_name AS away_team, m.away_logo,m.status, m.home_score, m.away_score").
		Joins("LEFT JOIN matches AS m ON m.id = s.matchId").
		Order("s.matchTime DESC").
		Scan(&matchDetails).Error

	if err != nil {
		return nil, err
	}

	fmt.Println(matchDetails)
	return matchDetails, nil
}

func GetMatchDetailsByProject(projectId int64) ([]MatchDetail, error) {
	var matchDetails []MatchDetail

	err := db.Debug().
		Table("schedules as s").
		Select("s.match_event AS event, m.phase AS round,s.project AS section, s.matchTime AS match_time, m.home_name AS home_team, m.home_logo, m.away_name AS away_team, m.away_logo,m.status, m.home_score, m.away_score").
		Joins("LEFT JOIN matches AS m ON m.id = s.matchId").
		Where("m.project_id = ?", projectId).
		Scan(&matchDetails).Error

	if err != nil {
		return nil, err
	}

	fmt.Println(matchDetails)
	return matchDetails, nil
}

func GetMatchDetailsByGroup(groupId int64) ([]MatchDetail, error) {
	var matchDetails []MatchDetail

	err := db.Debug().
		Table("schedules as s").
		Select("s.match_event AS event, m.phase AS round,s.project AS section, s.matchTime AS match_time, m.home_name AS home_team, m.home_logo, m.away_name AS away_team, m.away_logo,m.status, m.home_score, m.away_score,g.name AS `group`,g.phase").
		Joins("LEFT JOIN matches AS m ON m.id = s.matchId").
		Joins("LEFT JOIN groups AS g ON g.id = m.group_id").
		Where("m.group_id = ?", groupId).
		Scan(&matchDetails).Error

	if err != nil {
		return nil, err
	}

	fmt.Println(matchDetails)
	return matchDetails, nil
}
