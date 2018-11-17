// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// GetConfigByID controller
func GetConfigByID(c *gin.Context) {
	ID := c.Param("id")

	fmt.Println(ID)

	c.JSON(200, gin.H{
		"status": "ok",
	})
}

// CreateConfig controller
func CreateConfig(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

// DeleteConfigByID controller
func DeleteConfigByID(c *gin.Context) {
	ID := c.Param("id")

	fmt.Println(ID)

	c.JSON(200, gin.H{
		"status": "ok",
	})
}

// UpdateConfigByID controller
func UpdateConfigByID(c *gin.Context) {
	ID := c.Param("id")

	fmt.Println(ID)

	c.JSON(200, gin.H{
		"status": "ok",
	})
}
