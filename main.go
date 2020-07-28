package main

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/go-errors/errors"
	"gitlab.com/restbeast/cli/cmds"
	"log"
	"time"
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

	defer sentry.Flush(2 * time.Second)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("restbeast encountered an unknown error")
			sentry.WithScope(func(scope *sentry.Scope) {
				scope.SetLevel(sentry.LevelFatal)
				sentry.CaptureException(errors.Wrap(r,4))
			})
		}
	}()

	cmds.Execute(version)
}