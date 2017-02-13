package main

import (
	"flag"
	"fmt"
	"io"
	"strings"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// Opt structure
type Opt struct {
	Verbose     bool
	Version     bool
}

// CLI is the command line object
type CLI struct {
	outStream, errStream io.Writer
	opt                  Opt
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	c := DefaultConfig()

	flags.BoolVar(&cli.opt.Verbose, "verbose", false, "Print verbose log.")
	flags.BoolVar(&cli.opt.Version, "version", false, "Print version information and quit.")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	if cli.opt.Version {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	return cli.Cos()
}

// out
func (cli *CLI) out(format string, a ...interface{}) {
	if cli.opt.Verbose {
		fmt.Fprintln(cli.outStream, fmt.Sprintf(format, a...))
	}
}

// err
func (cli *CLI) err(format string, a ...interface{}) {
	fmt.Fprintln(cli.errStream, fmt.Sprintf(format, a...))
}

// Cos returns exit status
func (cli *CLI) Cos() int {
	c := DefaultConfig()
	c.Set(cli.opt)

	if IsFileExist(c.ConfigFile) {
		config, err := LoadConfig(c.ConfigFile)
		if err != nil {
			cli.err(fmt.Sprintf("%#v", err))
			return 1
		}
		c.Merge(config)
	}

	c.SetFromEnv()

	return 0
}

// IsCheckoutable returns bool
func (cli *CLI) IsCheckoutable() bool {
	return true
}
