// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"github.com/gin-gonic/gin"
)

// GetNodeInfo controller
func GetNodeInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
