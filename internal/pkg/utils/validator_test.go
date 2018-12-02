// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package utils

import (
	"fmt"
	"github.com/nbio/st"
	"os"
	"strings"
	"testing"
)

// init setup stuff
func init() {
	basePath := fmt.Sprintf("%s/src/github.com/clivern/beaver", os.Getenv("GOPATH"))
	configFile := fmt.Sprintf("%s/%s", basePath, "config.test.json")

	config := Config{}
	ok, err := config.Load(configFile)

	if !ok || err != nil {
		panic(err.Error())
	}
	config.Cache()
	config.GinEnv()
	if !strings.Contains(os.Getenv("LogPath"), basePath) {
		os.Setenv("LogPath", fmt.Sprintf("%s/%s", basePath, os.Getenv("LogPath")))
	}
}

// TestValidation test cases
func TestValidation(t *testing.T) {
	validate := Validator{}
	st.Expect(t, validate.In("public", []string{"public", "private"}), true)
	st.Expect(t, validate.Slug("customers_chat_0123", 5, 60), true)
	st.Expect(t, validate.Slug("customers_chat-0123", 5, 60), false)
	st.Expect(t, validate.Slug(" customers_chat_0123", 5, 60), false)
	st.Expect(t, validate.Slug("-customers_chat_0123", 5, 60), false)
	st.Expect(t, validate.Slug("customers_chat_0123_", 5, 60), false)
	st.Expect(t, validate.Slug("cu", 5, 60), false)
	st.Expect(t, validate.Slug("cu263hd53t3g363g3g36362gr3", 5, 10), false)
}
