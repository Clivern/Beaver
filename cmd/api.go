// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/clivern/beaver/core/controller"
	"github.com/clivern/beaver/core/middleware"
	"github.com/clivern/beaver/core/util"

	"github.com/drone/envsubst"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var towerCmd = &cobra.Command{
	Use:   "api",
	Short: "Start beaver api server",
	Run: func(cmd *cobra.Command, args []string) {
		configUnparsed, err := ioutil.ReadFile(config)

		if err != nil {
			panic(fmt.Sprintf(
				"Error while reading config file [%s]: %s",
				config,
				err.Error(),
			))
		}

		configParsed, err := envsubst.EvalEnv(string(configUnparsed))

		if err != nil {
			panic(fmt.Sprintf(
				"Error while parsing config file [%s]: %s",
				config,
				err.Error(),
			))
		}

		viper.SetConfigType("yaml")
		err = viper.ReadConfig(bytes.NewBuffer([]byte(configParsed)))

		if err != nil {
			panic(fmt.Sprintf(
				"Error while loading configs [%s]: %s",
				config,
				err.Error(),
			))
		}

		viper.SetDefault("app.name", util.GenerateUUID4())

		if viper.GetString("app.log.output") != "stdout" {
			dir, _ := filepath.Split(viper.GetString("app.log.output"))

			if !util.DirExists(dir) {
				if _, err := util.EnsureDir(dir, 775); err != nil {
					panic(fmt.Sprintf(
						"Directory [%s] creation failed with error: %s",
						dir,
						err.Error(),
					))
				}
			}

			if !util.FileExists(viper.GetString("app.log.output")) {
				f, err := os.Create(viper.GetString("app.log.output"))
				if err != nil {
					panic(fmt.Sprintf(
						"Error while creating log file [%s]: %s",
						viper.GetString("app.log.output"),
						err.Error(),
					))
				}
				defer f.Close()
			}
		}

		if viper.GetString("app.log.output") == "stdout" {
			gin.DefaultWriter = os.Stdout
			log.SetOutput(os.Stdout)
		} else {
			f, _ := os.Create(viper.GetString("app.log.output"))
			gin.DefaultWriter = io.MultiWriter(f)
			log.SetOutput(f)
		}

		lvl := strings.ToLower(viper.GetString("app.log.level"))
		level, err := log.ParseLevel(lvl)

		if err != nil {
			level = log.InfoLevel
		}

		log.SetLevel(level)

		if viper.GetString("app.mode") == "prod" {
			gin.SetMode(gin.ReleaseMode)
			gin.DefaultWriter = ioutil.Discard
			gin.DisableConsoleColor()
		}

		if viper.GetString("app.log.format") == "json" {
			log.SetFormatter(&log.JSONFormatter{})
		} else {
			log.SetFormatter(&log.TextFormatter{})
		}

		r := gin.Default()

		r.Use(middleware.Cors())
		r.Use(middleware.Correlation())
		r.Use(middleware.Logger())
		r.Use(middleware.Metric())
		r.Use(middleware.Auth())

		r.GET("/favicon.ico", func(c *gin.Context) {
			c.String(http.StatusNoContent, "")
		})

		r.GET("/", controller.Home)

		r.GET("/_health", controller.Health)

		r.GET(
			viper.GetString("app.metrics.prometheus.endpoint"),
			gin.WrapH(controller.Metrics()),
		)

		r.GET("/api/channel/:name", controller.GetChannelByName)
		r.POST("/api/channel", controller.CreateChannel)
		r.DELETE("/api/channel/:name", controller.DeleteChannelByName)
		r.PUT("/api/channel/:name", controller.UpdateChannelByName)

		r.GET("/api/client/:id", controller.GetClientByID)
		r.POST("/api/client", controller.CreateClient)
		r.DELETE("/api/client/:id", controller.DeleteClientByID)
		r.PUT("/api/client/:id/unsubscribe", controller.Unsubscribe)
		r.PUT("/api/client/:id/subscribe", controller.Subscribe)

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

		var runerr error

		if viper.GetBool("app.tls.status") {
			runerr = r.RunTLS(
				fmt.Sprintf(":%s", strconv.Itoa(viper.GetInt("app.port"))),
				viper.GetString("app.tls.pemPath"),
				viper.GetString("app.tls.keyPath"),
			)
		} else {
			runerr = r.Run(
				fmt.Sprintf(":%s", strconv.Itoa(viper.GetInt("app.port"))),
			)
		}

		if runerr != nil {
			panic(runerr.Error())
		}
	},
}

func init() {
	towerCmd.Flags().StringVarP(
		&config,
		"config",
		"c",
		"config.prod.yml",
		"Absolute path to config file (required)",
	)
	towerCmd.MarkFlagRequired("config")
	rootCmd.AddCommand(towerCmd)
}
