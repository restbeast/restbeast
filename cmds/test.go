package cmds

import (
	. "fmt"
	"github.com/restbeast/restbeast/lib"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
	"os"
	"strings"
)

var green = chalk.Green.NewStyle().Style
var red = chalk.Red.NewStyle().Style
var bold = chalk.White.NewStyle().WithTextStyle(chalk.Bold).Style

type FinalTestResult struct {
	OutputText   string
	SuccessCount int
	FailCount    int
	TotalCount   int
	ExitCode     int
}

func init() {
	testCmd.Flags().StringVar(&env, "env", "", "Selected environment")
	rootCmd.AddCommand(testCmd)
}

func resultOutput(test *lib.Test, padding string) (out FinalTestResult) {
	for _, result := range test.Assertions {
		if result.Pass {
			out.OutputText += Sprintf("%s%s: %s\n", padding, green("PASS"), result.Name)
			out.SuccessCount++
		} else {
			out.ExitCode = 1
			failText := strings.Replace(result.Text, "\n", Sprintf("\n%s  ", padding), -1)
			out.OutputText += Sprintf("%s%s: %s\n%s  %s\n", padding, red("FAIL"), result.Name, padding, failText)
			out.FailCount++
		}
	}

	out.TotalCount = len(test.Assertions)

	return out
}

func doTest(cmd *cobra.Command, args []string) {
	exitCode := 0
	var successCount, failCount, totalCount int

	if len(args) == 1 {
		test, err := lib.LoadTest(args[0], env, execCtx)
		if err != nil {
			Printf("Error: Failed to load given test\n%s\n", err)
			os.Exit(1)
		}

		result := resultOutput(test, "")
		successCount = +result.SuccessCount
		failCount = +result.FailCount
		totalCount = +result.TotalCount

		if result.ExitCode == 1 {
			exitCode = 1
		}

		Print(result.OutputText)
	} else if len(args) == 0 {
		tests, err := lib.LoadAllTests(env, execCtx)

		if err != nil {
			Printf("Error: Failed to load given request\n%s\n", err)
			os.Exit(1)
		}

		for _, test := range tests {
			Printf("Running %s\n", bold(test.Name))

			result := resultOutput(test, "  ")
			successCount = +result.SuccessCount
			failCount = +result.FailCount
			totalCount = +result.TotalCount

			if result.ExitCode == 1 {
				exitCode = 1
			}

			Print(result.OutputText)
		}
	}

	Printf("\n%d passes, %d failures, %d total.\n", successCount, failCount, totalCount)
	os.Exit(exitCode)
}

var testCmd = &cobra.Command{
	Use:     "test",
	Aliases: []string{"t"},
	Short:   "Execute one or all tests",
	Run:     doTest,
	Args:    cobra.MinimumNArgs(0),
}
