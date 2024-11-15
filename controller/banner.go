package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"server/cache/redis"
	"server/dao/mysql"
	"server/service"
	"server/utils"
)

func Banner(c *gin.Context) {
	cacheKey := "bannerList"
	if !utils.Contains(redis.KeyList, cacheKey) {
		redis.KeyList = append(redis.KeyList, cacheKey)
	}

	// 尝试从缓存中获取数据
	banners, err := getBannersFromCache(cacheKey)
	if err == nil {
		c.JSON(200, utils.ECode{
			Code:    0,
			Message: "轮播图获取成功",
			Data:    banners,
		})
		return
	}

	// 从数据库获取数据
	data, err := service.Banner()
	if err != nil {
		c.JSON(500, utils.ECode{
			Code:    -1,
			Message: "轮播图获取失败",
			Data:    err,
		})
		return
	}

	// 将数据存入缓存
	if err := redis.SetCache(cacheKey, data, 0); err != nil {
		c.JSON(200, utils.ECode{
			Code:    -1,
			Message: "缓存失败",
			Data:    nil,
		})
		return
	}

	c.JSON(200, utils.ECode{
		Code:    0,
		Message: "轮播图获取成功",
		Data:    data,
	})
}

func getBannersFromCache(cacheKey string) ([]mysql.Banner, error) {
	var banners []mysql.Banner

	// 尝试从缓存中获取数据
	cachedData, err := redis.GetCache("old" + cacheKey)
	if err == nil {
		if err := json.Unmarshal(cachedData, &banners); err == nil {
			fmt.Printf("从old缓存中获取到: %v\n", banners)
			return banners, nil
		}
	}

	cachedData, err = redis.GetCache(cacheKey)
	if err == nil {
		if err := json.Unmarshal(cachedData, &banners); err == nil {
			fmt.Printf("从缓存中获取到: %v\n", banners)
			return banners, nil
		}
	}

	return nil, fmt.Errorf("缓存中没有找到数据")
}
