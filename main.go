package main

import (
	"gitlab.com/restbeast/cli/cmds"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
)

var version, sentryDsn string

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: sentryDsn,
		Release: version,
		AttachStacktrace: true,
	})

	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)
	defer sentry.Recover()

	cmds.Execute(version)
}
