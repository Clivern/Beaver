// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

// Schema type
type Schema struct {
	FilePath     string
	DatabaseName string
}

// NewSchema returns schema instance
func NewSchema() *Schema {
	return &Schema{}
}

// WithFilePath define the file path
func (s *Schema) WithFilePath(filePath string) *Schema {
	s.FilePath = filePath
	return s
}

// WithDatabaseName define database name
func (s *Schema) WithDatabaseName(databaseName string) *Schema {
	s.DatabaseName = databaseName
	return s
}

// GetQueries gets the queries to run
func (s *Schema) GetQueries() []string {
	return []string{}
}
