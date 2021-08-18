// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package util

import (
	"os"

	"github.com/satori/go.uuid"
)

// GenerateUUID4 create a UUID
func GenerateUUID4() string {
	u := uuid.Must(uuid.NewV4(), nil)
	return u.String()
}

// Unset remove element at position i
func Unset(a []string, i int) []string {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	return a[:len(a)-1]
}

// FileExists reports whether the named file exists
func FileExists(path string) bool {
	if fi, err := os.Stat(path); err == nil {
		if fi.Mode().IsRegular() {
			return true
		}
	}
	return false
}

// DirExists reports whether the dir exists
func DirExists(path string) bool {
	if fi, err := os.Stat(path); err == nil {
		if fi.Mode().IsDir() {
			return true
		}
	}
	return false
}

// EnsureDir ensures that directory exists
func EnsureDir(dirName string, mode int) (bool, error) {
	err := os.MkdirAll(dirName, os.FileMode(mode))

	if err == nil || os.IsExist(err) {
		return true, nil
	}
	return false, err
}
