// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"io"
	"net/http"
	"os"

	"github.com/clivern/beaver/internal/app/api"
	"github.com/clivern/beaver/internal/app/cmd"
	"github.com/clivern/beaver/internal/app/controller"
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

	cmd.CreateMigrationTable()

	if exec != "" {
		switch exec {
		case "migrate.up":
			cmd.MigrationUp()
		case "migrate.down":
			cmd.MigrationDown()
		case "migrate.status":
			cmd.MigrationStatus()
		case "health":
			cmd.HealthStatus()
		default:
			utils.PrintCommands()
		}
		return
	}

	cmd.MigrationUp()

	if os.Getenv("AppMode") == "prod" {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
		f, _ := os.Create("var/logs/gin.log")
		gin.DefaultWriter = io.MultiWriter(f)
	}

	r := gin.Default()
	r.Static("/static", "./web/static/")
	r.LoadHTMLGlob("web/template/*")
	r.GET("/", controller.Index)
	r.GET("/_healthcheck", controller.HealthCheck)
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.String(http.StatusNoContent, "")
	})

	socket := &api.Websocket{}
	socket.Init()
	r.GET("/ws", func(c *gin.Context) {
		socket.HandleConnections(c.Writer, c.Request)
	})

	go socket.HandleMessages()

	r.Run()
}
