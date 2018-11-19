// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package middleware

import (
	"bytes"
	"github.com/clivern/beaver/internal/pkg/logger"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"time"
)

// Logger middleware
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		var bodyBytes []byte
		t := time.Now()

		// Workaround for issue https://github.com/gin-gonic/gin/issues/1651
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		c.Next()

		// after request
		latency := time.Since(t)
		status := c.Writer.Status()
		size := c.Writer.Size()

		logger.Infof(
			"Request %s -> %s %s - Response Code %d, Size %d Time Spent %s",
			c.Request.Method,
			c.Request.URL,
			string(bodyBytes),
			status,
			size,
			latency,
		)
	}
}
