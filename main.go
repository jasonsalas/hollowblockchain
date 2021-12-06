package main

import (
	"os"

	"github.com/jasonsalas/bc/cli"
)

func main() {
	defer os.Exit(0)
	cli := cli.CommandLine{}
	cli.Run()
}
