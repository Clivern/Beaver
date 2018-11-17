// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// GetConfigByKey controller
func GetConfigByKey(c *gin.Context) {
	Key := c.Param("key")

	fmt.Println(Key)

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

// DeleteConfigByKey controller
func DeleteConfigByKey(c *gin.Context) {
	Key := c.Param("key")

	fmt.Println(Key)

	c.JSON(200, gin.H{
		"status": "ok",
	})
}

// UpdateConfigByKey controller
func UpdateConfigByKey(c *gin.Context) {
	Key := c.Param("key")

	fmt.Println(Key)

	c.JSON(200, gin.H{
		"status": "ok",
	})
}
