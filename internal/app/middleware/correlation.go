// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/clivern/beaver/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

// Correlation middleware
func Correlation() gin.HandlerFunc {
	return func(c *gin.Context) {
		corralationID := c.Request.Header.Get("X-Correlation-ID")
		if strings.TrimSpace(corralationID) == "" {
			c.Request.Header.Add("CorrelationID", utils.GenerateUUID())
		}
		c.Next()
	}
}
