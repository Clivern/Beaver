// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"github.com/gin-gonic/gin"
)

// Channel
func Channel(c *gin.Context) {
	appID := c.Param("app_id")
	channelName := c.Param("channel_name")

	c.JSON(200, gin.H{
		"status":      "ok",
		"appID":       appID,
		"channelName": channelName,
	})
}
