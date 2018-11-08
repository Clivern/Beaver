// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/clivern/beaver/internal/model"
	"github.com/clivern/beaver/internal/pkg/driver"
	"github.com/clivern/beaver/internal/pkg/logger"
	"github.com/clivern/beaver/internal/pkg/utils"
	"strings"
)

// CreateMigrationTable creates migration table
func CreateMigrationTable() (bool, error) {
	var query string
	var mysql driver.MySQL = *driver.NewMySQL()

	if mysql.Ping() == false {
		logger.Error("Unable to connect to database!")
	}

	files := utils.ListFiles("internal/scheme")
	upMigrationFile := utils.FilterFiles(files, []string{"migration_up"})
	ok := true
	result := true

	if mysql.TableExists("migration") {
		return true, nil
	}

	fmt.Print("\n\033[35m")
	for _, file := range upMigrationFile {
		query = utils.ReadFile(file)
		result = mysql.Exec(query)
		if result {
			logger.Infof("Migration %s executed successfully!", file)
			fmt.Printf("Migration %s executed successfully!\n", file)
		}
		ok = ok && result
	}
	fmt.Print("\033[0m")

	mysql.Close()

	return ok, nil
}

// DropMigrationTable deletes migrations table
func DropMigrationTable() (bool, error) {
	var query string
	var mysql driver.MySQL = *driver.NewMySQL()

	if mysql.Ping() == false {
		logger.Error("Unable to connect to database!")
	}

	files := utils.ListFiles("internal/scheme")
	downMigrationFile := utils.FilterFiles(files, []string{"migration_down"})
	ok := true
	result := true

	if !mysql.TableExists("migration") {
		return true, nil
	}

	fmt.Print("\n\033[35m")
	for _, file := range downMigrationFile {
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

	return ok, nil
}

// Up runs up migrations
func Up() (bool, error) {
	var query string
	var mysql driver.MySQL = *driver.NewMySQL()

	if mysql.Ping() == false {
		logger.Error("Unable to connect to database!")
	}

	files := utils.ListFiles("internal/scheme")
	upFiles := utils.FilterFiles(files, []string{".up"})
	ok := true
	result := true

	fmt.Print("\n\033[35m")
	for _, file := range upFiles {
		count, _ := model.GetOneByMigration(mysql.Connection, file)
		if count == 0 {
			query = utils.ReadFile(file)
			result = mysql.Exec(query)
			if result {
				model.InsertOne(mysql.Connection, file)
				logger.Infof("Migration %s executed successfully!", file)
				fmt.Printf("Migration %s executed successfully!\n", file)
			}
			ok = ok && result
		}
	}
	fmt.Println("\033[0m")

	mysql.Close()

	return ok, nil
}

// Down runs down migrations
func Down() (bool, error) {
	var query string
	var mysql driver.MySQL = *driver.NewMySQL()

	if mysql.Ping() == false {
		logger.Error("Unable to connect to database!")
	}

	files := utils.ListFiles("internal/scheme")
	downFiles := utils.FilterFiles(files, []string{".down"})
	ok := true
	result := true

	fmt.Print("\n\033[35m")
	for _, file := range downFiles {
		count, _ := model.GetOneByMigration(mysql.Connection, strings.Replace(file, ".down.", ".up.", 1))
		if count != 0 {
			query = utils.ReadFile(file)
			result = mysql.Exec(query)
			if result {
				model.DeleteOneByMigration(mysql.Connection, strings.Replace(file, ".down.", ".up.", 1))
				logger.Infof("Migration %s executed successfully!", file)
				fmt.Printf("Migration %s executed successfully!\n", file)
			}
			ok = ok && result
		}
	}
	fmt.Println("\033[0m")

	mysql.Close()

	return ok, nil
}

// Status shows migration status
func Status() (bool, error) {
	fmt.Println("Status")
	return true, nil
}
