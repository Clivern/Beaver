// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Metrics controller
func Metrics() http.Handler {
	return promhttp.Handler()
}
