package main

import (
	"os"

	"github.com/wongnai/vac/internal/cli"
)

var version = ""

func main() {
	cli.Run(version, os.Args)
}
