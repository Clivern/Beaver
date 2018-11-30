// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetClientByID controller
func GetClientByID(c *gin.Context) {
	ID := c.Param("id")

	fmt.Println(ID)

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

// DeleteClientByID controller
func DeleteClientByID(c *gin.Context) {
	ID := c.Param("id")

	fmt.Println(ID)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// UpdateClientByID controller
func UpdateClientByID(c *gin.Context) {
	ID := c.Param("id")

	fmt.Println(ID)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
