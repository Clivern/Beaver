// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetClientByUUID controller
func GetClientByUUID(c *gin.Context) {
	UUID := c.Param("uuid")

	fmt.Println(UUID)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// CreateClient controller
func CreateClient(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// DeleteClientByUUID controller
func DeleteClientByUUID(c *gin.Context) {
	UUID := c.Param("uuid")

	fmt.Println(UUID)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// UpdateClientByUUID controller
func UpdateClientByUUID(c *gin.Context) {
	UUID := c.Param("uuid")

	fmt.Println(UUID)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
