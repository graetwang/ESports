package mysql

type EventStanding struct {
	TeamId                 int64   `json:"team_id"`
	Team                   string  `json:"team"`
	TeamLogo               string  `json:"team_logo"`
	MatchesWon             int     `json:"matches_won"`
	MatchesLost            int     `json:"matches_lost"`
	DamagePerMinute        float32 `json:"dpm"`
	GoldPerMinute          float32 `json:"gpm"`
	CreepScorePerMinute    float32 `json:"cspm"`
	SmallDragonControlRate float32 `json:"sdcr"`
	AverageKillsPerGame    float32 `json:"akpm"`
	AverageAssistsPerGame  float32 `json:"apm"`
	AverageDeathsPerGame   float32 `json:"adpm"`
	BaronControlRate       float32 `json:"bcr"`
}

func GetTeamStandings(eventId int64) ([]EventStanding, error) {
	var teamStandings []EventStanding
	err := db.Table("event_standings").
		Where("event_id =?", eventId).
		Order("id DESC").
		Scan(&teamStandings).Error
	return teamStandings, err
}
