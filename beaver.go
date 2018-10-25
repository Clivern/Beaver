// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"io"
	"net/http"
	"os"

	"github.com/clivern/beaver/internal/app/controller"
	"github.com/clivern/beaver/internal/pkg/broadcast"
	"github.com/gin-gonic/gin"
)

func main() {

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
	r.GET("/chat", controller.Chat)

	socket := &broadcast.Websocket{}
	socket.Init()
	r.GET("/ws", func(c *gin.Context) {
		socket.HandleConnections(c.Writer, c.Request, c.DefaultQuery("channel", ""))
	})
	go socket.HandleMessages()

	r.Run()
}
