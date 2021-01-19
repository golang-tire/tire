package cmd

import (
	"flag"
	"strings"

	"github.com/golang-tire/tire/internal/cmd/embed"

	"github.com/golang-tire/tire/internal/cmd/gen"

	"github.com/golang-tire/tire/internal/cmd/version"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const baseName = "tire"

func NewTireCommand() *cobra.Command {

	cmds := &cobra.Command{
		Use:   "tire",
		Short: "tire generate golang projects",
		Long: `
	tire generate golang projects

    Find more information at:
		https://github.com/golang-tire
`,
		Run: runHelp,
	}

	flags := cmds.PersistentFlags()
	flags.SetNormalizeFunc(wordSepNormalizeFunc)

	addProfilingFlags(flags)
	cmds.PersistentFlags().AddGoFlagSet(flag.CommandLine)

	cmds.AddCommand(gen.NewCmdGen(baseName))
	cmds.AddCommand(embed.NewCmdEmbed(baseName))
	cmds.AddCommand(version.NewCmdVersion())

	return cmds
}

func runHelp(cmd *cobra.Command, args []string) {
	_ = cmd.Help()
}

func wordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		return pflag.NormalizedName(strings.Replace(name, "_", "-", -1))
	}
	return pflag.NormalizedName(name)
}
