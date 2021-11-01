package main

import (
	"os"

	"github.com/quantum-box/tachyon-sdk-go/internal/cli"
)

func main() {
	cli := cli.New(os.Stdout, os.Stderr)
	os.Exit(cli.Run(os.Args))
}
