// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package utils

import (
	"fmt"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"os"
	"path/filepath"
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

// ListFiles lists all files inside a dir
func ListFiles(basePath string) []string {
	var files []string

	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if basePath != path && !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return files
	}

	return files
}

// ReadFile get the file content
func ReadFile(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err.Error()
	}
	return string(data)
}
