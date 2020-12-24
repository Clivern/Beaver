// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"github.com/clivern/beaver/internal/app/api"
	"github.com/clivern/beaver/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetClientByID controller
func GetClientByID(c *gin.Context) {

	var clientResult api.ClientResult
	validate := utils.Validator{}

	ID := c.Param("id")

	if validate.IsEmpty(ID) || !validate.IsUUID4(ID) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Client ID is invalid.",
		})
		return
	}

	client := api.Client{
		CorrelationID: c.Request.Header.Get("X-Correlation-ID"),
	}

	if !client.Init() {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "error",
			"error":  "Internal server error",
		})
		return
	}

	clientResult, err := client.GetClientByID(ID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         clientResult.ID,
		"token":      clientResult.Token,
		"channels":   clientResult.Channels,
		"created_at": clientResult.CreatedAt,
		"updated_at": clientResult.UpdatedAt,
	})
}

// CreateClient controller
func CreateClient(c *gin.Context) {

	var clientResult api.ClientResult
	validate := utils.Validator{}

	client := api.Client{
		CorrelationID: c.Request.Header.Get("X-Correlation-ID"),
	}
	channel := api.Channel{
		CorrelationID: c.Request.Header.Get("X-Correlation-ID"),
	}

	rawBody, err := c.GetRawData()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid request",
		})
		return
	}

	ok, err := clientResult.LoadFromJSON(rawBody)

	if !ok || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid request",
		})
		return
	}

	if !validate.IsSlugs(clientResult.Channels, 3, 60) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Provided client channels are invalid.",
		})
		return
	}

	if !client.Init() || !channel.Init() {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "error",
			"error":  "Internal server error",
		})
		return
	}

	ok, err = channel.ChannelsExist(clientResult.Channels)

	if !ok || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	ok, err = clientResult.GenerateClient()

	if !ok || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	ok, err = client.CreateClient(clientResult)

	if !ok || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         clientResult.ID,
		"token":      clientResult.Token,
		"channels":   clientResult.Channels,
		"created_at": clientResult.CreatedAt,
		"updated_at": clientResult.UpdatedAt,
	})
}

// DeleteClientByID controller
func DeleteClientByID(c *gin.Context) {

	validate := utils.Validator{}
	ID := c.Param("id")

	if validate.IsEmpty(ID) || !validate.IsUUID4(ID) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Client ID is invalid.",
		})
		return
	}

	client := api.Client{
		CorrelationID: c.Request.Header.Get("X-Correlation-ID"),
	}

	if !client.Init() {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "error",
			"error":  "Internal server error",
		})
		return
	}

	_, err := client.DeleteClientByID(ID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

// Unsubscribe controller
func Unsubscribe(c *gin.Context) {

	var clientResult api.ClientResult
	validate := utils.Validator{}
	ID := c.Param("id")

	if validate.IsEmpty(ID) || !validate.IsUUID4(ID) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Client ID is invalid.",
		})
		return
	}

	client := api.Client{
		CorrelationID: c.Request.Header.Get("X-Correlation-ID"),
	}
	channel := api.Channel{
		CorrelationID: c.Request.Header.Get("X-Correlation-ID"),
	}

	rawBody, err := c.GetRawData()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid request",
		})
		return
	}

	ok, err := clientResult.LoadFromJSON(rawBody)

	if !ok || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid request",
		})
		return
	}

	if !validate.IsSlugs(clientResult.Channels, 3, 60) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Provided client channels are invalid.",
		})
		return
	}

	if !client.Init() || !channel.Init() {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "error",
			"error":  "Internal server error",
		})
		return
	}

	ok, err = channel.ChannelsExist(clientResult.Channels)

	if !ok || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	ok, err = client.Unsubscribe(ID, clientResult.Channels)

	if !ok || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}

// Subscribe controller
func Subscribe(c *gin.Context) {

	var clientResult api.ClientResult
	validate := utils.Validator{}
	ID := c.Param("id")

	if validate.IsEmpty(ID) || !validate.IsUUID4(ID) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Client ID is invalid.",
		})
		return
	}

	client := api.Client{
		CorrelationID: c.Request.Header.Get("X-Correlation-ID"),
	}
	channel := api.Channel{
		CorrelationID: c.Request.Header.Get("X-Correlation-ID"),
	}

	rawBody, err := c.GetRawData()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid request",
		})
		return
	}

	ok, err := clientResult.LoadFromJSON(rawBody)

	if !ok || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid request",
		})
		return
	}

	if !validate.IsSlugs(clientResult.Channels, 3, 60) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Provided client channels are invalid.",
		})
		return
	}

	if !client.Init() || !channel.Init() {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "error",
			"error":  "Internal server error",
		})
		return
	}

	ok, err = channel.ChannelsExist(clientResult.Channels)

	if !ok || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	ok, err = client.Subscribe(ID, clientResult.Channels)

	if !ok || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}
