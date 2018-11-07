// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"github.com/gin-gonic/gin"

	"fmt"
	"github.com/gorilla/websocket"
	_ "log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(_ *http.Request) bool {
		return true
	},
}

// Websocket controller
func Websocket(c *gin.Context) {
	//key := c.Param("key")

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()

	if err != nil {
		fmt.Println("1")
		fmt.Println(err.Error())
		return
	}

	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			fmt.Println("2")
			fmt.Println(err.Error())
			break
		}

		conn.WriteJSON(fmt.Sprintf(`{"item":"%s"}`, message))
		fmt.Println(string(message))
	}
}
