// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/clivern/beaver/core/driver"
	"github.com/clivern/beaver/core/module"
	"github.com/clivern/beaver/core/util"

	"github.com/gin-gonic/gin"
)

// GetClientByID controller
func GetClientByID(c *gin.Context) {

	var clientResult module.ClientResult

	validate := util.Validator{}

	ID := c.Param("id")

	if validate.IsEmpty(ID) || !validate.IsUUID4(ID) {
		c.JSON(http.StatusBadRequest, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Client ID is invalid UUID v4.",
		})
		return
	}

	db := driver.NewEtcdDriver()

	err := db.Connect()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Internal server error",
		})
		return
	}

	defer db.Close()

	client := module.NewClient(db)

	clientResult, err = client.GetClientByID(ID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  fmt.Sprintf("Client with ID %s not found.", ID),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        clientResult.ID,
		"token":     clientResult.Token,
		"channels":  clientResult.Channels,
		"createdAt": time.Unix(clientResult.CreatedAt, 0),
		"updatedAt": time.Unix(clientResult.UpdatedAt, 0),
	})
}

// CreateClient controller
func CreateClient(c *gin.Context) {

	var clientResult module.ClientResult

	validate := util.Validator{}

	rawBody, err := c.GetRawData()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errorMessage":  "Error! Invalid request",
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
		})
		return
	}

	ok, err := clientResult.LoadFromJSON(rawBody)

	if !ok || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errorMessage":  "Error! Invalid request",
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
		})
		return
	}

	if !validate.IsSlugs(clientResult.Channels, 3, 60) {
		c.JSON(http.StatusBadRequest, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Provided client channels are invalid.",
		})
		return
	}

	db := driver.NewEtcdDriver()

	err = db.Connect()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Internal server error",
		})
		return
	}

	defer db.Close()

	client := module.NewClient(db)
	channel := module.NewChannel(db)

	ok, err = channel.ChannelsExist(clientResult.Channels)

	if !ok || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Error! Provided client channels not found",
		})
		return
	}

	newClient := module.GenerateClient(clientResult.Channels)

	ok, err = client.CreateClient(*newClient)

	if !ok || err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        newClient.ID,
		"token":     newClient.Token,
		"channels":  newClient.Channels,
		"createdAt": time.Unix(newClient.CreatedAt, 0),
		"updatedAt": time.Unix(newClient.UpdatedAt, 0),
	})
}

// DeleteClientByID controller
func DeleteClientByID(c *gin.Context) {

	validate := util.Validator{}
	ID := c.Param("id")

	if validate.IsEmpty(ID) || !validate.IsUUID4(ID) {
		c.JSON(http.StatusBadRequest, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Client ID is invalid.",
		})
		return
	}

	db := driver.NewEtcdDriver()

	err := db.Connect()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Internal server error",
		})
		return
	}

	defer db.Close()

	client := module.NewClient(db)

	_, err = client.GetClientByID(ID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  fmt.Sprintf("Client with ID %s not found.", ID),
		})
		return
	}

	ok, err := client.DeleteClientByID(ID)

	if !ok || err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Internal server error",
		})
		return
	}

	c.Status(http.StatusNoContent)
}

// Unsubscribe controller
func Unsubscribe(c *gin.Context) {

	var clientResult module.ClientResult

	validate := util.Validator{}

	ID := c.Param("id")

	if validate.IsEmpty(ID) || !validate.IsUUID4(ID) {
		c.JSON(http.StatusBadRequest, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Error! Client ID is invalid UUID v4",
		})
		return
	}

	rawBody, err := c.GetRawData()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Error! Invalid request",
		})
		return
	}

	ok, err := clientResult.LoadFromJSON(rawBody)

	if !ok || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Error! Invalid request",
		})
		return
	}

	if !validate.IsSlugs(clientResult.Channels, 3, 60) {
		c.JSON(http.StatusBadRequest, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Error! Provided client channels are invalid",
		})
		return
	}

	db := driver.NewEtcdDriver()

	err = db.Connect()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Internal server error",
		})
		return
	}

	defer db.Close()

	client := module.NewClient(db)
	channel := module.NewChannel(db)

	ok, err = channel.ChannelsExist(clientResult.Channels)

	if !ok || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Error! Channels not found",
		})
		return
	}

	_, err = client.GetClientByID(ID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  fmt.Sprintf("Client with ID %s not found.", ID),
		})
		return
	}

	ok, err = client.Unsubscribe(ID, clientResult.Channels)

	if !ok || err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Internal server error",
		})
		return
	}

	c.Status(http.StatusOK)
}

// Subscribe controller
func Subscribe(c *gin.Context) {
	var clientResult module.ClientResult

	validate := util.Validator{}

	ID := c.Param("id")

	if validate.IsEmpty(ID) || !validate.IsUUID4(ID) {
		c.JSON(http.StatusBadRequest, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Client ID is invalid UUID v4.",
		})
		return
	}

	rawBody, err := c.GetRawData()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Invalid request",
		})
		return
	}

	ok, err := clientResult.LoadFromJSON(rawBody)

	if !ok || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Invalid request",
		})
		return
	}

	if !validate.IsSlugs(clientResult.Channels, 3, 60) {
		c.JSON(http.StatusBadRequest, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Provided client channels are invalid.",
		})
		return
	}

	db := driver.NewEtcdDriver()

	err = db.Connect()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Internal server error",
		})
		return
	}

	defer db.Close()

	client := module.NewClient(db)
	channel := module.NewChannel(db)

	ok, err = channel.ChannelsExist(clientResult.Channels)

	if !ok || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Error! Channels not found",
		})
		return
	}

	_, err = client.GetClientByID(ID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  fmt.Sprintf("Client with ID %s not found.", ID),
		})
		return
	}

	ok, err = client.Subscribe(ID, clientResult.Channels)

	if !ok || err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Internal server error",
		})
		return
	}

	c.Status(http.StatusOK)
}
