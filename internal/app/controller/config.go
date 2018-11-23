// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"github.com/clivern/beaver/internal/app/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetConfigByKey controller
func GetConfigByKey(c *gin.Context) {
	key := c.Param("key")
	config := api.Config{
		CorrelationID: c.Request.Header.Get("X-Correlation-ID"),
	}

	if !config.Init() {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "error",
			"error":  "Internal server error",
		})
		return
	}

	value, err := config.GetConfigByKey(key)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"key":   key,
		"value": value,
	})
}

// CreateConfig controller
func CreateConfig(c *gin.Context) {

	var configRequest api.ConfigResult

	config := api.Config{
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

	ok, err := configRequest.LoadFromJSON(rawBody)

	if !ok || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid request",
		})
		return
	}

	if configRequest.Key == "" || configRequest.Value == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid request",
		})
		return
	}

	if !config.Init() {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":  "error",
			"message": "Internal server error",
		})
		return
	}

	_, err = config.CreateConfig(configRequest.Key, configRequest.Value)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

// DeleteConfigByKey controller
func DeleteConfigByKey(c *gin.Context) {
	key := c.Param("key")
	config := api.Config{
		CorrelationID: c.Request.Header.Get("X-Correlation-ID"),
	}

	if !config.Init() {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "error",
			"error":  "Internal server error",
		})
		return
	}

	_, err := config.DeleteConfigByKey(key)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

// UpdateConfigByKey controller
func UpdateConfigByKey(c *gin.Context) {

	var configRequest api.ConfigResult

	config := api.Config{
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

	configRequest.LoadFromJSON(rawBody)
	configRequest.Key = c.Param("key")

	if configRequest.Key == "" || configRequest.Value == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid request",
		})
		return
	}

	if !config.Init() {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "error",
			"error":  "Internal server error",
		})
		return
	}

	_, err = config.UpdateConfigByKey(configRequest.Key, configRequest.Value)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}
