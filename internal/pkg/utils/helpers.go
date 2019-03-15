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
	"strings"
)

// GenerateUUID4 create a UUID
func GenerateUUID4() string {
	u := uuid.Must(uuid.NewV4(), nil)
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
	fmt.Println("\033[32mBeaver - A Real Time Messaging Server.\033[0m")
}

// PrintCommands show the available commands
func PrintCommands() {
	fmt.Print("\033[31m")
	fmt.Println("\nAvailable Commands:")
	fmt.Print("\033[0m")
	fmt.Println("\033[32m  -exec=health\033[0m")
	fmt.Print("\n")
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

// FilterFiles filters files list based on specific sub-strings
func FilterFiles(files, filters []string) []string {
	var filteredFiles []string

	for _, file := range files {
		ok := true
		for _, filter := range filters {

			ok = ok && strings.Contains(file, filter)
		}
		if ok {
			filteredFiles = append(filteredFiles, file)
		}
	}

	return filteredFiles
}

// Unset remove element at position i
func Unset(a []string, i int) []string {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	return a[:len(a)-1]
}
