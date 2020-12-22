package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	versionExample = `
		# Print the tire cli version
		tire version`
)

// Options is a struct to support version command
type Options struct {
	Short bool
}

// NewCmdVersion returns a cobra command for printing versions
func NewCmdVersion() *cobra.Command {
	o := &Options{Short: false}
	cmd := &cobra.Command{
		Use:     "version",
		Short:   "Print the tire cli version",
		Long:    "Print the tire cli version",
		Example: versionExample,
		Run:     o.run,
	}
	cmd.Flags().BoolVar(&o.Short, "short", o.Short, "If true, print just the version number.")
	return cmd
}

func (o *Options) run(cmd *cobra.Command, args []string) {
	fmt.Println("Tire cli version:")
}
