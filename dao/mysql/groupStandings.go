package mysql

type GroupStanding struct {
	TeamName    string `json:"teamName"`
	TeamLogo    string `json:"teamLogo"`
	MatchesWon  int64  `json:"matchesWon"`
	MatchesLost int64  `json:"matchesLost"`
	Points      int64  `json:"points"`
	Group       string `json:"group"`
	Phase       string `json:"phase"`
}

func GetGroupStandingsByGroup(groupId int64) ([]GroupStanding, error) {
	var groupStandings []GroupStanding

	err := db.Debug().
		Table("group_standings AS gs").
		Select("gs.team AS team_name, gs.team_logo AS team_logo, gs.matches_won, gs.matches_lost, gs.points,g.name AS `group`,g.phase").
		Joins("LEFT JOIN `groups` AS g ON g.id = gs.group_id").
		Where("gs.group_id =?", groupId).
		Scan(&groupStandings).Error

	return groupStandings, err
}
