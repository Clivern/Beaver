// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package driver

import (
	"database/sql"
	"fmt"
	"github.com/clivern/beaver/internal/pkg/logger"
	_ "github.com/go-sql-driver/mysql"
)

// MySQL driver
type MySQL struct {
	Username   string
	Password   string
	Host       string
	Port       int
	Database   string
	Protocol   string
	Connection *sql.DB
}

// Ping method check the db connection
func (e *MySQL) Ping() bool {

	var err error

	e.Connection, err = sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@%s(%s:%d)/%s", e.Username, e.Password, e.Protocol, e.Host, e.Port, e.Database),
	)

	if err != nil {
		logger.Errorf("Error connecting to DB: %s", err.Error())
		return false
	}

	err = e.Connection.Ping()

	if err != nil {
		logger.Errorf("Error while checking DB connection: %s", err.Error())
		return false
	}

	return true
}

// Close method closes the db connection
func (e *MySQL) Close() {
	e.Connection.Close()
}

// Exec run a query
func (e *MySQL) Exec(query string) bool {
	_, err := e.Connection.Exec(query)

	if err == nil {
		return true
	}
	return false
}

// Migrate runs migrations inside specifc path
func (e *MySQL) Migrate(path string, direction string) bool {
	return true
}

//func (e *MySQL) ScanMigration(path string, direction string) []string {

//}

// TableExists checks if table exists or not
func (e *MySQL) TableExists(tableName string) bool {

	var count int

	rows, err := e.Connection.Query(
		"SELECT count(*) as count FROM information_schema.tables WHERE table_schema = ? AND table_name = ?",
		e.Database,
		tableName,
	)

	if err != nil {
		return false
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return false
		}
		if count == 1 {
			return true
		}
		return false
	}

	err = rows.Err()

	if err != nil {
		return false
	}
	return false
}
