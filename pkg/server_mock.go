// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package pkg

import (
	"net/http"
	"net/http/httptest"
)

// ServerMock mocks http server
func ServerMock(uri, response string, statusCode int) *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Write([]byte(response))
	})

	srv := httptest.NewServer(handler)

	return srv
}
