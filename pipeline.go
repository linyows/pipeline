package pipeline

import (
	"fmt"
	"io"
	"os"

	flag "github.com/linyows/mflag"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// Options is structure
type Options struct {
	Config  string
	Version bool
}

var usageText = `
Usage: pipeline [options] [args]

Options:`

var exampleText = `
Examples:
  $ pipeline --config /etc/pipeline.conf

`

// Pipeline is structure
type Pipeline struct {
	outStream, errStream io.Writer
	inStream             *os.File
}

// New for pipeline
func New(stdin *os.File, stdout io.Writer, stderr io.Writer) *Pipeline {
	return &Pipeline{
		inStream:  stdin,
		outStream: stdout,
		errStream: stderr,
	}
}

// Run invokes the CLI with the given arguments.
func (p *Pipeline) Run(args []string) int {
	f := flag.NewFlagSet(Name, flag.ContinueOnError)
	f.SetOutput(p.outStream)

	f.Usage = func() {
		fmt.Fprintf(p.outStream, usageText)
		f.PrintDefaults()
		fmt.Fprint(p.outStream, exampleText)
	}

	var opt Options

	f.StringVar(&opt.Config, []string{"c", "-config"}, "/etc/pipeline.conf", "the path to the configuration file")
	f.BoolVar(&opt.Version, []string{"v", "-version"}, false, "print the version and exit")

	if err := f.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	//parsedArgs := f.Args()

	if opt.Version {
		fmt.Fprintf(p.outStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	return ExitCodeOK
}
