// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	AppMode       string `json:"app_mode"`
	AppPort       string `json:"app_port"`
	AppLogLevel   string `json:"app_log_level"`
	AppDomain     string `json:"app_domain"`
	MySQLUsername string `json:"mysql_username"`
	MySQLPassword string `json:"mysql_password"`
	MySQLProtocol string `json:"mysql_protocol"`
	MySQLHost     string `json:"mysql_host"`
	MySQLPort     string `json:"mysql_port"`
	MySQLDatabase string `json:"mysql_database"`
}

func (e *Config) Load(file string) (bool, error) {

	_, err := os.Stat(file)

	if err != nil {
		return false, fmt.Errorf("config file %s not found", file)
	}

	data, err := ioutil.ReadFile(file)

	if err != nil {
		return false, err
	}

	err = json.Unmarshal(data, &e)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (e *Config) Cache() {
	if os.Getenv("AppMode") == "" {
		os.Setenv("AppMode", e.AppMode)
		os.Setenv("AppLogLevel", e.AppLogLevel)
		os.Setenv("AppPort", e.AppPort)
		os.Setenv("AppDomain", e.AppDomain)
		os.Setenv("MySQLUsername", e.MySQLUsername)
		os.Setenv("MySQLPassword", e.MySQLPassword)
		os.Setenv("MySQLProtocol", e.MySQLProtocol)
		os.Setenv("MySQLHost", e.MySQLHost)
		os.Setenv("MySQLPort", e.MySQLPort)
		os.Setenv("MySQLDatabase", e.MySQLDatabase)
	}
}

func (e *Config) GinEnv() {
	// Used by gin framework
	// https://github.com/gin-gonic/gin/blob/d510595aa58c2417373d89a8d8ffa21cf58673cb/utils.go#L140
	os.Setenv("PORT", os.Getenv("AppPort"))
}
