package main

import (
	"os"

	"github.com/joshvanl/gopem/cmd"
)

var (
	commit   = "unknown"
	date     = ""
	version  = "dev"
	wingHash = "unknown"
)

func main() {
	cmd.Version.Version = version
	cmd.Version.Commit = commit
	cmd.Version.BuildDate = date
	cmd.Execute(os.Args[1:])
}
