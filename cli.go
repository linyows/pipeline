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

const escape = "\x1b"

var blue = fmt.Sprintf("%s[%sm", escape, "1;34")
var clear = fmt.Sprintf("%s[%sm", escape, "0")

var logo = `
╔═╗╦╔═╗╔═╗╦  ╦╔╗╔╔═╗
╠═╝║╠═╝║╣ ║  ║║║║║╣
╩  ╩╩  ╚═╝╩═╝╩╝╚╝╚═╝
`

var usageText = `
Usage: pipeline [options] [args]

Options:`

var exampleText = `
Examples:
  $ pipeline --config /etc/pipeline.yml

`

// CLI is structure
type CLI struct {
	outStream, errStream io.Writer
	inStream             *os.File
}

// NewCLI returns CLI struct
func NewCLI(o io.Writer, e io.Writer, i *os.File) *CLI {
	return &CLI{
		outStream: o,
		errStream: e,
		inStream:  i,
	}
}

// Run invokes the CLI with the given arguments.
func (c *CLI) Run(args []string) int {
	f := flag.NewFlagSet(Name, flag.ContinueOnError)
	f.SetOutput(c.outStream)

	f.Usage = func() {
		fmt.Fprintf(c.outStream, blue)
		fmt.Fprintf(c.outStream, logo)
		fmt.Fprintf(c.outStream, clear)
		fmt.Fprintf(c.outStream, usageText)
		f.PrintDefaults()
		fmt.Fprint(c.outStream, exampleText)
	}

	var opt Options

	f.StringVar(&opt.Config, []string{"c", "-config"}, "/etc/pipeline.yml", "the path to the configuration file")
	f.BoolVar(&opt.Version, []string{"v", "-version"}, false, "print the version and exit")

	if err := f.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	//parsedArgs := f.Args()

	if opt.Version {
		fmt.Fprintf(c.outStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	f.Usage()

	p := NewPipeline()
	p.Run(f.Args())

	return ExitCodeOK
}
