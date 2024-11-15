package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"server/cache/redis"
	"server/dao/mysql"
	"server/service"
	"server/utils"
	"strconv"
	"time"
)

type teamStanding struct {
	mysql.EventStanding
	TeamMembers []mysql.Player `json:"team_members"`
}

func EventStandings(c *gin.Context) {
	str := c.Query("event_id")
	eventId, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		c.JSON(200, utils.ECode{
			Code:    -1,
			Message: "参数错误",
			Data:    nil,
		})
		return
	}

	cacheKey := "teamStandingsByEventId_" + fmt.Sprintf("%d", eventId)
	if !utils.Contains(redis.KeyList, cacheKey) {
		redis.KeyList = append(redis.KeyList, cacheKey)
	}

	// 尝试从缓存中获取数据
	teamStandings, err := getTeamStandingsFromCache(cacheKey)
	if err == nil {
		c.JSON(200, utils.ECode{
			Code:    0,
			Message: "获取赛事积分数据成功",
			Data:    teamStandings,
		})
		return
	}

	// 从数据库获取数据
	data, err := service.EventStandings(eventId)
	if err != nil {
		c.JSON(500, utils.ECode{
			Code:    -1,
			Message: "获取赛事积分数据失败",
			Data:    err,
		})
		return
	}

	// 将数据存入缓存
	if err := redis.SetCache(cacheKey, data, 10*time.Minute); err != nil {
		c.JSON(200, utils.ECode{
			Code:    -1,
			Message: "缓存失败",
			Data:    nil,
		})
		return
	}

	c.JSON(200, utils.ECode{
		Code:    0,
		Message: "获取赛事积分数据成功",
		Data:    data,
	})
}

func getTeamStandingsFromCache(cacheKey string) ([]teamStanding, error) {
	var teamStandings []teamStanding

	// 尝试从缓存中获取数据
	cachedData, err := redis.GetCache("old" + cacheKey)
	if err == nil {
		if err := json.Unmarshal(cachedData, &teamStandings); err == nil {
			fmt.Printf("从old缓存中获取到: %v\n", teamStandings)
			return teamStandings, nil
		}
	}

	cachedData, err = redis.GetCache(cacheKey)
	if err == nil {
		if err := json.Unmarshal(cachedData, &teamStandings); err == nil {
			fmt.Printf("从缓存中获取到: %v\n", teamStandings)
			return teamStandings, nil
		}
	}

	return nil, fmt.Errorf("缓存中没有找到数据")
}
