// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/clivern/beaver/internal/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

// Auth middleware
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method

		if strings.Contains(path, "/api/") {
			authToken := c.GetHeader("X-AUTH-TOKEN")
			if authToken != viper.GetString("api.token") {
				logger.Infof(
					`Unauthorized access to %s:%s with token %s {"correlationId":"%s"}`,
					method,
					path,
					authToken,
					c.Request.Header.Get("X-Correlation-ID"),
				)
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		}
	}
}
