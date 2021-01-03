// Copyright 2018 Clivern. All rights reserved.
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

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start beaver server",
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

		if viper.GetString("app.log.output") != "stdout" {
			dir, _ := filepath.Split(viper.GetString("app.log.output"))

			if !util.DirExists(dir) {
				if _, err := util.EnsureDir(dir, 755); err != nil {
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

		viper.SetDefault("app.name", util.GenerateUUID4())

		r := gin.Default()

		r.Use(middleware.Correlation())
		r.Use(middleware.Auth())
		r.Use(middleware.Cors())
		r.Use(middleware.Logger())
		r.Use(middleware.Metric())

		r.GET("/favicon.ico", func(c *gin.Context) {
			c.String(http.StatusNoContent, "")
		})

		r.GET("/", controller.Index)
		r.GET("/_health", controller.HealthCheck)

		apiv1 := r.Group("/api/v2")
		{
			apiv1.GET("/client/:id", controller.GetClientByID)
			apiv1.POST("/client", controller.CreateClient)
			apiv1.DELETE("/client/:id", controller.DeleteClientByID)
			apiv1.PUT("/client/:id/unsubscribe", controller.Unsubscribe)
			apiv1.PUT("/client/:id/subscribe", controller.Subscribe)

			apiv1.GET("/channel/:name", controller.GetChannelByName)
			apiv1.POST("/channel", controller.CreateChannel)
			apiv1.DELETE("/channel/:name", controller.DeleteChannelByName)
			apiv1.PUT("/channel/:name", controller.UpdateChannelByName)
			apiv1.GET("/channel/:name/subscribers", controller.GetChannelSubscribersByName)
			apiv1.GET("/channel/:name/listeners", controller.GetChannelListenersByName)
		}

		r.GET(
			viper.GetString("app.metrics.prometheus.endpoint"),
			gin.WrapH(controller.Metrics()),
		)

		// Goroutine for node heartbeat
		go controller.Heartbeat()

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
	serveCmd.Flags().StringVarP(
		&config,
		"config",
		"c",
		"config.prod.yml",
		"Absolute path to config file (required)",
	)
	serveCmd.MarkFlagRequired("config")
	rootCmd.AddCommand(serveCmd)
}
