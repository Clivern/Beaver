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

// MySQL
type MySQL struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
	Protocol string
}

// Ping check the db connection
func (e *MySQL) Ping() bool {

	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@%s(%s:%d)/%s", e.Username, e.Password, e.Protocol, e.Host, e.Port, e.Database),
	)

	if err != nil {
		logger.Errorf("Error connecting to DB: %s", err.Error())
		return false
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		logger.Errorf("Error while checking DB connection: %s", err.Error())
		return false
	}

	return true
}
