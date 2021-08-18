// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"net/http"
	"time"

	"github.com/clivern/beaver/core/api"
	"github.com/clivern/beaver/core/util"

	"github.com/gin-gonic/gin"
)

// GetChannelByName controller
func GetChannelByName(c *gin.Context) {

	var channelResult api.ChannelResult
	validate := util.Validator{}

	name := c.Param("name")

	if validate.IsEmpty(name) || !validate.IsSlug(name, 3, 60) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Channel name must be alphanumeric with length from 3 to 60",
		})
		return
	}

	channel := api.Channel{}

	if !channel.Init() {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "error",
			"error":  "Internal server error",
		})
		return
	}

	channelResult, err := channel.GetChannelByName(name)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	if channelResult.Type == "presence" {
		c.JSON(http.StatusOK, gin.H{
			"name":              channelResult.Name,
			"type":              channelResult.Type,
			"subscribers_count": channel.CountSubscribers(name),
			"listeners_count":   channel.CountListeners(name),
			"subscribers":       channel.GetSubscribers(name),
			"listeners":         channel.GetListeners(name),
			"created_at":        channelResult.CreatedAt,
			"updated_at":        channelResult.UpdatedAt,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":              channelResult.Name,
		"type":              channelResult.Type,
		"subscribers_count": channel.CountSubscribers(name),
		"listeners_count":   channel.CountListeners(name),
		"created_at":        channelResult.CreatedAt,
		"updated_at":        channelResult.UpdatedAt,
	})
}

// CreateChannel controller
func CreateChannel(c *gin.Context) {

	var channelResult api.ChannelResult
	validate := util.Validator{}

	channel := api.Channel{}

	rawBody, err := c.GetRawData()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid request",
		})
		return
	}

	ok, err := channelResult.LoadFromJSON(rawBody)

	if !ok || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid request",
		})
		return
	}

	if validate.IsEmpty(channelResult.Name) || !validate.IsSlug(channelResult.Name, 3, 60) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Channel name must be alphanumeric with length from 3 to 60",
		})
		return
	}

	if !validate.IsIn(channelResult.Type, []string{"public", "private", "presence"}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Channel type must be public, private or presence",
		})
		return
	}

	if !channel.Init() {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "error",
			"error":  "Internal server error",
		})
		return
	}

	channelResult.CreatedAt = time.Now().Unix()
	channelResult.UpdatedAt = time.Now().Unix()

	ok, err = channel.CreateChannel(channelResult)

	if !ok || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

// DeleteChannelByName controller
func DeleteChannelByName(c *gin.Context) {

	validate := util.Validator{}

	name := c.Param("name")

	if validate.IsEmpty(name) || !validate.IsSlug(name, 3, 60) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Channel name must be alphanumeric with length from 3 to 60",
		})
		return
	}

	channel := api.Channel{}

	if !channel.Init() {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "error",
			"error":  "Internal server error",
		})
		return
	}

	_, err := channel.DeleteChannelByName(name)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

// UpdateChannelByName controller
func UpdateChannelByName(c *gin.Context) {

	var channelResult api.ChannelResult
	var currentChannelResult api.ChannelResult
	validate := util.Validator{}

	channel := api.Channel{}

	rawBody, err := c.GetRawData()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid request",
		})
		return
	}

	channelResult.LoadFromJSON(rawBody)
	channelResult.Name = c.Param("name")

	if validate.IsEmpty(channelResult.Name) || !validate.IsSlug(channelResult.Name, 3, 60) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Channel name must be alphanumeric with length from 3 to 60",
		})
		return
	}

	if !validate.IsIn(channelResult.Type, []string{"public", "private", "presence"}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Channel type must be public, private or presence",
		})
		return
	}

	if !channel.Init() {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "error",
			"error":  "Internal server error",
		})
		return
	}

	currentChannelResult, err = channel.GetChannelByName(channelResult.Name)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	// Update type & updated_at
	currentChannelResult.Type = channelResult.Type
	currentChannelResult.UpdatedAt = time.Now().Unix()

	ok, err := channel.UpdateChannelByName(currentChannelResult)

	if !ok || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}
