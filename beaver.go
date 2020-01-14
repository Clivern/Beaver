// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"github.com/clivern/beaver/internal/app/cmd"
	"github.com/clivern/beaver/internal/app/controller"
	"github.com/clivern/beaver/internal/app/middleware"
	"github.com/clivern/beaver/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {

	var exec string
	var configFile string

	utils.PrintBanner()

	flag.StringVar(&exec, "exec", "", "exec")
	flag.StringVar(&configFile, "config", "config.dist.yml", "config")
	flag.Parse()

	viper.SetConfigFile(configFile)

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Sprintf(
			"Error while loading config file [%s]: %s",
			configFile,
			err.Error(),
		))
	}

	if exec != "" {
		switch exec {
		case "health":
			cmd.HealthStatus()
		default:
			utils.PrintCommands()
		}
		return
	}

	if viper.GetString("app.mode") == "prod" {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
		f, _ := os.Create(fmt.Sprintf("%s/gin.log", viper.GetString("log.path")))
		gin.DefaultWriter = io.MultiWriter(f)
	}

	r := gin.Default()
	r.Use(middleware.Correlation())
	r.Use(middleware.Auth())
	r.Use(middleware.Logger())
	r.GET("/", controller.Index)
	r.GET("/_healthcheck", controller.HealthCheck)
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.String(http.StatusNoContent, "")
	})

	r.GET("/api/channel/:name", controller.GetChannelByName)
	r.POST("/api/channel", controller.CreateChannel)
	r.DELETE("/api/channel/:name", controller.DeleteChannelByName)
	r.PUT("/api/channel/:name", controller.UpdateChannelByName)

	r.GET("/api/client/:id", controller.GetClientByID)
	r.POST("/api/client", controller.CreateClient)
	r.DELETE("/api/client/:id", controller.DeleteClientByID)
	r.PUT("/api/client/:id/unsubscribe", controller.Unsubscribe)
	r.PUT("/api/client/:id/subscribe", controller.Subscribe)

	r.GET("/api/node", controller.GetNodeInfo)
	r.GET("/api/metrics", controller.GetMetrics)

	r.GET("/api/config/:key", controller.GetConfigByKey)
	r.POST("/api/config", controller.CreateConfig)
	r.DELETE("/api/config/:key", controller.DeleteConfigByKey)
	r.PUT("/api/config/:key", controller.UpdateConfigByKey)

	socket := &controller.Websocket{}
	socket.Init()

	r.GET("/ws/:id/:token", func(c *gin.Context) {
		socket.HandleConnections(
			c.Writer,
			c.Request,
			c.Param("id"),
			c.Param("token"),
			c.Request.Header.Get("X-Correlation-ID"),
		)
	})

	r.POST("/api/broadcast", func(c *gin.Context) {
		rawBody, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  "Invalid request",
			})
			return
		}
		socket.BroadcastAction(c, rawBody)
	})

	r.POST("/api/publish", func(c *gin.Context) {
		rawBody, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  "Invalid request",
			})
			return
		}
		socket.PublishAction(c, rawBody)
	})

	go socket.HandleMessages()

	if viper.GetBool("app.tls.status") {
		r.RunTLS(
			fmt.Sprintf(":%s", strconv.Itoa(viper.GetInt("app.port"))),
			viper.GetString("app.tls.pemPath"),
			viper.GetString("app.tls.keyPath"),
		)
	} else {
		r.Run(
			fmt.Sprintf(":%s", strconv.Itoa(viper.GetInt("app.port"))),
		)
	}
}
