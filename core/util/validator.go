// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package util

import (
	"encoding/json"
	"regexp"
	"strings"
)

const (
	// UUID3 regex expr
	UUID3 string = "^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}$"
	// UUID4 regex expr
	UUID4 string = "^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	// UUID5 regex expr
	UUID5 string = "^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	// UUID regex expr
	UUID string = "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
	// SLUG regex expr
	SLUG string = "^[a-z0-9]+(?:_[a-z0-9]+)*$"
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

	if regexp.MustCompile(SLUG).MatchString(slug) {
		return true
	}

	return false
}

// IsSlugs checks if string is a valid slug
func (v *Validator) IsSlugs(slugs []string, min int, max int) bool {

	for _, slug := range slugs {
		if !v.IsSlug(slug, min, max) {
			return false
		}
	}

	return true
}

// IsEmpty checks if item is empty
func (v *Validator) IsEmpty(item string) bool {
	if strings.TrimSpace(item) == "" {
		return true
	}
	return false
}

// IsUUID validates a UUID
func (v *Validator) IsUUID(uuid string) bool {
	if regexp.MustCompile(UUID).MatchString(uuid) {
		return true
	}

	return false
}

// IsUUID3 validates a UUID3
func (v *Validator) IsUUID3(uuid string) bool {
	if regexp.MustCompile(UUID3).MatchString(uuid) {
		return true
	}

	return false
}

// IsUUID4 validates a UUID4
func (v *Validator) IsUUID4(uuid string) bool {
	if regexp.MustCompile(UUID4).MatchString(uuid) {
		return true
	}

	return false
}

// IsUUID5 validates a UUID5
func (v *Validator) IsUUID5(uuid string) bool {
	if regexp.MustCompile(UUID5).MatchString(uuid) {
		return true
	}

	return false
}

// IsJSON validates a JSON string
func (v *Validator) IsJSON(str string) bool {
	var jsonStr map[string]interface{}
	err := json.Unmarshal([]byte(str), &jsonStr)
	return err == nil
}
