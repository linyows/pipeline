package main

import (
	"flag"
	"fmt"
	"io"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		test    string
		percent string
		base    string
		format  string
		api     string
		comment bool
		verbose bool

		version bool
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.StringVar(&test, "test", "", "Test Command")
	flags.StringVar(&test, "t", "", "Test Command(Short)")

	flags.StringVar(&percent, "percent", "", "Covered Percent")
	flags.StringVar(&percent, "p", "", "Covered Percent(Short)")

	flags.StringVar(&base, "base", "", "Base Branch")
	flags.StringVar(&base, "b", "", "Base Branch(Short)")

	flags.StringVar(&format, "format", "", "Status Format")
	flags.StringVar(&format, "f", "", "Status Format(Short)")

	flags.StringVar(&api, "api", "", "API Endpoint")
	flags.StringVar(&api, "a", "", "API Endpoint(Short)")

	flags.BoolVar(&comment, "comment", false, "")
	flags.BoolVar(&comment, "c", false, "(Short)")
	flags.BoolVar(&verbose, "verbose", false, "")
	flags.BoolVar(&verbose, "v", false, "(Short)")

	flags.BoolVar(&version, "version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if version {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	_ = test

	_ = percent

	_ = base

	_ = format

	_ = api

	_ = comment

	_ = verbose

	return ExitCodeOK
}
