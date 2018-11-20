// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetChannelByName controller
func GetChannelByName(c *gin.Context) {
	name := c.Param("name")

	fmt.Println(name)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
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

	fmt.Println(name)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// UpdateChannelByName controller
func UpdateChannelByName(c *gin.Context) {
	name := c.Param("name")

	fmt.Println(name)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
