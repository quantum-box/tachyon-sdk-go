package cli

import (
	"flag"
	"fmt"
	tachyon "github.com/quantum-box/tachyon-sdk-go"
	"io"
)

const (
	ExitCodeOK             = 0
	ExitCodeParseFlagError = 1
)

type CLI struct {
	outStream, errStream io.Writer
}

func New(outStream, errStream io.Writer) *CLI {
	return &CLI{
		outStream: outStream,
		errStream: errStream,
	}
}

func (c *CLI) Run(args []string) int {
    fmt.Println(args)
	var version bool
	flags := flag.NewFlagSet("cms", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.BoolVar(&version, "version", false, "Print version information and quit")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	if version {
		fmt.Fprintf(c.errStream, "cms-cli version %s", tachyon.Version)
		return ExitCodeOK
	}
	fmt.Fprintf(c.outStream, "Do cms work")
	return ExitCodeOK
}
