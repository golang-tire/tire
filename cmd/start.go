package cmd

import (
	"strings"

	"github.com/golang-tire/tire/internal/engine"
	"github.com/spf13/cobra"
)

func newService(serviceName string) {

	n := strings.ToLower(serviceName + "s")
	data := engine.Data{
		Pkg: engine.Package{
			Name:        serviceName,
			NamePlural:  serviceName + "s",
			EntityAlias: string(n[0]),
		},
	}

	tmpls := []struct {
		Tmpl     string
		Dest     string
		Filename string
	}{

	}

	var err error
	for _, tmpl := range tmpls {
		err = engine.Render(
			tmpl.Tmpl,
			tmpl.Dest,
			tmpl.Filename,
			data,
		)
		if err != nil {
			panic(err)
		}
	}
}

var startCmd = &cobra.Command{
	Use:   "new-service [Service name]",
	Short: "create a service",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		newService(args[0])
	},
}

func init() {
	RootCmd.AddCommand(startCmd)
}
