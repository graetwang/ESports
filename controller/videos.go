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

func Videos(c *gin.Context) {
	section := c.Query("section")
	pageSize := c.Query("pageSize")
	page := c.Query("page")

	if pageSize == "" || page == "" {
		cacheKey := "videosBySection_" + section
		if !utils.Contains(redis.KeyList, cacheKey) {
			redis.KeyList = append(redis.KeyList, cacheKey)
		}

		// 尝试从缓存中获取数据
		videos, err := getVideosFromCache(cacheKey)
		if err == nil {
			c.JSON(200, utils.ECode{
				Code:    0,
				Message: "获取节目列表成功",
				Data:    videos,
			})
			return
		}

		// 从数据库获取数据
		data, err := service.GetVideosBySection(section)
		if err != nil {
			c.JSON(500, utils.ECode{
				Code:    -1,
				Message: "获取节目列表失败",
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
			Message: "获取节目列表成功",
			Data:    data,
		})
		return
	} else {
		pageSizeInt, _ := strconv.Atoi(pageSize)
		pageInt, _ := strconv.Atoi(page)
		cacheKey := "videosBySectionLimit_" + section + fmt.Sprintf("_%d_%d", pageInt, pageSizeInt)
		if !utils.Contains(redis.KeyList, cacheKey) {
			redis.KeyList = append(redis.KeyList, cacheKey)
		}

		// 尝试从缓存中获取数据
		videos, err := getVideosFromCache(cacheKey)
		if err == nil {
			c.JSON(200, utils.ECode{
				Code:    0,
				Message: "获取节目列表成功",
				Data:    videos,
			})
			return
		}

		// 从数据库获取数据
		data, err := service.GetVideosBySectionLimit(section, pageSizeInt, pageInt)
		if err != nil {
			c.JSON(500, utils.ECode{
				Code:    -1,
				Message: "获取节目列表失败",
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
			Message: "获取节目列表成功",
			Data:    data,
		})
		return
	}
}

func getVideosFromCache(cacheKey string) ([]mysql.Video, error) {
	var videos []mysql.Video

	// 尝试从缓存中获取数据
	cachedData, err := redis.GetCache("old" + cacheKey)
	if err == nil {
		if err := json.Unmarshal(cachedData, &videos); err == nil {
			fmt.Printf("从old缓存中获取到: %v\n", videos)
			return videos, nil
		}
	}

	cachedData, err = redis.GetCache(cacheKey)
	if err == nil {
		if err := json.Unmarshal(cachedData, &videos); err == nil {
			fmt.Printf("从缓存中获取到: %v\n", videos)
			return videos, nil
		}
	}

	return nil, fmt.Errorf("缓存中没有找到数据")
}
