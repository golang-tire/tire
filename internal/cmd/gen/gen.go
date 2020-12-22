package gen

import (
	"github.com/golang-tire/tire/internal/cmd/util"

	"github.com/spf13/cobra"
)

const genBaseName = "gen"

// GenOptions defines flags and other configuration parameters for the `gen` command
type GenOptions struct {
	cmdBaseName string
}

var (
	genLong    = ``
	genExample = ``
)

// NewGenOption creates new ApplyOptions for the `gen` command
func NewGenOption() *GenOptions {
	return &GenOptions{
		cmdBaseName: "",
	}
}

func NewCmdGen(baseName string) *cobra.Command {
	o := NewGenOption()
	o.cmdBaseName = baseName

	cmd := &cobra.Command{
		Use:                   "gen [type]",
		DisableFlagsInUseLine: true,
		Short:                 "generate codes based on selected type",
		Long:                  genLong,
		Example:               genExample,
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Run())
		},
	}

	util.AddDryRunFlag(cmd)
	cmd.AddCommand(NewCmdSwagger(genBaseName))
	return cmd
}

func (o *GenOptions) Run() error {
	return nil
}
