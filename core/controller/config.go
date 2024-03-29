// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"net/http"

	"github.com/clivern/beaver/core/api"
	"github.com/clivern/beaver/core/util"

	"github.com/gin-gonic/gin"
)

// GetConfigByKey controller
func GetConfigByKey(c *gin.Context) {
	key := c.Param("key")
	validate := util.Validator{}

	if validate.IsEmpty(key) || !validate.IsSlug(key, 3, 60) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Config key must be alphanumeric with length from 3 to 60",
		})
		return
	}

	config := api.Config{}

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

	validate := util.Validator{}
	var configRequest api.ConfigResult

	config := api.Config{}

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

	if validate.IsEmpty(configRequest.Key) || !validate.IsSlug(configRequest.Key, 3, 60) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Config key must be alphanumeric with length from 3 to 60",
		})
		return
	}

	if validate.IsEmpty(configRequest.Value) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Config value must not be empty",
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
	validate := util.Validator{}

	if validate.IsEmpty(key) || !validate.IsSlug(key, 3, 60) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Config key must be alphanumeric with length from 3 to 60",
		})
		return
	}

	config := api.Config{}

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
	validate := util.Validator{}

	config := api.Config{}

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

	if validate.IsEmpty(configRequest.Key) || !validate.IsSlug(configRequest.Key, 3, 60) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Config key must be alphanumeric with length from 3 to 60",
		})
		return
	}

	if validate.IsEmpty(configRequest.Value) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Config value must not be empty",
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
