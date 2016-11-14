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

// Ops structure
type Ops struct {
	TestCMD     string
	CoverageCMD string
	BaseBranch  string
	StatusName  string
	StatusOK    string
	StatusNG    string
	APIEndpoint string
	AccessToken string
	Comment     bool
	Verbose     bool
	Version     bool
}

// CLI is the command line object
type CLI struct {
	outStream, errStream io.Writer
	ops                  Ops
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	c := DefaultConfig()

	flags.StringVar(&cli.ops.TestCMD, "test-command", c.TestCMD, "Command to the Test")
	flags.StringVar(&cli.ops.TestCMD, "t", c.TestCMD, "Command to the Test(Short)")

	flags.StringVar(&cli.ops.CoverageCMD, "coverage-percentage-command", c.CoverageCMD, "Command to get the Coverage Percentage")
	flags.StringVar(&cli.ops.CoverageCMD, "p", c.CoverageCMD, "Command to get the Coverage Percentage(Short)")

	flags.StringVar(&cli.ops.BaseBranch, "base-branch", c.BaseBranch, "Base Branch")
	flags.StringVar(&cli.ops.BaseBranch, "b", c.BaseBranch, "Base Branch(Short)")

	flags.StringVar(&cli.ops.StatusName, "status-name", c.StatusName, "Status Name")
	flags.StringVar(&cli.ops.StatusName, "n", c.StatusName, "Status Name(Short)")

	flags.StringVar(&cli.ops.StatusOK, "status-ok", c.StatusOK, "Status Format for OK")
	flags.StringVar(&cli.ops.StatusNG, "status-ng", c.StatusNG, "Status Format for NG")

	flags.StringVar(&cli.ops.APIEndpoint, "api-endpoint", c.APIEndpoint, "API Endpoint")
	flags.StringVar(&cli.ops.APIEndpoint, "a", c.APIEndpoint, "API Endpoint(Short)")

	flags.StringVar(&cli.ops.AccessToken, "access-token", c.AccessToken, "Access Token")
	flags.StringVar(&cli.ops.AccessToken, "s", c.AccessToken, "Access Token(Short)")

	flags.BoolVar(&cli.ops.Comment, "comment", false, "Comment Coverage to the Pull-Request")
	flags.BoolVar(&cli.ops.Comment, "c", false, "(Short)")

	flags.BoolVar(&cli.ops.Verbose, "verbose", false, "Print verbose log.")
	flags.BoolVar(&cli.ops.Version, "version", false, "Print version information and quit.")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	if cli.ops.Version {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	if testCMD := flags.Args(); len(testCMD) != 0 {
		cli.ops.TestCMD = strings.Join(testCMD, " ")
	}

	return cli.Cos()
}

// out
func (cli *CLI) out(format string, a ...interface{}) {
	if cli.ops.Verbose {
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
	c.Set(cli.ops)

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
