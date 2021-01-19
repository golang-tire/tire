package embed

import (
	"github.com/golang-tire/tire/internal/cmd/util"
	"github.com/golang-tire/tire/internal/swagger"

	"github.com/spf13/cobra"
)

// SwaggerOptions defines flags and other configuration parameters for the `swagger` command
type SwaggerOptions struct {
	JsonFile      string
	OutputFile    string
	OutputPackage string
	cmdBaseName   string
}

var (
	swaggerLong    = ``
	swaggerExample = ``
)

// NewSwaggerOption creates new ApplyOptions for the `swagger` command
func NewSwaggerOption() *SwaggerOptions {
	return &SwaggerOptions{
		JsonFile:      "",
		OutputFile:    "Stdout",
		OutputPackage: "",
		cmdBaseName:   "",
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
	cmd.Flags().StringVarP(&o.JsonFile, "json-file", "i", o.JsonFile, "path to swagger json file")
	cmd.Flags().StringVarP(&o.OutputFile, "output-file", "o", o.OutputFile, "Full path to output golang file or 'Stdout'")
	cmd.Flags().StringVarP(&o.OutputPackage, "output-pkg", "p", o.OutputPackage, "Output package name")
	return cmd
}

func (o *SwaggerOptions) Complete(cmd *cobra.Command) error {
	var err error
	return err
}

func (o *SwaggerOptions) Run() error {
	return swagger.EmbedSwagger(o.JsonFile, o.OutputFile, o.OutputPackage)
}
