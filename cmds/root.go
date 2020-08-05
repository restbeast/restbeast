package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.com/restbeast/cli/lib"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "restbeast",
	Short: "RestBeast is an api testing tool.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

var execCtx *lib.ExecutionContext

func Execute(ctx *lib.ExecutionContext) {
	execCtx = ctx

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
