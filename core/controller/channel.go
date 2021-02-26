// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetChannelByName controller
func GetChannelByName(c *gin.Context) {
	c.Status(http.StatusOK)
}

// CreateChannel controller
func CreateChannel(c *gin.Context) {
	c.Status(http.StatusOK)
}

// DeleteChannelByName controller
func DeleteChannelByName(c *gin.Context) {
	c.Status(http.StatusOK)
}

// UpdateChannelByName controller
func UpdateChannelByName(c *gin.Context) {
	c.Status(http.StatusOK)
}

// GetChannelSubscribersByName controller
func GetChannelSubscribersByName(c *gin.Context) {
	c.Status(http.StatusOK)
}

// GetChannelListenersByName controller
func GetChannelListenersByName(c *gin.Context) {
	c.Status(http.StatusOK)
}

// GetChannelMessages controller
func GetChannelMessages(c *gin.Context) {
	c.Status(http.StatusOK)
}
