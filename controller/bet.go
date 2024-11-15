package controller

import (
	"github.com/gin-gonic/gin"
	"server/service"
	"server/utils"
	"strconv"
)

func PlaceBet(c *gin.Context) {
	str1 := c.Query("user_id")
	userId, _ := strconv.ParseInt(str1, 10, 64)
	str2 := c.Query("bet_item_id")
	betItemId, _ := strconv.ParseInt(str2, 10, 64)
	str3 := c.Query("bet_amount")
	betAmount, _ := strconv.ParseInt(str3, 10, 64)
	err := service.PlaceBet(userId, betItemId, float64(betAmount))
	if err != nil {
		c.JSON(200, utils.ECode{
			Code:    -1,
			Message: "下注失败",
			Data:    err,
		})
		return
	}
	c.JSON(200, utils.ECode{
		Code:    0,
		Message: "下注成功",
		Data:    nil,
	})
}
