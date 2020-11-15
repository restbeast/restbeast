package cmds

import (
	"fmt"
	"github.com/restbeast/restbeast/lib"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "restbeast",
	Short: "RestBeast is an api testing tool.",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
		os.Exit(0)
	},
}

var execCtx *lib.ExecutionContext

func Execute(ctx *lib.ExecutionContext) {
	execCtx = ctx

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
