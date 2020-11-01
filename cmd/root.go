package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "tire",
	Short: "golang-tire cli",
	Long:  `an small tools to generate project parts`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		_, err = fmt.Fprintln(os.Stderr, err)
		if err!=nil{
			panic(err)
		}
		os.Exit(1)
	}
}
