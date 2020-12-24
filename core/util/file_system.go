// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package service

import (
	"os"
)

// FileSystem struct
type FileSystem struct{}

// PathExists reports whether the path exists
func (fs *FileSystem) PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// FileExists reports whether the named file exists
func (fs *FileSystem) FileExists(path string) bool {
	if fi, err := os.Stat(path); err == nil {
		if fi.Mode().IsRegular() {
			return true
		}
	}
	return false
}

// DirExists reports whether the dir exists
func (fs *FileSystem) DirExists(path string) bool {
	if fi, err := os.Stat(path); err == nil {
		if fi.Mode().IsDir() {
			return true
		}
	}
	return false
}

// EnsureDir ensures that directory exists
func (fs *FileSystem) EnsureDir(dirName string, mode int) (bool, error) {
	err := os.MkdirAll(dirName, os.FileMode(mode))

	if err == nil || os.IsExist(err) {
		return true, nil
	}
	return false, err
}
