package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"server/cache/redis"
	"server/service"
	"server/utils"
	"time"
)

func Sections(c *gin.Context) {
	cacheKey := "sections"
	if !utils.Contains(redis.KeyList, cacheKey) {
		redis.KeyList = append(redis.KeyList, cacheKey)
	}

	// 尝试从缓存中获取数据
	sections, err := getSectionsFromCache(cacheKey)
	if err == nil {
		c.JSON(200, utils.ECode{
			Code:    0,
			Message: "获取分区信息成功",
			Data:    sections,
		})
		return
	}

	// 从数据库获取数据
	data, err := service.GetTodaySection()
	if err != nil {
		c.JSON(500, utils.ECode{
			Code:    -1,
			Message: "获取分区信息失败",
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
		Message: "获取分区信息成功",
		Data:    data,
	})
}

func getSectionsFromCache(cacheKey string) ([]service.Section, error) {
	var sections []service.Section

	// 尝试从缓存中获取数据
	cachedData, err := redis.GetCache("old" + cacheKey)
	if err == nil {
		if err := json.Unmarshal(cachedData, &sections); err == nil {
			fmt.Printf("从old缓存中获取到: %v\n", sections)
			return sections, nil
		}
	}

	cachedData, err = redis.GetCache(cacheKey)
	if err == nil {
		if err := json.Unmarshal(cachedData, &sections); err == nil {
			fmt.Printf("从缓存中获取到: %v\n", sections)
			return sections, nil
		}
	}

	return nil, fmt.Errorf("缓存中没有找到数据")
}
