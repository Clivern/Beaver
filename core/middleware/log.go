// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package middleware

import (
	"bytes"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Logger middleware
// TODO Hide Secure Data from Logs
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		var bodyBytes []byte

		// Workaround for issue https://github.com/gin-gonic/gin/issues/1651
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		log.WithFields(log.Fields{
			"correlation_id": c.GetHeader("x-correlation-id"),
			"http_method":    c.Request.Method,
			"http_path":      c.Request.URL.Path,
			"request_body":   string(bodyBytes),
		}).Info("Request started")

		c.Next()

		// after request
		status := c.Writer.Status()
		size := c.Writer.Size()

		log.WithFields(log.Fields{
			"correlation_id": c.GetHeader("x-correlation-id"),
			"http_status":    status,
			"response_size":  size,
		}).Info(`Request finished`)
	}
}
