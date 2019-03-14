// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package logger

import (
	"fmt"
	"github.com/nbio/st"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

// init setup stuff
func init() {
	basePath := fmt.Sprintf("%s/src/github.com/clivern/beaver", os.Getenv("GOPATH"))
	configFile := fmt.Sprintf("%s/%s", basePath, "config.test.yml")

	viper.SetConfigFile(configFile)

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Sprintf(
			"Error while loading config file [%s]: %s",
			configFile,
			err.Error(),
		))
	}

	os.Setenv("BeaverBasePath", fmt.Sprintf("%s/", basePath))
	os.Setenv("PORT", strconv.Itoa(viper.GetInt("app.port")))
}

// TestLogging test cases
func TestLogging(t *testing.T) {

	currentTime := time.Now().Local()

	logFile := fmt.Sprintf(
		"%s%s/%s.log",
		os.Getenv("BeaverBasePath"),
		viper.GetString("log.path"),
		currentTime.Format("2006-01-02"),
	)

	// Start Test Cases
	Info("Info")
	Infoln("Infoln")
	Infof("Infof")
	Warning("Warning")
	Warningln("Warningln")
	Warningf("Warningf")

	data, err := ioutil.ReadFile(logFile)

	if err != nil {
		panic(err.Error())
	}

	st.Expect(t, strings.Contains(string(data), "Info\n"), true)
	st.Expect(t, strings.Contains(string(data), "Infoln\n"), true)
	st.Expect(t, strings.Contains(string(data), "Infof\n"), true)
	st.Expect(t, strings.Contains(string(data), "Warning\n"), true)
	st.Expect(t, strings.Contains(string(data), "Warningln\n"), true)
	st.Expect(t, strings.Contains(string(data), "Warningf\n"), true)

	os.Remove(logFile)
}
