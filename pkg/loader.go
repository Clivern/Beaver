// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package pkg

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/drone/envsubst"
	"github.com/spf13/viper"
)

// GetBaseDir returns the project base dir
// Base dir identified if dirName found
// This function for testing purposes only
func GetBaseDir(dirName string) string {
	baseDir, _ := os.Getwd()
	cacheDir := fmt.Sprintf("%s/%s", baseDir, dirName)

	for {
		if fi, err := os.Stat(cacheDir); err == nil {
			if fi.Mode().IsDir() {
				return baseDir
			}
		}
		baseDir = filepath.Dir(baseDir)
		cacheDir = fmt.Sprintf("%s/%s", baseDir, dirName)
	}
}

// LoadConfigs load configs for testing purposes using viper
func LoadConfigs(path string) error {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	data1, err := envsubst.EvalEnv(string(data))

	if err != nil {
		return err
	}

	viper.SetConfigType("yaml")

	viper.ReadConfig(bytes.NewBuffer([]byte(data1)))

	viper.SetDefault("app.name", "x-x-x-x")

	return nil
}
