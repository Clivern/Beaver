// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home controller
func Home(c *gin.Context) {
	homeTpl := `<!doctype html><html> <head> <meta charset="utf-8"> <meta name="viewport" content="width=device-width, initial-scale=1"> <title>Beaver</title> <link href="https://fonts.googleapis.com/css?family=Nunito:200,600" rel="stylesheet" type="text/css"> <style> html, body { background-color: #fff; color: #636b6f; font-family: 'Nunito', sans-serif; font-weight: 200; height: 100vh; margin: 0; } .full-height { height: 100vh; } .flex-center { align-items: center; display: flex; justify-content: center; } .position-ref { position: relative; } .top-right { position: absolute; right: 10px; top: 18px; } .content { text-align: center; } .title { font-size: 30px; } .links > a { color: #636b6f; padding: 0 25px; font-size: 12px; font-weight: 600; letter-spacing: .1rem; text-decoration: none; text-transform: uppercase; } .m-b-md { margin-bottom: 30px; } </style> </head> <body> <div class="flex-center position-ref full-height"> <div class="content"> <div class="title"> Beaver </div> <p>A Real Time Messaging Server.</p> </div> </div> </body></html>`
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte(homeTpl))
	return
}
