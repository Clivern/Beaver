// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package utils

import (
	"regexp"
)

// Validator util
type Validator struct {
}

// In checks if item in array
func (v *Validator) In(item string, list []string) bool {
	for _, a := range list {
		if a == item {
			return true
		}
	}
	return false
}

// Slug checks if string is a valid slug
func (v *Validator) Slug(slug string, min int, max int) bool {

	if len(slug) < min {
		return false
	}

	if len(slug) > max {
		return false
	}

	if regexp.MustCompile(`^[a-z0-9]+(?:_[a-z0-9]+)*$`).MatchString(slug) {
		return true
	}
	return false
}
