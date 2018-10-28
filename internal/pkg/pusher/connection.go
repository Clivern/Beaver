// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package pusher

import (
	"time"
)

// socket interface to write to the client
type Socket interface {
	WriteJSON(interface{}) error
}

// An User Connection
type Connection struct {
	SocketID  string
	Socket    Socket
	CreatedAt time.Time
}

// Create a new Subscriber
func NewConnection(socketID string, s Socket) *Connection {
	return &Connection{SocketID: socketID, Socket: s, CreatedAt: time.Now()}
}

// Publish the message to websocket atached to this client
func (conn *Connection) Publish(m interface{}) {
	conn.Socket.WriteJSON(m)
}
