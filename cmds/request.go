package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.com/restbeast/cli/lib"
	"golang.org/x/crypto/ssh/terminal"
	"net/http"
	"os"
)

var outputTiming, outputDetailedTiming bool
var env string

func init() {
	requestCmd.Flags().BoolVar(&outputTiming, "timing", false, "Displays timings")
	requestCmd.Flags().BoolVar(&outputDetailedTiming, "detailed-timing", false, "Displays detailed timings")
	requestCmd.Flags().StringVar(&env, "env", "", "Selected environment")

	rootCmd.AddCommand(requestCmd)
}

func doRequest(cmd *cobra.Command, args []string) {
	isTerminal := terminal.IsTerminal(int(os.Stdout.Fd()))

	if len(args) == 0 {
		fmt.Println("Error: Specify a request name")
		os.Exit(1)
	}

	request, err := lib.LoadWhole(args[0], env, execCtx)
	if err != nil {
		fmt.Printf("Error: Failed to load given request\n%s\n", err)
		os.Exit(1)
	}

	response, requestErr := lib.DoRequest(*request, execCtx)
	if requestErr != nil {
		fmt.Printf("Error: Failed to execute request\n%s\n", requestErr)
		os.Exit(1)
	}

	// Check if output is terminal or pipe
	if isTerminal {
		// Print out response information
		fmt.Printf("%s %d %s", response.Proto, response.StatusCode, http.StatusText(response.StatusCode))

		if outputTiming {
			fmt.Printf("\nTotal Time: %s", response.Timing.Total)
		} else if outputDetailedTiming {
			fmt.Printf("\nDNS Resolve Time: %s", response.Timing.Dns)
			fmt.Printf("\nConenction Time: %s", response.Timing.Conn)
			if response.Timing.Tls > 0 {
				fmt.Printf("\nTLS Handshake Time: %s", response.Timing.Tls)
			}
			fmt.Printf("\nFirst Byte Time: %s", response.Timing.FirstByte)
			fmt.Printf("\nTotal Time: %s", response.Timing.Total)
		}

		fmt.Printf("\n\n%s", response.Body)
	} else { // piped output
		fmt.Printf("%s", response.Body)
	}
}

var requestCmd = &cobra.Command{
	Use:     "request",
	Aliases: []string{"r"},
	Short:   "Execute a http request",
	Run:     doRequest,
	Args:    cobra.MinimumNArgs(0),
}
