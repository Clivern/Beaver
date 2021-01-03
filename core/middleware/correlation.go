// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package middleware

import (
	"strings"

	"github.com/clivern/beaver/core/util"

	"github.com/gin-gonic/gin"
)

// Correlation middleware
func Correlation() gin.HandlerFunc {
	return func(c *gin.Context) {
		corralationID := c.Request.Header.Get("X-Correlation-ID")

		if strings.TrimSpace(corralationID) == "" {
			c.Request.Header.Add("X-Correlation-ID", util.GenerateUUID4())
		}
		c.Next()
	}
}
