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

func GetMatchDetails(c *gin.Context) {
	cacheKey := "matchDetailsAll"
	if !utils.Contains(redis.KeyList, cacheKey) {
		redis.KeyList = append(redis.KeyList, cacheKey)
	}

	// 尝试从缓存中获取数据
	matchDetails, err := getMatchDetailsFromCache(cacheKey)
	if err == nil {
		c.JSON(200, utils.ECode{
			Code:    0,
			Message: "获取赛程详情成功",
			Data:    matchDetails,
		})
		return
	}

	// 从数据库获取数据
	data, err := service.GetMatchDetails()
	if err != nil {
		c.JSON(500, utils.ECode{
			Code:    -1,
			Message: "获取赛程详情失败",
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
		Message: "获取赛程详情成功",
		Data:    data,
	})
}

func GetMatchDetailsByProject(c *gin.Context) {
	str := c.Query("project_id")
	projectId, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		c.JSON(200, utils.ECode{
			Code:    -1,
			Message: "参数错误",
			Data:    nil,
		})
		return
	}

	cacheKey := "matchDetailsByProject_" + fmt.Sprintf("%d", projectId)
	if !utils.Contains(redis.KeyList, cacheKey) {
		redis.KeyList = append(redis.KeyList, cacheKey)
	}

	// 尝试从缓存中获取数据
	matchDetails, err := getMatchDetailsFromCache(cacheKey)
	if err == nil {
		c.JSON(200, utils.ECode{
			Code:    0,
			Message: "获取赛程详情成功",
			Data:    matchDetails,
		})
		return
	}

	// 从数据库获取数据
	data, err := service.GetMatchDetailsByProject(projectId)
	if err != nil {
		c.JSON(500, utils.ECode{
			Code:    -1,
			Message: "获取赛程详情失败",
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
		Message: "获取赛程详情成功",
		Data:    data,
	})
}

func GetMatchDetailsByGroup(c *gin.Context) {
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

	cacheKey := "matchDetailsByGroup_" + fmt.Sprintf("%d", groupId)
	if !utils.Contains(redis.KeyList, cacheKey) {
		redis.KeyList = append(redis.KeyList, cacheKey)
	}

	// 尝试从缓存中获取数据
	matchDetails, err := getMatchDetailsFromCache(cacheKey)
	if err == nil {
		c.JSON(200, utils.ECode{
			Code:    0,
			Message: "获取赛程详情成功",
			Data:    matchDetails,
		})
		return
	}

	// 从数据库获取数据
	data, err := service.GetMatchDetailsByGroup(groupId)
	if err != nil {
		c.JSON(500, utils.ECode{
			Code:    -1,
			Message: "获取赛程详情失败",
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
		Message: "获取赛程详情成功",
		Data:    data,
	})
}

func getMatchDetailsFromCache(cacheKey string) ([]mysql.MatchDetail, error) {
	var matchDetails []mysql.MatchDetail

	// 尝试从缓存中获取数据
	cachedData, err := redis.GetCache("old" + cacheKey)
	if err == nil {
		if err := json.Unmarshal(cachedData, &matchDetails); err == nil {
			fmt.Printf("从old缓存中获取到: %v\n", matchDetails)
			return matchDetails, nil
		}
	}

	cachedData, err = redis.GetCache(cacheKey)
	if err == nil {
		if err := json.Unmarshal(cachedData, &matchDetails); err == nil {
			fmt.Printf("从缓存中获取到: %v\n", matchDetails)
			return matchDetails, nil
		}
	}

	return nil, fmt.Errorf("缓存中没有找到数据")
}
