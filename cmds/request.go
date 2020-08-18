package cmds

import (
	. "fmt"
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

func printJustTiming(response lib.Response, padding string) {
	var extraPadding string
	if len(padding) > 0 {
		extraPadding = "  "
	} else {
		padding = "  "
	}

	Printf("%s %s\n", response.Method, response.Url)
	Printf("%s%s│  Total Time: %d ms\n", padding, extraPadding, response.Timing.Total.Milliseconds())
}

func printDetailedTiming(response lib.Response, padding string) {
	var extraPadding string
	if len(padding) > 0 {
		extraPadding = "  "
	} else {
		padding = "  "
	}

	Printf("%s %s\n", response.Method, response.Url)
	Printf("%s%s│  DNS Resolve Time: %d ms\n", padding, extraPadding, response.Timing.Dns.Milliseconds())
	Printf("%s%s│  Conenction Time: %d ms\n", padding, extraPadding, response.Timing.Conn.Milliseconds())
	if response.Timing.Tls > 0 {
		Printf("%s%s│  TLS Handshake Time: %d ms\n", padding, extraPadding, response.Timing.Tls.Milliseconds())
	}
	Printf("%s%s│  First Byte Time: %d ms\n", padding, extraPadding, response.Timing.FirstByte.Milliseconds())
	Printf("%s%s│  Total Time: %d ms\n", padding, extraPadding, response.Timing.Total.Milliseconds())
}

func printTiming(outputTiming bool, outputDetailedTiming bool, request lib.Request, response lib.Response, padding string) {
	if outputTiming {
		Printf("%s  │\n", padding)
		Printf("%s  ├──", padding)
		printJustTiming(response, padding)
	} else if outputDetailedTiming {
		Printf("%s  │\n", padding)
		Printf("%s  ├──", padding)
		printDetailedTiming(response, padding)
	}

	if len(request.PrecedingRequests) > 0 {
		for _, resP := range request.PrecedingRequests {
			res := *resP
			printTiming(outputTiming, outputDetailedTiming, *res.Request, res, padding + "  │  ")
		}
	}
}

func doRequest(cmd *cobra.Command, args []string) {
	isTerminal := terminal.IsTerminal(int(os.Stdout.Fd()))

	if len(args) == 0 {
		Println("Error: Specify a request name")
		os.Exit(1)
	}

	request, err := lib.LoadWhole(args[0], env, execCtx)
	if err != nil {
		Printf("Error: Failed to load given request\n%s\n", err)
		os.Exit(1)
	}

	response, requestErr := lib.DoRequest(*request, execCtx)
	if requestErr != nil {
		Printf("Error: Failed to execute request\n%s\n", requestErr)
		os.Exit(1)
	}

	// Check if output is terminal or pipe
	if isTerminal {
		// Print out response information
		Printf("%s %d %s\n", response.Proto, response.StatusCode, http.StatusText(response.StatusCode))
		printTiming(outputTiming, outputDetailedTiming, *request, *response, "")

		Printf("\n\n%s", response.Body)
	} else { // piped output
		Printf("%s", response.Body)
	}
}

var requestCmd = &cobra.Command{
	Use:     "request",
	Aliases: []string{"r"},
	Short:   "Execute a http request",
	Run:     doRequest,
	Args:    cobra.MinimumNArgs(0),
}
