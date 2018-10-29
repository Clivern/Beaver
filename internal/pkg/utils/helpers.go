// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package utils

import (
	"fmt"
	"github.com/satori/go.uuid"
)

// GenerateUUID create a UUID
func GenerateUUID() string {
	u := uuid.Must(uuid.NewV4())
	return u.String()
}

// PrintBanner add a banner to app
func PrintBanner() {
	fmt.Print("\033[31m")
	fmt.Print(`
     .-"""-.__     Hey Dude!
    /      ' o'\
 ,-;  '.  :   _c
:_."\._ ) ::-"
       ""m "m
    `)
	fmt.Println("\033[0m")
	fmt.Println("\033[32mWelcome to Beaver - Pusher Server Implementation (https://github.com/clivern/beaver)\033[0m")
	fmt.Println("\033[33mBy: Clivern <hello@clivern.com>\033[0m")
}
