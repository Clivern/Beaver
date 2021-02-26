// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	totalNodes = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "beaver",
			Name:      "cluster_total_nodes",
			Help:      "Total nodes in the cluster",
		})

	totalClients = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "beaver",
			Name:      "cluster_total_clients",
			Help:      "Total clients in the cluster",
		})

	totalChannels = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "beaver",
			Name:      "cluster_total_channels",
			Help:      "Total channels in the cluster",
		})
)

func init() {
	prometheus.MustRegister(totalNodes)
	prometheus.MustRegister(totalClients)
	prometheus.MustRegister(totalChannels)
}

// Heartbeat node heartbeat
func Heartbeat() {
}
