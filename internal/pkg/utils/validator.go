// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package utils

import (
	"regexp"
	"strings"
)

// Validator util
type Validator struct {
}

// IsIn checks if item in array
func (v *Validator) IsIn(item string, list []string) bool {
	for _, a := range list {
		if a == item {
			return true
		}
	}
	return false
}

// IsSlug checks if string is a valid slug
func (v *Validator) IsSlug(slug string, min int, max int) bool {

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

// IsEmpty checks if item is empty
func (v *Validator) IsEmpty(item string) bool {
	if strings.TrimSpace(item) == "" {
		return true
	}
	return false
}
