package main

import (
	"bufio"
	. "fmt"
	"github.com/getsentry/sentry-go"
	"github.com/go-errors/errors"
	"gitlab.com/restbeast/cli/cmds"
	"gitlab.com/restbeast/cli/lib"
	"log"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

var version, sentryDsn string

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              sentryDsn,
		Release:          version,
		AttachStacktrace: true,
	})

	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	execCtx := lib.ExecutionContext{
		Version: version,
		Debug:   os.Getenv("DEBUG") != "",
	}

	defer sentry.Flush(2 * time.Second)

	defer func() {
		if r := recover(); r != nil {
			Println("RestBeast encountered an unknown error")

			if execCtx.Debug {
				log.Printf("%s", r)
				log.Printf("%s", debug.Stack())
			}

			stdinReader := bufio.NewReader(os.Stdin)
			Print("Do you want to send crash report [y/N]: ")
			choice, _ := stdinReader.ReadString('\n')

			if strings.Trim(strings.ToLower(choice), "\t \n") == "y" {
				sentry.WithScope(func(scope *sentry.Scope) {
					scope.SetLevel(sentry.LevelFatal)
					sentry.CaptureException(errors.Wrap(r, 4))
				})
				Println("Crash report sent.")
			}
		}
	}()

	cmds.Execute(&execCtx)
}
