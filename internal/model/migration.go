// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	"database/sql"
	"fmt"
)

// InsertOne inserts a record
func InsertOne(connection *sql.DB, migration string) (int64, error) {

	stmt, err := connection.Prepare("INSERT INTO migration(migration) VALUES(?)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(migration)
	if err != nil {
		return 0, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastId, nil
}

// GetOneByMigration get a record by migration
func GetOneByMigration(connection *sql.DB, migration string) (int64, error) {
	var id int64
	rows, err := connection.Query("select id from migration where migration = ?", migration)
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return 0, err
		}
		return id, nil
	}

	err = rows.Err()

	if err != nil {
		return 0, err
	}
	return 0, nil
}

// DeleteOneByMigration deletes a record by migration
func DeleteOneByMigration(connection *sql.DB, migration string) (bool, error) {
	_, err := connection.Exec(fmt.Sprintf("DELETE FROM migration where migration='%s'", migration))
	if err != nil {
		return false, err
	}
	return true, nil
}
