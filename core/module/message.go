// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"github.com/clivern/beaver/core/driver"

	"github.com/gocql/gocql"
)

// MessageModule type
type MessageModule struct {
	db *driver.Cassandra
}

// MessageModel struct
type MessageModel struct {
	ID            gocql.UUID `json:"id"`
	FromClientId  gocql.UUID `json:"from_client_id"`
	ToChannelId   gocql.UUID `json:"to_channel_id"`
	ToClientId    gocql.UUID `json:"to_client_id"`
	Message       string     `json:"message"`
	ToChannelName string     `json:"to_channel_name"`
	CreatedAt     int64      `json:"created_at"`
	UpdatedAt     int64      `json:"updated_at"`
}

// NewMessageModel creates a message module instance
func NewMessageModel(db *driver.Cassandra) *MessageModule {
	result := new(MessageModule)
	result.db = db

	return result
}
