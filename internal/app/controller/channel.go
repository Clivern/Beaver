// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"github.com/clivern/beaver/internal/app/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetChannelByName controller
func GetChannelByName(c *gin.Context) {
	var channelResult api.ChannelResult

	name := c.Param("name")
	channel := api.Channel{}

	if !channel.Init() {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "error",
			"error":  "Internal server error",
		})
		return
	}

	channelResult, err := channel.GetChannelByName(name)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":        channelResult.Name,
		"type":        channelResult.Type,
		"listeners":   channelResult.Listeners,
		"subscribers": channelResult.Subscribers,
		"created_at":  channelResult.CreatedAt,
		"updated_at":  channelResult.UpdatedAt,
	})
}

// CreateChannel controller
func CreateChannel(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// DeleteChannelByName controller
func DeleteChannelByName(c *gin.Context) {

	name := c.Param("name")
	channel := api.Channel{}

	if !channel.Init() {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "error",
			"error":  "Internal server error",
		})
		return
	}

	_, err := channel.DeleteChannelByName(name)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

// UpdateChannelByName controller
func UpdateChannelByName(c *gin.Context) {
	name := c.Param("name")

	fmt.Println(name)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
