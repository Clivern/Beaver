// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package util

import (
	"testing"

	"github.com/clivern/beaver/pkg"
)

// TestInArray test cases
func TestInArray(t *testing.T) {
	// TestInArray
	t.Run("TestInArray", func(t *testing.T) {
		pkg.Expect(t, InArray("A", []string{"A", "B", "C", "D"}), true)
		pkg.Expect(t, InArray("B", []string{"A", "B", "C", "D"}), true)
		pkg.Expect(t, InArray("H", []string{"A", "B", "C", "D"}), false)
		pkg.Expect(t, InArray(1, []int{2, 3, 1}), true)
		pkg.Expect(t, InArray(9, []int{2, 3, 1}), false)
	})
}

// TestEncryption test cases
func TestEncryption(t *testing.T) {
	// TestEncryption
	t.Run("TestEncryption", func(t *testing.T) {
		ciphertext, err := Encrypt([]byte("Hello World"), "password")
		plaintext, err := Decrypt(ciphertext, "password")
		pkg.Expect(t, "Hello World", string(plaintext))
		pkg.Expect(t, err, nil)
	})
}

// TestHashing test cases
func TestHashing(t *testing.T) {
	// TestHashing
	t.Run("TestHashing", func(t *testing.T) {
		hash, err := HashPassword("password")
		pkg.Expect(t, true, CheckPasswordHash("password", hash))
		pkg.Expect(t, false, CheckPasswordHash("PasSword", hash))
		pkg.Expect(t, err, nil)
	})
}
