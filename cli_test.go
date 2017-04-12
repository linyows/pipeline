package pipeline

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestRunVersionFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := NewCLI(outStream, errStream, io.Stdin)
	args := strings.Split("pipeline --version", " ")

	status := cli.Run(args)
	if status != ExitCodeOK {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}

	expected := fmt.Sprintf("pipeline version %s", Version)
	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}
