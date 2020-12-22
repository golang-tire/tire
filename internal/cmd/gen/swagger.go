package gen

import (
	"github.com/golang-tire/tire/internal/cmd/util"

	"github.com/spf13/cobra"
)

// SwaggerOptions defines flags and other configuration parameters for the `swagger` command
type SwaggerOptions struct {
	DryRun      bool
	cmdBaseName string
}

var (
	swaggerLong    = ``
	swaggerExample = ``
)

// NewSwaggerOption creates new ApplyOptions for the `swagger` command
func NewSwaggerOption() *SwaggerOptions {
	return &SwaggerOptions{
		DryRun:      false,
		cmdBaseName: "",
	}
}

func NewCmdSwagger(baseName string) *cobra.Command {
	o := NewSwaggerOption()
	o.cmdBaseName = baseName

	cmd := &cobra.Command{
		Use:                   "swagger",
		DisableFlagsInUseLine: true,
		Short:                 "generate swagger bindings from swagger json files",
		Long:                  swaggerLong,
		Example:               swaggerExample,
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Complete(cmd))
			util.CheckErr(o.Run())
		},
	}
	return cmd
}

func (o *SwaggerOptions) Complete(cmd *cobra.Command) error {
	var err error
	o.DryRun, err = util.GetDryRun(cmd)
	return err
}

func (o *SwaggerOptions) Run() error {
	return nil
}
