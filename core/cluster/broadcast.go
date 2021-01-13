// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"github.com/clivern/beaver/core/driver"
	"github.com/clivern/beaver/core/util"

	"github.com/spf13/viper"
)

// Message type
type Message struct {
}

// Broadcast type
type Broadcast struct {
	broker *driver.RabbitMQ
}

// NewBroadcast creates a new Broadcast object
func NewBroadcast() (*Broadcast, error) {
	var broadcast *Broadcast

	broker, err := driver.NewRabbitMQDriver()

	if err != nil {
		return broadcast, err
	}

	broadcast = &Broadcast{
		broker: broker,
	}

	return broadcast, nil
}

// Publish publish messages to broker queue
func (b *Broadcast) Publish(m *Message) error {
	body, err := util.ConvertToJSON(m)

	if err != nil {
		return err
	}

	err = b.broker.Send(
		viper.GetString("cluster.broker.rabbitmq.queue"),
		viper.GetString("app.name"),
		body,
	)

	return err
}

// Listen listens to incoming broker messages
func (b *Broadcast) Listen(callback func(msg string)) error {
	err := b.broker.Consume(
		viper.GetString("cluster.broker.rabbitmq.queue"),
		viper.GetString("app.name"),
		callback,
	)

	return err
}

// Close closes broker connection
func (b *Broadcast) Close() {
	b.broker.Close()
}
