// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package pusher

import (
	"strings"
	"sync"
	"time"
)

// A Channel
type Channel struct {
	sync.Mutex

	ChannelID     string
	CreatedAt     time.Time
	Subscriptions map[string]*Subscription
}

// Create a new Channel
func NewChannel(channelID string) *Channel {
	return &Channel{ChannelID: channelID, CreatedAt: time.Now(), Subscriptions: make(map[string]*Subscription)}
}

// Return true if the channel has at least one subscriber
func (c *Channel) IsOccupied() bool {
	return c.TotalSubscriptions() > 0
}

// Check if the type of the channel is presence or is private
func (c *Channel) IsPresenceOrPrivate() bool {
	return c.IsPresence() || c.IsPrivate()
}

// Check if the type of the channel is public
func (c *Channel) IsPublic() bool {
	return !c.IsPresenceOrPrivate()
}

// Check if the type of the channel is presence
func (c *Channel) IsPresence() bool {
	return strings.HasPrefix(c.ChannelID, "presence-")
}

// Check if the type of the channel is private
func (c *Channel) IsPrivate() bool {
	return strings.HasPrefix(c.ChannelID, "private-")
}

// Get the total of subscribers
func (c *Channel) TotalSubscriptions() int {
	return len(c.Subscriptions)
}

// Get the total of users.
func (c *Channel) TotalUsers() int {
	total := make(map[string]int)

	for _, s := range c.Subscriptions {
		total[s.ID]++
	}

	return len(total)
}

// IsSubscribed check if the user is subscribed
func (c *Channel) IsSubscribed(conn *Connection) bool {
	_, exists := c.Subscriptions[conn.SocketID]
	return exists
}
