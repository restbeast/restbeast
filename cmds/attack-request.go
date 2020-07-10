package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.com/restbeast/cli/lib"
	"os"
	"sort"
	"sync"
	"time"
)

var count int
var period string

type sortByTime []lib.Response

func (a sortByTime) Len() int           { return len(a) }
func (a sortByTime) Less(i, j int) bool { return a[i].Timing.Total < a[j].Timing.Total }
func (a sortByTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func init() {
	attackRequestCmd.Flags().BoolVarP(&outputTiming, "timing", "t", true, "Displays timings")
	attackRequestCmd.Flags().StringVarP(&env, "env", "e", "", "Selected environment")
	attackRequestCmd.Flags().IntVarP(&count, "count", "c", 60, "Request count")
	attackRequestCmd.Flags().StringVarP(&period, "period", "p", "60s", "Period")
	rootCmd.AddCommand(attackRequestCmd)
}

func doAttackRequest(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Error: Specify a request name")
		os.Exit(1)
	}

	duration, err := time.ParseDuration(period)
	if err != nil {
		fmt.Println("Error: Failed to parse duration")
		fmt.Println("Valid time units are  \"s\", \"m\", \"h\"")
		fmt.Println("Examples: \"300s\", \"1.5h\" or \"2h45m\"")
		os.Exit(1)
	}

	perSecond := count / (int(duration) / int(time.Second))

	var requestsMade int
	requestsMade = 0

	var wg sync.WaitGroup

	var request lib.Request
	request, err = lib.LoadOne(args[0], env, version)
	if err != nil {
		fmt.Println("Error: Failed to load request")
		os.Exit(1)
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	var responses []lib.Response

	for {
		<-ticker.C
		for i := 1; i <= perSecond; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				response := lib.DoRequest(request, version)
				responses = append(responses, response)
			}()
		}

		requestsMade = requestsMade + perSecond
		if requestsMade >= count {
			break
		}
	}

	wg.Wait()

	sort.Stable(sortByTime(responses))

	var totalTime time.Duration
	statusCodes := make(map[int]int)

	for i := 0; i < len(responses); i++ {
		totalTime = totalTime + responses[i].Timing.Total

		if val, ok := statusCodes[responses[i].StatusCode]; ok {
			statusCodes[responses[i].StatusCode] = val + 1
		} else {
			statusCodes[responses[i].StatusCode] = 1
		}
	}

	p95 := len(responses) * 95 / 100
	p99 := len(responses) * 99 / 100
	averageTime := time.Duration(float64(totalTime) / float64(count))

	keys := make([]int, 0, len(statusCodes))
	for k := range statusCodes {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for i := 0; i < len(keys); i++ {
		fmt.Printf("Status %d response: %%%d (%d)\n", keys[i], 100*statusCodes[keys[i]]/len(responses), statusCodes[keys[i]])
	}

	fmt.Printf("95 Percentile: %s\n", responses[p95].Timing.Total)
	fmt.Printf("99 Percentile: %s\n", responses[p99].Timing.Total)
	fmt.Printf("AverageTime: %s\n", averageTime)
}

var attackRequestCmd = &cobra.Command{
	Use:     "attack-request",
	Aliases: []string{"ar"},
	Short:   "http request",
	Run:     doAttackRequest,
	Args:    cobra.MinimumNArgs(0),
}
