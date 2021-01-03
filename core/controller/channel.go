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

// GetChannelByName controller
func GetChannelByName(c *gin.Context) {

	var channelResult module.ChannelResult

	validate := util.Validator{}

	name := c.Param("name")

	if validate.IsEmpty(name) || !validate.IsSlug(name, 3, 60) {
		c.JSON(http.StatusBadRequest, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Channel name must be alphanumeric with length from 3 to 60 and lowercase",
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

	channel := module.NewChannel(db)

	channelResult, err = channel.GetChannelByName(name)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Error! Channel not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":      channelResult.Name,
		"type":      channelResult.Type,
		"createdAt": time.Unix(channelResult.CreatedAt, 0),
		"updatedAt": time.Unix(channelResult.UpdatedAt, 0),
	})
}

// CreateChannel controller
func CreateChannel(c *gin.Context) {

	var channelResult module.ChannelResult

	validate := util.Validator{}

	rawBody, err := c.GetRawData()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Error! Invalid request",
		})
		return
	}

	ok, err := channelResult.LoadFromJSON(rawBody)

	if !ok || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Error! Invalid request",
		})
		return
	}

	if validate.IsEmpty(channelResult.Name) || !validate.IsSlug(channelResult.Name, 3, 60) {
		c.JSON(http.StatusBadRequest, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Channel name must be alphanumeric with length from 3 to 60 and lowercase",
		})
		return
	}

	if !validate.IsIn(channelResult.Type, []string{"public", "private", "presence"}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Channel type must be public, private or presence",
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

	channel := module.NewChannel(db)

	ok, _ = channel.ChannelExist(channelResult.Name)

	if ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  fmt.Sprintf("Error! Channel %s already exists", channelResult.Name),
		})
		return
	}

	channelResult.CreatedAt = time.Now().Unix()
	channelResult.UpdatedAt = time.Now().Unix()

	ok, err = channel.CreateChannel(channelResult)

	if !ok || err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"name":      channelResult.Name,
		"type":      channelResult.Type,
		"createdAt": time.Unix(channelResult.CreatedAt, 0),
		"updatedAt": time.Unix(channelResult.UpdatedAt, 0),
	})
}

// DeleteChannelByName controller
func DeleteChannelByName(c *gin.Context) {

	validate := util.Validator{}

	name := c.Param("name")

	if validate.IsEmpty(name) || !validate.IsSlug(name, 3, 60) {
		c.JSON(http.StatusBadRequest, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Channel name must be alphanumeric with length from 3 to 60 and lowercase",
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

	channel := module.NewChannel(db)

	ok, err := channel.ChannelExist(name)

	if !ok || err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  fmt.Sprintf("Error! Channel %s not found", name),
		})
		return
	}

	_, err = channel.DeleteChannelByName(name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Internal server error",
		})
		return
	}

	c.Status(http.StatusNoContent)
}

// UpdateChannelByName controller
func UpdateChannelByName(c *gin.Context) {

	var channelResult module.ChannelResult
	var currentChannelResult module.ChannelResult

	validate := util.Validator{}

	rawBody, err := c.GetRawData()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Error! Invalid request",
		})
		return
	}

	channelResult.LoadFromJSON(rawBody)
	channelResult.Name = c.Param("name")

	if validate.IsEmpty(channelResult.Name) || !validate.IsSlug(channelResult.Name, 3, 60) {
		c.JSON(http.StatusBadRequest, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Channel name must be alphanumeric with length from 3 to 60 and lowercase",
		})
		return
	}

	if !validate.IsIn(channelResult.Type, []string{"public", "private", "presence"}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Channel type must be public, private or presence",
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

	channel := module.NewChannel(db)

	ok, err := channel.ChannelExist(channelResult.Name)

	if !ok || err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  fmt.Sprintf("Error! Channel %s not found", channelResult.Name),
		})
		return
	}

	currentChannelResult, err = channel.GetChannelByName(channelResult.Name)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  fmt.Sprintf("Error! Channel %s not found", channelResult.Name),
		})
		return
	}

	// Update type & updated_at
	currentChannelResult.Type = channelResult.Type
	currentChannelResult.UpdatedAt = time.Now().Unix()

	ok, err = channel.UpdateChannelByName(currentChannelResult)

	if !ok || err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":      channelResult.Name,
		"type":      channelResult.Type,
		"createdAt": time.Unix(channelResult.CreatedAt, 0),
		"updatedAt": time.Unix(channelResult.UpdatedAt, 0),
	})
}

// GetChannelSubscribersByName controller
func GetChannelSubscribersByName(c *gin.Context) {

	validate := util.Validator{}

	name := c.Param("name")

	if validate.IsEmpty(name) || !validate.IsSlug(name, 3, 60) {
		c.JSON(http.StatusBadRequest, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Channel name must be alphanumeric with length from 3 to 60 and lowercase",
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

	channel := module.NewChannel(db)

	ok, err := channel.ChannelExist(name)

	if !ok || err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  fmt.Sprintf("Error! Channel %s not found", name),
		})
		return
	}

	subscribers, err := channel.GetSubscribers(name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Internal server error",
		})
		return
	}

	count, err := channel.CountSubscribers(name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"subscribers": subscribers,
		"count":       count,
	})
}

// GetChannelListenersByName controller
func GetChannelListenersByName(c *gin.Context) {

	validate := util.Validator{}

	name := c.Param("name")

	if validate.IsEmpty(name) || !validate.IsSlug(name, 3, 60) {
		c.JSON(http.StatusBadRequest, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Channel name must be alphanumeric with length from 3 to 60 and lowercase",
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

	channel := module.NewChannel(db)

	ok, err := channel.ChannelExist(name)

	if !ok || err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  fmt.Sprintf("Error! Channel %s not found", name),
		})
		return
	}

	listeners, err := channel.GetListeners(name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Internal server error",
		})
		return
	}

	count, err := channel.CountListeners(name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"correlationID": c.Request.Header.Get("X-Correlation-ID"),
			"errorMessage":  "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"listeners": listeners,
		"count":     count,
	})
}
