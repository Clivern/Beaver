// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// GetChannelByID controller
func GetChannelByID(c *gin.Context) {
	ID := c.Param("id")

	fmt.Println(ID)

	c.JSON(200, gin.H{
		"status": "ok",
	})
}

// CreateChannel controller
func CreateChannel(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

// DeleteChannelByID controller
func DeleteChannelByID(c *gin.Context) {
	ID := c.Param("id")

	fmt.Println(ID)

	c.JSON(200, gin.H{
		"status": "ok",
	})
}

// UpdateChannelByID controller
func UpdateChannelByID(c *gin.Context) {
	ID := c.Param("id")

	fmt.Println(ID)

	c.JSON(200, gin.H{
		"status": "ok",
	})
}
