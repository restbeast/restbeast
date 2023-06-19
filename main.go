package main

import (
	. "fmt"
	"log"
	"os"
	"runtime/debug"

	"github.com/restbeast/restbeast/cmds"
	"github.com/restbeast/restbeast/lib"
)

var version string

func main() {
	execCtx := lib.ExecutionContext{
		Version: version,
		Debug:   os.Getenv("DEBUG") != "",
	}

	defer func() {
		if r := recover(); r != nil {
			Println("RestBeast encountered an unknown error")

			if execCtx.Debug {
				log.Printf("%s", r)
				log.Printf("%s", debug.Stack())
			}
		}
	}()

	cmds.Execute(&execCtx)
}
