package service

import (
	"fmt"
	"server/dao/mysql"
)

type page struct {
	PhaseName      string
	Group          string
	MatchDetails   []mysql.MatchDetail
	GroupStandings []mysql.GroupStanding
}

func GetPageByGroup(eventId int64) ([]page, error) {
	var pages []page
	phases, err := mysql.GetPhasesByEvent(eventId)
	if err != nil {
		return nil, err
	}
	fmt.Println("phases", phases)
	for _, phase := range phases {
		var page page
		page.PhaseName = phase.Name
		groups, err := mysql.GetGroupsByPhase(phase.Id)
		if err != nil {
			return nil, err
		}
		fmt.Println("groups", groups)
		for _, group := range groups {
			matchDetails, err := mysql.GetMatchDetailsByGroup(group.Id)
			if err != nil {
				return nil, err
			}
			fmt.Println("matchDetails", matchDetails)
			groupStandings, err := mysql.GetGroupStandingsByGroup(group.Id)
			if err != nil {
				return nil, err
			}
			fmt.Println("groupStandings", groupStandings)
			fmt.Println("page", page)
			page.Group = group.Name
			page.GroupStandings = groupStandings
			page.MatchDetails = matchDetails
			pages = append(pages, page)
		}
	}
	return pages, nil
}
