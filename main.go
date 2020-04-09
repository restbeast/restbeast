package main

import (
	"gitlab.com/restbeast/cli/cmds"
)

var version string

func main() {
	cmds.Execute(version)
}
