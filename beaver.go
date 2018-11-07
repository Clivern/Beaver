// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/clivern/beaver/internal/app/cmd"
	"github.com/clivern/beaver/internal/app/controller"
	"github.com/clivern/beaver/internal/pkg/pusher"
	"github.com/clivern/beaver/internal/pkg/utils"
	"github.com/gin-gonic/gin"
)

func main() {

	var exec string

	utils.PrintBanner()

	// Load config.json file and store on env
	config := &utils.Config{}
	config.Load("config.dist.json")
	// This will never override ENV Vars if exists
	config.Cache()
	config.GinEnv()

	flag.StringVar(&exec, "exec", "", "exec")
	flag.Parse()
	if exec != "" {
		switch exec {
		case "migrate.up":
			migrate := &cmd.Migrate{}
			migrate.Up()
		case "migrate.down":
			migrate := &cmd.Migrate{}
			migrate.Down()
		case "migrate.status":
			migrate := &cmd.Migrate{}
			migrate.Status()
		case "health":
			health := &cmd.Health{}
			health.Status()
		default:
			utils.PrintCommands()
		}
		return
	}

	safeMigrate := &cmd.Migrate{}
	safeMigrate.SafeUp()

	if os.Getenv("AppMode") == "prod" {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
		f, _ := os.Create("var/logs/gin.log")
		gin.DefaultWriter = io.MultiWriter(f)
	}

	files := []string{}

	files = utils.ListFiles("internal/scheme")
	files = utils.FilterFiles(files, []string{"down"})
	fmt.Println(files)

	r := gin.Default()
	r.Static("/static", "./web/static/")
	r.LoadHTMLGlob("web/template/*")
	r.GET("/", controller.Index)
	r.GET("/_healthcheck", controller.HealthCheck)
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.String(http.StatusNoContent, "")
	})

	r.POST("/apps/:app_id/events", controller.Events)
	r.GET("/apps/:app_id/channels", controller.Channels)
	r.GET("/apps/:app_id/channels/:channel_name", controller.Channel)
	r.GET("/apps/:app_id/channels/:channel_name/users", controller.ChannelUsers)

	socket := &pusher.Websocket{}
	socket.Init()
	r.GET("/app/:key", func(c *gin.Context) {
		socket.HandleConnections(c.Writer, c.Request, c.Param("key"))
	})

	go socket.HandleMessages()

	r.Run()
}
