package cmds

import (
	. "fmt"
	"github.com/dustin/go-humanize"
	"github.com/nathan-fiscaletti/consolesize-go"
	"github.com/restbeast/restbeast/lib"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"net/http"
	"os"
)

var outputTiming, outputDetailedTiming, showHeaders bool
var env string
var screenWidth int
var repeat int

func init() {
	requestCmd.Flags().BoolVarP(&showHeaders, "headers", "H", false, "Show response headers")
	requestCmd.Flags().BoolVar(&outputTiming, "timing", false, "Displays timings")
	requestCmd.Flags().BoolVar(&outputDetailedTiming, "detailed-timing", false, "Displays detailed timings")
	requestCmd.Flags().StringVar(&env, "env", "", "Selected environment")
	requestCmd.Flags().IntVar(&repeat, "repeat", 1, "Repeat this request X times")

	rootCmd.AddCommand(requestCmd)
	screenWidth, _ = consolesize.GetConsoleSize()
	if screenWidth == 0 {
		screenWidth = 30
	}
}

func printJustTiming(response lib.Response, padding string) string {
	var returnVal string
	var extraPadding string
	if len(padding) > 0 {
		extraPadding = "  "
	} else {
		padding = "  "
	}

	urlLine := Sprintf("%s %s\n", response.Method, response.Url)
	if len(urlLine) > screenWidth && screenWidth-len(padding)-4 < len(urlLine) {
		returnVal += urlLine[:screenWidth-len(padding)-4] + "...\n"
	} else {
		returnVal += urlLine
	}
	returnVal += Sprintf("%s%s│  Total Time: %d ms\n", padding, extraPadding, response.Timing.Total.Milliseconds())
	if response.BytesSend > 0 {
		returnVal += Sprintf("%s%s│  Bytes Sent: %s\n", padding, extraPadding, humanize.Bytes(response.BytesSend))
	}
	returnVal += Sprintf("%s%s│  Bytes Received: %s\n", padding, extraPadding, humanize.Bytes(response.BytesReceived))

	return returnVal
}

func printDetailedTiming(response lib.Response, padding string) string {
	var returnVal string
	var extraPadding string
	if len(padding) > 0 {
		extraPadding = "  "
	} else {
		padding = "  "
	}

	urlLine := Sprintf("%s %s\n", response.Method, response.Url)
	if len(urlLine) > screenWidth && screenWidth-len(padding)-4 < len(urlLine) {
		returnVal += urlLine[:screenWidth-len(padding)-4] + "...\n"
	} else {
		returnVal += urlLine
	}
	returnVal += Sprintf("%s%s│  DNS Resolve Time: %d ms\n", padding, extraPadding, response.Timing.Dns.Milliseconds())
	returnVal += Sprintf("%s%s│  Connection Time: %d ms\n", padding, extraPadding, response.Timing.Conn.Milliseconds())
	if response.Timing.Tls > 0 {
		returnVal += Sprintf("%s%s│  TLS Handshake Time: %d ms\n", padding, extraPadding, response.Timing.Tls.Milliseconds())
	}
	returnVal += Sprintf("%s%s│  First Byte Time: %d ms\n", padding, extraPadding, response.Timing.FirstByte.Milliseconds())
	returnVal += Sprintf("%s%s│  Total Time: %d ms\n", padding, extraPadding, response.Timing.Total.Milliseconds())
	if response.BytesSend > 0 {
		returnVal += Sprintf("%s%s│  Bytes Sent: %s\n", padding, extraPadding, humanize.Bytes(response.BytesSend))
	}
	returnVal += Sprintf("%s%s│  Bytes Received: %s\n", padding, extraPadding, humanize.Bytes(response.BytesReceived))

	return returnVal
}

func printTiming(outputTiming bool, outputDetailedTiming bool, request lib.Request, response lib.Response, padding string) string {
	var returnVal string

	if outputTiming {
		returnVal += Sprintf("%s  │\n", padding)
		returnVal += Sprintf("%s  ├──", padding)
		returnVal += printJustTiming(response, padding)
	} else if outputDetailedTiming {
		returnVal += Sprintf("%s  │\n", padding)
		returnVal += Sprintf("%s  ├──", padding)
		returnVal += printDetailedTiming(response, padding)
	}

	if len(request.PrecedingRequests) > 0 {
		for _, resP := range request.PrecedingRequests {
			res := *resP
			returnVal += printTiming(outputTiming, outputDetailedTiming, *res.Request, res, padding+"  │  ")
		}
	}

	return returnVal
}

func printHeaders(response lib.Response) string {
	var returnVal string

	if showHeaders {
		returnVal += Sprintf("\n")
		response.Headers.OrderedCallBack(func(k, v string) {
			returnVal += Sprintf("\033[1m%s\033[0m: %s\n", k, v)
		})
	}

	return returnVal
}

func doRequest(cmd *cobra.Command, args []string) {
	isTerminal := terminal.IsTerminal(int(os.Stdout.Fd()))

	if len(args) == 0 {
		Println("Error: Specify a request name")
		os.Exit(1)
	}

	evalCtx, err := lib.LoadEvalCtx(env, execCtx)
	if err != nil {
		Printf("Error: Failed to load given request\n%s\n", err)
		os.Exit(1)
	}

	repeatCount, err := lib.LoadRepeatCount(args[0], evalCtx)
	if err != nil {
		Printf("Error: Failed to load given request\n%s\n", err)
		os.Exit(1)
	}

	if repeat == 1 {
		if repeatCount > 0 {
			repeat = repeatCount
		}
	}

	for i := 0; i < repeat; i++ {
		request, err := lib.LoadOnlyRequest(args[0], evalCtx, execCtx)
		if err != nil {
			Printf("Error: Failed to load given request\n%s\n", err)
			os.Exit(1)
		}

		requestErr := request.Exec()
		if requestErr != nil {
			Printf("Error: Failed to execute request\n%s\n", requestErr)
			os.Exit(1)
		}

		// Check if output is terminal or pipe
		if isTerminal {
			// Print out response information
			if i > 0 {
				Print("\n")
			}

			Printf("%s %d %s\n", request.Response.Proto, request.Response.StatusCode, http.StatusText(request.Response.StatusCode))
			Print(printTiming(outputTiming, outputDetailedTiming, *request, *request.Response, ""))
			Print(printHeaders(*request.Response))

			if len(request.Response.Body) > 0 {
				Printf("\n\n%s", request.Response.Body)
			}
		} else { // piped output
			Printf("%s", request.Response.Body)
		}
	}
}

var requestCmd = &cobra.Command{
	Use:     "request",
	Aliases: []string{"r"},
	Short:   "Execute a http request",
	Run:     doRequest,
	Args:    cobra.MinimumNArgs(0),
}
