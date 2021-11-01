package main

import (
	"flag"
	"fmt"
	"io"
)

const (
	ExitCodeOK             = 0
	ExitCodeParseFlagError = 1
)

type CLI struct {
	outStream, errStream io.Writer
}

func (c *CLI) Run(args []string) int {
	var version bool
	flags := flag.NewFlagSet("cms-cli", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.BoolVar(&version, "version", false, "Print version information and quit")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	if version {
		fmt.Fprintf(c.errStream, "cms-cli version %s", Version)
		return ExitCodeOK
	}
	fmt.Fprintf(c.outStream, "Do cms work")
	return ExitCodeOK
}
