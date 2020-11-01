package main

import (
	"fmt"

	"github.com/golang-tire/tire/cmd"

	"github.com/spf13/cobra"
)

var Version string

func init() {
	cmd.RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of cli",
	Long:  `All software has versions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("tire cli %s -- HEAD\n", Version)
	},
}
