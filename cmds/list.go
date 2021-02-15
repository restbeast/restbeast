package cmds

import (
	. "fmt"
	"github.com/nathan-fiscaletti/consolesize-go"
	"github.com/restbeast/restbeast/lib"
	"github.com/spf13/cobra"
	"os"
)

var longFormat, showOnlyRequests, showOnlyTests bool

func init() {
	listCmd.Flags().BoolVarP(&longFormat, "long", "l", false, "Use long listing format")
	listCmd.Flags().BoolVarP(&showOnlyRequests, "requests", "r", false, "List only requests")
	listCmd.Flags().BoolVarP(&showOnlyTests, "tests", "t", false, "List only tests")
	rootCmd.AddCommand(listCmd)
}

func doList(cmd *cobra.Command, args []string) {
	var lsType string
	if showOnlyRequests {
		lsType = "request"
	} else if showOnlyTests {
		lsType = "test"
	}

	list, maxNameLen, err := lib.ListRequestsAndTests(lsType, execCtx)

	if err != nil {
		Printf("Error: Failed to load given request\n%s\n", err)
		os.Exit(1)
	}

	if longFormat {
		for _, item := range list {
			Printf("%-"+Sprintf("%d", maxNameLen)+"s → %s\n", item.Name, item.Type)
		}
	} else {
		// Get screen width
		screenWidth, _ = consolesize.GetConsoleSize()
		if screenWidth == 0 {
			screenWidth = 30
		}

		// Calculate column count by dividing max width with max name length
		columnCount := screenWidth / (maxNameLen + 2)

		// Calculate row count
		// Int divisions results as floored ints
		// Add 1 if there is a remainder
		rowCount := len(list) / columnCount
		if len(list)%columnCount != 0 {
			rowCount++
		}

		// Loop for each row and each column
		for row := 0; row < rowCount; row++ {
			for c := 0; c < columnCount; c++ {
				currIndex := (c * rowCount) + row

				// There might be short rows
				if len(list) > currIndex {
					// Pad names with spaces to max name length + 2 for column spacing
					Printf("%-"+Sprintf("%d", maxNameLen+2)+"s", list[currIndex].Name)
				}
			}

			// Print a new line at the end of each row
			Printf("\n")
		}
	}

}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List http requests",
	Run:     doList,
	Args:    cobra.MinimumNArgs(0),
}
