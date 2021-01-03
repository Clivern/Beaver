// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

// This module used to broadcast messages to all cluster nodes
// By default if the node is part of a cluster, it will join
// RabbitMQ

type Broadcast struct {
}

type Message struct {
}

func (b *Broadcast) Publish(m *Message) error {
	return nil
}

func (b *Broadcast) Listen() {

}
