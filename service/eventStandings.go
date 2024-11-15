package service

import (
	"server/dao/mysql"
)

type teamStanding struct {
	mysql.EventStanding
	TeamMembers []mysql.Player `json:"team_members"`
}

func EventStandings(eventId int64) ([]teamStanding, error) {
	var teamStandings []teamStanding
	standings, err := mysql.GetTeamStandings(eventId)
	if err != nil {
		return nil, err
	}
	for _, standing := range standings {
		teamMembers, err := mysql.GetPlayersByTeams(standing.TeamId)
		if err != nil {
			return nil, err
		}
		teamStandings = append(teamStandings, teamStanding{
			EventStanding: standing,
			TeamMembers:   teamMembers,
		})
	}
	return teamStandings, nil
}
