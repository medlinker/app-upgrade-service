package main

import (
	"flag"
	"net/http"
	"strings"

	"app-upgrade-service/config"
	. "app-upgrade-service/logger"
	"app-upgrade-service/models"

	"github.com/gin-gonic/gin"
)

func init() {
	confFile := flag.String("env", "./conf/env.conf", "global configuration!")
	flag.Parse()
	config.InitConf(*confFile)
	InitLog()
	models.ListenConfiguration()
}

func main() {
	router := gin.Default()

	// use middleware
	router.Use(gin.Recovery())

	// set mode
	if config.GetString("mode", false, "prod") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	router.GET("/version", upgrade)
	router.POST("/version", upgrade)

	router.Run(config.GetString("hostAndPort", true, ":8866"))
}

func upgrade(c *gin.Context) {
	req := models.RequestParam{}
	reqMethod := c.Request.Method
	if reqMethod == "GET" {
		req.Platform = strings.ToLower(c.Query("platform"))
		req.ClientVersion = strings.ToLower(c.Query("clientVersion"))
		req.ClientChannel = c.Query("clientChannel")
	} else if reqMethod == "POST" {
		req.Platform = strings.ToLower(c.PostForm("platform"))
		req.ClientVersion = strings.ToLower(c.PostForm("clientVersion"))
		req.ClientChannel = c.PostForm("clientChannel")
	}
	if req.Illegal() {
		Logger.Error(req)
		c.JSON(http.StatusOK, "request params illegal!")
		return
	}
	if platform, ok := models.CurretPlatforms[req.Platform]; ok {
		resp := models.Response{}
		url := strings.Replace(platform.URL, "${channel}", req.ClientChannel, -1)
		data := models.Data{Title: platform.Title, Desc: platform.Desc, URL: url}
		resp.Data = data
		resp.Platform = req.Platform
		resp.ClientVersion = req.ClientVersion
		resp.ClientChannel = req.ClientChannel
		resp.CurrentVersion = platform.Version
		if strings.TrimSpace(platform.Version) == req.ClientVersion {
			resp.Data.Action = models.NEWEST
		} else {
			resp.Data.Action = generateResponse(req)
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	c.JSON(http.StatusOK, "unsupport platform")
}

func generateResponse(req models.RequestParam) int {
	for _, rule := range (*models.VersionConfig).Rules {
		if strings.ToLower(rule.Platform) == req.Platform {
			if strings.Compare(rule.MinVersion, req.ClientVersion) <= 0 &&
				strings.Compare(rule.MaxVersion, req.ClientVersion) >= 0 {
				return rule.Action
			}
		}
	}
	return -1
}
