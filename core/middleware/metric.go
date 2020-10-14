// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

var (
	httpRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "beaver",
			Name:      "total_http_requests",
			Help:      "How many HTTP requests processed, partitioned by status code and HTTP method.",
		}, []string{"code", "method", "handler", "host", "url"})

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Subsystem: "beaver",
			Name:      "request_duration_seconds",
			Help:      "The HTTP request latencies in seconds.",
		},
		[]string{"code", "method", "url"},
	)

	responseSize = prometheus.NewSummary(
		prometheus.SummaryOpts{
			Namespace: "beaver",
			Name:      "response_size_bytes",
			Help:      "The HTTP response sizes in bytes.",
		},
	)
)

func init() {
	prometheus.MustRegister(httpRequests)
	prometheus.MustRegister(requestDuration)
	prometheus.MustRegister(responseSize)
}

// Metric middleware
func Metric() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		start := time.Now()

		c.Next()

		// after request
		elapsed := float64(time.Since(start)) / float64(time.Second)

		log.WithFields(log.Fields{
			"correlation_id": c.Request.Header.Get("X-Correlation-ID"),
		}).Info(`Collecting metrics`)

		// Collect Metrics
		httpRequests.WithLabelValues(
			strconv.Itoa(c.Writer.Status()),
			c.Request.Method,
			c.HandlerName(),
			c.Request.Host,
			c.Request.URL.Path,
		).Inc()

		requestDuration.WithLabelValues(
			strconv.Itoa(c.Writer.Status()),
			c.Request.Method,
			c.Request.URL.Path,
		).Observe(elapsed)

		responseSize.Observe(float64(c.Writer.Size()))
	}
}
