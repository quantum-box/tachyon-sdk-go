package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestRun_versionFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("cms-cli -version", " ")

	status := cli.Run(args)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus=%d, want %d", status, ExitCodeOK)
	}

	expected := fmt.Sprintf("cms-cli version %s", Version)
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("Output=%q, want %q", errStream.String(), expected)
	}
}
