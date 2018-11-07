// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/clivern/beaver/internal/pkg/driver"
	"github.com/clivern/beaver/internal/pkg/logger"
	"github.com/clivern/beaver/internal/pkg/utils"
)

// Migrate struct
type Migrate struct {
}

// Up runs the up migrations
func (e *Migrate) Up() (error, bool) {

	var query string
	var mysql driver.MySQL = *driver.NewMySQL()

	if mysql.Ping() == false {
		logger.Error("Unable to connect to database!")
	}

	files := []string{}
	files = utils.ListFiles("internal/scheme")
	files = utils.FilterFiles(files, []string{"up"})

	ok := true
	result := true

	fmt.Print("\n\033[35m")
	for _, file := range files {
		query = utils.ReadFile(file)
		result = mysql.Exec(query)

		if result {
			logger.Infof("Migration %s executed successfully!", file)
			fmt.Printf("Migration %s executed successfully!\n", file)
		}

		ok = ok && result
	}
	fmt.Println("\033[0m")

	mysql.Close()

	return nil, ok
}

// SafeUp runs the needed migrations only
func (e *Migrate) SafeUp() (error, bool) {

	var query string
	var mysql driver.MySQL = *driver.NewMySQL()

	if mysql.Ping() == false {
		logger.Error("Unable to connect to database!")
	}

	files := []string{}
	files = utils.ListFiles("internal/scheme")
	files = utils.FilterFiles(files, []string{"up"})

	ok := true
	result := true

	for _, file := range files {
		query = utils.ReadFile(file)
		result = mysql.Exec(query)

		if result {
			logger.Infof("Migration %s executed successfully!", file)
		}

		ok = ok && result
	}

	mysql.Close()

	return nil, ok
}

// Down runs the down migrations
func (e *Migrate) Down() (error, bool) {

	var query string
	var mysql driver.MySQL = *driver.NewMySQL()

	if mysql.Ping() == false {
		logger.Error("Unable to connect to database!")
	}

	files := []string{}
	files = utils.ListFiles("internal/scheme")
	files = utils.FilterFiles(files, []string{"down"})

	ok := true
	result := true

	fmt.Print("\n\033[35m")
	for _, file := range files {
		query = utils.ReadFile(file)
		result = mysql.Exec(query)

		if result {
			logger.Infof("Migration %s executed successfully!", file)
			fmt.Printf("Migration %s executed successfully!\n", file)
		}

		ok = ok && result
	}
	fmt.Println("\033[0m")

	mysql.Close()

	return nil, ok
}

// Status shows migration status
func (e *Migrate) Status() (error, bool) {
	fmt.Println("Status")
	return nil, true
}
