package router

import (
	"github.com/gin-gonic/gin"
	"server/controller"
	"server/utils/timer"
	"time"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.GET("/matchDetails", controller.GetMatchDetails)
	r.GET("/matchDetailsByProject", controller.GetMatchDetailsByProject)
	r.GET("/banner", controller.Banner)
	r.GET("/section", controller.Sections)
	r.GET("/groupStandingsByGroup", controller.GetGroupStandingsByGroup)
	r.GET("/matchDetailsByGroup", controller.GetMatchDetailsByGroup)
	r.GET("/page2", controller.Page2)
	r.GET("/eventStandings", controller.EventStandings)
	r.GET("/videos", controller.Videos)
	r.GET("/2024LPLSummer", controller.MatchesLPLSummer2024)
	r.POST("/placeBet", controller.PlaceBet)
	go timer.SyncCache(time.Second * 5)
	go timer.OddsUpdate()
	go timer.SettleBets()
	return r
}
