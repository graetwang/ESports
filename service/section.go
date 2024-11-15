package service

import (
	"server/dao/mysql"
)

type Section struct {
	MatchDetails []mysql.MatchDetail `json:"matchDetails"`
	Videos       []mysql.Video       `json:"videos"`
	Background   mysql.Background    `json:"background"`
}

func GetTodaySection() ([]Section, error) {
	matchDetails, err := mysql.GetMatchDetails()
	if err != nil {
		return nil, err
	}
	var (
		LOLmatchDetails        []mysql.MatchDetail
		wangzhematchDetails    []mysql.MatchDetail
		diwurengematchDetails  []mysql.MatchDetail
		wuweiqiyuematchDetails []mysql.MatchDetail
		DOTAmatchDetails       []mysql.MatchDetail
	)
	for i, detail := range matchDetails {
		if detail.Section == "英雄联盟" {
			if detail.Status != 1 && len(LOLmatchDetails) <= 2 {
				LOLmatchDetails = append(LOLmatchDetails, matchDetails[i])
			}
		} else if detail.Section == "王者荣耀" {
			if detail.Status != 1 && len(wangzhematchDetails) <= 2 {
				wangzhematchDetails = append(wangzhematchDetails, matchDetails[i])
			}
		} else if detail.Section == "第五人格" {
			if detail.Status != 1 && len(diwurengematchDetails) <= 2 {
				diwurengematchDetails = append(diwurengematchDetails, matchDetails[i])
			} else if detail.Section == "无畏契约" {
				if detail.Status != 1 && len(wuweiqiyuematchDetails) <= 2 {
					wuweiqiyuematchDetails = append(wuweiqiyuematchDetails, matchDetails[i])
				} else {
					if detail.Status != 1 && len(DOTAmatchDetails) <= 2 {
						DOTAmatchDetails = append(DOTAmatchDetails, matchDetails[i])
					}
				}
			}
		}
	}
	videos, err := mysql.GetVideoList()
	if err != nil {
		return nil, err
	}
	var (
		LOLvideos        []mysql.Video
		wangzhevideos    []mysql.Video
		diwurengevideos  []mysql.Video
		wuweiqiyuevideos []mysql.Video
		DOTAvideos       []mysql.Video
	)
	for _, video := range videos {
		if video.Section == "英雄联盟" {
			if len(LOLvideos) <= 6 {
				LOLvideos = append(LOLvideos, video)
			}
		} else if video.Section == "王者荣耀" {
			if len(wangzhevideos) <= 6 {
				wangzhevideos = append(wangzhevideos, video)
			}
		} else if video.Section == "第五人格" {
			if len(diwurengevideos) <= 6 {
				diwurengevideos = append(diwurengevideos, video)
			}
		} else if video.Section == "无畏契约" {
			if len(wuweiqiyuevideos) <= 6 {
				wuweiqiyuevideos = append(wuweiqiyuevideos, video)
			}
		} else {
			if len(DOTAvideos) <= 6 {
				DOTAvideos = append(DOTAvideos, video)
			}
		}
	}
	backgrounds, err := mysql.BackgroundList()
	if err != nil {
		return nil, err
	}
	LOLsection := Section{
		MatchDetails: LOLmatchDetails,
		Videos:       LOLvideos,
	}
	wangzhesection := Section{
		MatchDetails: wangzhematchDetails,
		Videos:       wangzhevideos,
	}
	diwurengesection := Section{
		MatchDetails: diwurengematchDetails,
		Videos:       diwurengevideos,
	}
	wuweiqiyuesection := Section{
		MatchDetails: wuweiqiyuematchDetails,
		Videos:       wuweiqiyuevideos,
	}
	DOTAsection := Section{
		MatchDetails: DOTAmatchDetails,
		Videos:       DOTAvideos,
	}
	for _, background := range backgrounds {
		LOLsection.Background = background
		wangzhesection.Background = background
		diwurengesection.Background = background
		wuweiqiyuesection.Background = background
		DOTAsection.Background = background
	}
	sections := []Section{LOLsection, wangzhesection, diwurengesection, wuweiqiyuesection, DOTAsection}
	return sections, nil
}
