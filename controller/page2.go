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

type page struct {
	PhaseName      string
	Group          string
	MatchDetails   []mysql.MatchDetail
	GroupStandings []mysql.GroupStanding
}

func Page2(c *gin.Context) {
	str := c.Query("event_id")
	eventId, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		c.JSON(400, utils.ECode{
			Code:    -1,
			Message: "参数错误",
			Data:    nil,
		})
		return
	}

	cacheKey := "pagesByEventId_" + fmt.Sprintf("%d", eventId)
	if !utils.Contains(redis.KeyList, cacheKey) {
		redis.KeyList = append(redis.KeyList, cacheKey)
	}

	// 尝试从缓存中获取数据
	pages, err := getPageFromCache(cacheKey)
	if err == nil {
		c.JSON(200, utils.ECode{
			Code:    0,
			Message: "获取页面二成功",
			Data:    pages,
		})
		return
	}

	// 从数据库获取数据
	data, err := service.GetPageByGroup(eventId)
	if err != nil {
		c.JSON(500, utils.ECode{
			Code:    -1,
			Message: "获取页面二失败",
			Data:    data,
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
		Message: "获取页面二成功",
		Data:    data,
	})
}

func getPageFromCache(cacheKey string) ([]page, error) {
	var pages []page

	// 尝试从缓存中获取数据
	cachedData, err := redis.GetCache("old" + cacheKey)
	if err == nil {
		if err := json.Unmarshal(cachedData, &pages); err == nil {
			fmt.Printf("从old缓存中获取到: %v\n", pages)
			return pages, nil
		}
	}

	cachedData, err = redis.GetCache(cacheKey)
	if err == nil {
		if err := json.Unmarshal(cachedData, &pages); err == nil {
			fmt.Printf("从缓存中获取到: %v\n", pages)
			return pages, nil
		}
	}

	return nil, fmt.Errorf("缓存中没有找到数据")
}
