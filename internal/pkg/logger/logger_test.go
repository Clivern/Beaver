// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package logger

import (
	"fmt"
	"github.com/clivern/beaver/internal/pkg/utils"
	"github.com/nbio/st"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

// TestLogging test cases
func TestLogging(t *testing.T) {
	// Setup Env Vars
	basePath := fmt.Sprintf("%s/src/github.com/clivern/beaver", os.Getenv("GOPATH"))
	configFile := fmt.Sprintf("%s/%s", basePath, "config.test.json")

	config := &utils.Config{}
	ok, err := config.Load(configFile)

	if !ok || err != nil {
		panic(err.Error())
	}
	config.Cache()
	config.GinEnv()
	os.Setenv("LogPath", fmt.Sprintf("%s/%s", basePath, os.Getenv("LogPath")))

	// Start Test Cases
	Info("Info")
	Infoln("Infoln")
	Infof("Infof")
	Warning("Warning")
	Warningln("Warningln")
	Warningf("Warningf")

	currentTime := time.Now().Local()
	data, err := ioutil.ReadFile(fmt.Sprintf("%s/%s.log", os.Getenv("LogPath"), currentTime.Format("2006-01-02")))

	if err != nil {
		panic(err.Error())
	}

	st.Expect(t, strings.Contains(string(data), "Info\n"), true)
	st.Expect(t, strings.Contains(string(data), "Infoln\n"), true)
	st.Expect(t, strings.Contains(string(data), "Infof\n"), true)
	st.Expect(t, strings.Contains(string(data), "Warning\n"), true)
	st.Expect(t, strings.Contains(string(data), "Warningln\n"), true)
	st.Expect(t, strings.Contains(string(data), "Warningf\n"), true)

	os.Remove(fmt.Sprintf("%s/%s.log", os.Getenv("LogPath"), currentTime.Format("2006-01-02")))
}
