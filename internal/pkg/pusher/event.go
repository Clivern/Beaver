// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package pusher

import (
	"encoding/json"
)

type RawEvent struct {
	Event   string          `json:"event"`
	Channel string          `json:"channel"`
	Data    json.RawMessage `json:"data"`
}

type ResponseEvent struct {
	Event   string      `json:"event"`
	Channel string      `json:"channel"`
	Data    interface{} `json:"data"`
}

// The response event that is broadcasted to the client sockets
func NewResponseEvent(name, channel string, data interface{}) ResponseEvent {
	return ResponseEvent{Event: name, Channel: channel, Data: data}
}
