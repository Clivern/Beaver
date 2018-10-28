// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package pusher

// A Channel Subscription
type Subscription struct {
	Connection *Connection
	ID         string
	Data       string
}

// Create a new Subscription
func NewSubscription(conn *Connection, data string) *Subscription {
	return &Subscription{Connection: conn, Data: data}
}
