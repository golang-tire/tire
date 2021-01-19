package embed

import (
	"github.com/golang-tire/tire/internal/cmd/util"

	"github.com/spf13/cobra"
)

const embedBaseName = "embed"

// EmbedOptions defines flags and other configuration parameters for the `embed` command
type EmbedOptions struct {
	cmdBaseName string
}

var (
	embedLong    = ``
	embedExample = ``
)

// NewEmbedOption creates new ApplyOptions for the `embed` command
func NewEmbedOption() *EmbedOptions {
	return &EmbedOptions{
		cmdBaseName: "",
	}
}

func NewCmdEmbed(baseName string) *cobra.Command {
	o := NewEmbedOption()
	o.cmdBaseName = baseName

	cmd := &cobra.Command{
		Use:                   "embed",
		DisableFlagsInUseLine: true,
		Short:                 "embed static codes into golang files",
		Long:                  embedLong,
		Example:               embedExample,
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Run())
		},
	}

	cmd.AddCommand(NewCmdSwagger(embedBaseName))
	return cmd
}

func (o *EmbedOptions) Run() error {
	return nil
}
