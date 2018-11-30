// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package utils

import (
	"github.com/dgrijalva/jwt-go"
)

// GenerateJWTToken generate a jwt token for frontend
func GenerateJWTToken(data string, timestamp int64, secret string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data":      data,
		"timestamp": timestamp,
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
