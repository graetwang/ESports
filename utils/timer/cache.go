package timer

import (
	"fmt"
	"server/cache/redis"
	"server/service"
	"strconv"
	"strings"
	"time"
)

// SyncCache 定时器同步函数
func SyncCache(interval time.Duration) {
	for {
		fmt.Println(redis.KeyList)
		for _, key := range redis.KeyList {
			if key == "bannerList" {
				data, _ := service.Banner()
				redis.SetCache(key, data, 0)
			} else if strings.Contains(key, "teamStandingsByEventId") {
				split := strings.Split(key, "_")
				eventId, _ := strconv.ParseInt(split[1], 10, 64)
				data, _ := service.EventStandings(eventId)
				redis.SetCache(key, data, 0)
			} else if strings.Contains(key, "groupStandings") {
				split := strings.Split(key, "_")
				groupId, _ := strconv.ParseInt(split[1], 10, 64)
				data, _ := service.GetGroupStandingsByGroup(groupId)
				redis.SetCache(key, data, 0)
			} else if strings.Contains(key, "matchDetailsAll") {
				data, _ := service.GetMatchDetails()
				redis.SetCache(key, data, 0)
			} else if strings.Contains(key, "matchDetailsByProject") {
				split := strings.Split(key, "_")
				projectId, _ := strconv.ParseInt(split[1], 10, 64)
				data, _ := service.GetMatchDetailsByProject(projectId)
				redis.SetCache(key, data, 0)
			} else if strings.Contains(key, "matchDetailsByGroup_") {
				split := strings.Split(key, "_")
				groupId, _ := strconv.ParseInt(split[1], 10, 64)
				data, _ := service.GetMatchDetailsByGroup(groupId)
				redis.SetCache(key, data, 0)
			} else if strings.Contains(key, "matchDetailsByEventAndPhase_") {
				split := strings.Split(key, "_")
				data, _ := service.GetMatchesByEventAndPhase(split[1], split[2])
				redis.SetCache(key, data, 0)
			} else if strings.Contains(key, "pagesByEventId_") {
				split := strings.Split(key, "_")
				eventId, _ := strconv.ParseInt(split[1], 10, 64)
				data, _ := service.GetPageByGroup(eventId)
				redis.SetCache(key, data, 0)
			} else if strings.Contains(key, "sections") {
				data, _ := service.GetTodaySection()
				redis.SetCache(key, data, 0)
			} else if strings.Contains(key, "videosBySection_") {
				split := strings.Split(key, "_")
				data, _ := service.GetVideosBySection(split[1])
				redis.SetCache(key, data, 0)
			} else if strings.Contains(key, "videosBySectionLimit_") {
				split := strings.Split(key, "_")
				page, _ := strconv.Atoi(split[2])
				pageSize, _ := strconv.Atoi(split[3])
				data, _ := service.GetVideosBySectionLimit(split[1], pageSize, page)
				redis.SetCache(key, data, 0)
			}
		}
		fmt.Println("缓存已更新")
		time.Sleep(interval)
	}
}
