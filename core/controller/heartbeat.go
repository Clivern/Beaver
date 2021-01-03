// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"time"

	"github.com/clivern/beaver/core/cluster"
	"github.com/clivern/beaver/core/driver"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
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
	db := driver.NewEtcdDriver()

	err := db.Connect()

	if err != nil {
		panic(fmt.Sprintf(
			"Error while connecting to etcd: %s",
			err.Error(),
		))
	}

	defer db.Close()

	node := cluster.NewNode(db)
	stats := cluster.NewStats(db)

	log.Info(`Start heartbeat daemon`)

	count := 0

	for {
		err := node.Alive(5)

		if err != nil {
			log.WithFields(log.Fields{
				"error": err.Error(),
			}).Error(`Error while connecting to etcd`)
		} else {
			log.Debug(`Node heartbeat done`)
		}

		time.Sleep(1 * time.Second)

		count, _ = stats.GetTotalNodes()
		totalNodes.Set(float64(count))

		count, _ = stats.GetTotalClients()
		totalClients.Set(float64(count))

		count, _ = stats.GetTotalChannels()
		totalChannels.Set(float64(count))

		time.Sleep(2 * time.Second)
	}
}
