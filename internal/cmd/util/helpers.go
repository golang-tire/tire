package util

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

const (
	DefaultErrorExitCode = 1
)

var (
	fatalErrHandler = fatal
	ErrExit         = fmt.Errorf("exit")
)

// AddDryRunFlag adds dry-run flag to a command. Usually used by mutations.
func AddDryRunFlag(cmd *cobra.Command) {
	cmd.Flags().String(
		"dry-run",
		"false",
		`Must be "true" or "false". if command will run without persisting`,
	)
	cmd.Flags().Lookup("dry-run").NoOptDefVal = "unchanged"
}

func GetDryRun(cmd *cobra.Command) (bool, error) {
	var dryRunFlag = GetFlagString(cmd, "dry-run")
	b, err := strconv.ParseBool(dryRunFlag)
	if err != nil {
		return false, fmt.Errorf(`invalid dry-run value (%v). must be "true" or "false"`, dryRunFlag)
	}
	return b, err
}

func GetFlagString(cmd *cobra.Command, flag string) string {
	s, err := cmd.Flags().GetString(flag)
	if err != nil {
		log.Fatalf("error accessing flag %s for command %s: %v", flag, cmd.Name(), err)
	}
	return s
}

func GetFlagBool(cmd *cobra.Command, flag string) bool {
	b, err := cmd.Flags().GetBool(flag)
	if err != nil {
		log.Fatalf("error accessing flag %s for command %s: %v", flag, cmd.Name(), err)
	}
	return b
}

func CheckErr(err error) {
	checkErr(err, fatalErrHandler)
}

func fatal(msg string, code int) {
	if len(msg) > 0 {
		// add newline if needed
		if !strings.HasSuffix(msg, "\n") {
			msg += "\n"
		}
		fmt.Fprint(os.Stderr, msg)
	}
	os.Exit(code)
}

func checkErr(err error, handleErr func(string, int)) {
	if err == nil {
		return
	}
	handleErr(err.Error(), DefaultErrorExitCode)
}
