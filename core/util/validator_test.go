// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package utils

import (
	"fmt"
	"github.com/nbio/st"
	"github.com/spf13/viper"
	"os"
	"strconv"
	"testing"
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

// TestValidation test cases
func TestValidation(t *testing.T) {
	validate := Validator{}
	st.Expect(t, validate.IsIn("public", []string{"public", "private"}), true)
	st.Expect(t, validate.IsSlug("customers_chat_0123", 5, 60), true)
	st.Expect(t, validate.IsSlug("customers_chat-0123", 5, 60), false)
	st.Expect(t, validate.IsSlug(" customers_chat_0123", 5, 60), false)
	st.Expect(t, validate.IsSlug("-customers_chat_0123", 5, 60), false)
	st.Expect(t, validate.IsSlug("customers_chat_0123_", 5, 60), false)
	st.Expect(t, validate.IsSlug("cu", 5, 60), false)
	st.Expect(t, validate.IsSlug("cu263hd53t3g363g3g36362gr3", 5, 10), false)
	st.Expect(t, validate.IsEmpty(" "), true)
	st.Expect(t, validate.IsEmpty(" Test \t "), false)
	st.Expect(t, validate.IsEmpty(" Test "), false)
	st.Expect(t, validate.IsEmpty(" \t "), true)

	st.Expect(t, validate.IsJSON(`{"id": "12", "name": "Joe"}`), true)
	st.Expect(t, validate.IsJSON(`"id": "12", "name": "Joe"}`), false)
	st.Expect(t, validate.IsJSON(`{"id": "12" "name": "Joe"}`), false)
	st.Expect(t, validate.IsJSON(`{"id": "12", "name": "Joe}`), false)
}
