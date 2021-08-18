// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Version buildinfo item
	Version = "dev"
	// Commit buildinfo item
	Commit = "none"
	// Date buildinfo item
	Date = "unknown"
	// BuiltBy buildinfo item
	BuiltBy = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(
			fmt.Sprintf(
				`Current Peanut Version %v Commit %v, Built @%v By %v.`,
				Version,
				Commit,
				Date,
				BuiltBy,
			),
		)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
