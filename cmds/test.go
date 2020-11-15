package cmds

import (
	. "fmt"
	"github.com/restbeast/restbeast/lib"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
	"os"
)

var green = chalk.Green.NewStyle().Style
var red = chalk.Red.NewStyle().Style

func init() {
	testCmd.Flags().StringVar(&env, "env", "", "Selected environment")
	rootCmd.AddCommand(testCmd)
}

func doTest(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		Println("Error: Specify a test name")
		os.Exit(1)
	}

	test, err := lib.LoadTest(args[0], env, execCtx)
	if err != nil {
		Printf("Error: Failed to load given request\n%s\n", err)
		os.Exit(1)
	}

	exitCode := 0
	for _, result := range test.Assertions {
		if result.Pass {
			Printf("%s: %s\n", green("PASS"), result.Name)
		} else {
			exitCode = 1
			Printf("%s: %s\n%s\n------------\n", red("FAIL"), result.Name, result.Text)
		}
	}

	os.Exit(exitCode)
}

var testCmd = &cobra.Command{
	Use:     "test",
	Aliases: []string{"t"},
	Short:   "Execute a test suite",
	Run:     doTest,
	Args:    cobra.MinimumNArgs(0),
}
