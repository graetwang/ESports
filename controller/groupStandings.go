package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"server/cache/redis"
	"server/service"
	"server/utils"
	"strconv"
	"time"
)

type GroupStanding struct {
	TeamName string `json:"teamName"`
	TeamLogo string `json:"teamLogo"`
	WonLost  string `json:"wonLost"`
	Points   int64  `json:"points"`
}

func GetGroupStandingsByGroup(c *gin.Context) {
	str := c.Query("group_id")
	groupId, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		c.JSON(200, utils.ECode{
			Code:    -1,
			Message: "参数错误",
			Data:    nil,
		})
		return
	}

	cacheKey := "groupStandings_" + fmt.Sprintf("%d", groupId)
	if !utils.Contains(redis.KeyList, cacheKey) {
		redis.KeyList = append(redis.KeyList, cacheKey)
	}

	// 尝试从缓存中获取数据
	groupStandings, err := getGroupStandingsFromCache(cacheKey)
	if err == nil {
		c.JSON(200, utils.ECode{
			Code:    0,
			Message: "获取组赛成绩成功",
			Data:    groupStandings,
		})
		return
	}

	// 从数据库获取数据
	data, err := service.GetGroupStandingsByGroup(groupId)
	if err != nil {
		c.JSON(500, utils.ECode{
			Code:    -1,
			Message: "获取组赛成绩失败",
			Data:    nil,
		})
		return
	}

	// 将数据转换为GroupStanding类型
	for _, groupStanding := range data {
		x := GroupStanding{
			TeamName: groupStanding.TeamName,
			TeamLogo: groupStanding.TeamLogo,
			WonLost:  fmt.Sprintf("%d/%d", groupStanding.MatchesWon, groupStanding.MatchesLost),
			Points:   groupStanding.Points,
		}
		groupStandings = append(groupStandings, x)
	}

	// 将数据存入缓存
	if err := redis.SetCache(cacheKey, groupStandings, 10*time.Minute); err != nil {
		c.JSON(200, utils.ECode{
			Code:    -1,
			Message: "缓存失败",
			Data:    nil,
		})
		return
	}

	c.JSON(200, utils.ECode{
		Code:    0,
		Message: "获取组赛成绩成功",
		Data:    groupStandings,
	})
}

func getGroupStandingsFromCache(cacheKey string) ([]GroupStanding, error) {
	var groupStandings []GroupStanding

	// 尝试从缓存中获取数据
	cachedData, err := redis.GetCache("old" + cacheKey)
	if err == nil {
		if err := json.Unmarshal(cachedData, &groupStandings); err == nil {
			fmt.Printf("从old缓存中获取到: %v\n", groupStandings)
			return groupStandings, nil
		}
	}

	cachedData, err = redis.GetCache(cacheKey)
	if err == nil {
		if err := json.Unmarshal(cachedData, &groupStandings); err == nil {
			fmt.Printf("从缓存中获取到: %v\n", groupStandings)
			return groupStandings, nil
		}
	}

	return nil, fmt.Errorf("缓存中没有找到数据")
}
