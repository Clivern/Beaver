// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/clivern/beaver/internal/pkg/logger"
	"github.com/gin-gonic/gin"
	"time"
)

// Logger middleware
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		t := time.Now()
		rawBody, _ := c.GetRawData()

		c.Next()

		// after request
		latency := time.Since(t)
		status := c.Writer.Status()
		size := c.Writer.Size()

		logger.Infof(
			"Request %s:%s %s - Response Code %d, Size %d Time Spent %s",
			c.Request.Method,
			c.Request.URL,
			string(rawBody),
			status,
			size,
			latency,
		)
	}
}
