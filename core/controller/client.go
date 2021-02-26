// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetClientByID controller
func GetClientByID(c *gin.Context) {
	c.Status(http.StatusOK)
}

// CreateClient controller
func CreateClient(c *gin.Context) {
	c.Status(http.StatusOK)
}

// DeleteClientByID controller
func DeleteClientByID(c *gin.Context) {
	c.Status(http.StatusOK)
}

// Unsubscribe controller
func Unsubscribe(c *gin.Context) {
	c.Status(http.StatusOK)
}

// Subscribe controller
func Subscribe(c *gin.Context) {
	c.Status(http.StatusOK)
}

// GetClientMessages controller
func GetClientMessages(c *gin.Context) {
	c.Status(http.StatusOK)
}
