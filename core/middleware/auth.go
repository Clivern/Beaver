// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Auth middleware
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		if strings.Contains(path, "/api/") {
			apiKey := c.GetHeader("x-api-key")
			if apiKey != viper.GetString("app.api.key") && viper.GetString("app.api.key") != "" {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		}
	}
}
