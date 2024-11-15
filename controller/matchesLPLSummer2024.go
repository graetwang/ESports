package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"server/cache/redis"
	"server/dao/mysql"
	"server/service"
	"server/utils"
	"time"
)

func MatchesLPLSummer2024(c *gin.Context) {
	phase := c.Query("phase")
	event := "2024LPLSummer"
	cacheKey := "matchDetailsByEventAndPhase_" + event + "_" + phase
	if !utils.Contains(redis.KeyList, cacheKey) {
		redis.KeyList = append(redis.KeyList, cacheKey)
	}

	// 尝试从缓存中获取数据
	groupedByRound, err := getMatchDetailsFromCache(cacheKey)
	if err == nil {
		c.JSON(200, utils.ECode{
			Code:    0,
			Message: "获取对阵图信息成功",
			Data:    groupedByRound,
		})
		return
	}

	// 从数据库获取数据
	data, err := service.GetMatchesByEventAndPhase(event, phase)
	if err != nil {
		c.JSON(500, utils.ECode{
			Code:    -1,
			Message: "获取对阵图信息失败",
			Data:    nil,
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
		Message: "获取对阵图信息成功",
		Data:    data,
	})
}

func get2024SummerMatchDetailsFromCache(cacheKey string) (map[string][]mysql.Match, error) {
	var groupedByRound map[string][]mysql.Match

	// 尝试从缓存中获取数据
	cachedData, err := redis.GetCache("old" + cacheKey)
	if err == nil {
		if err := json.Unmarshal(cachedData, &groupedByRound); err == nil {
			fmt.Printf("从old缓存中获取到: %v\n", groupedByRound)
			return groupedByRound, nil
		}
	}

	cachedData, err = redis.GetCache(cacheKey)
	if err == nil {
		if err := json.Unmarshal(cachedData, &groupedByRound); err == nil {
			fmt.Printf("从缓存中获取到: %v\n", groupedByRound)
			return groupedByRound, nil
		}
	}

	return nil, fmt.Errorf("缓存中没有找到数据")
}
