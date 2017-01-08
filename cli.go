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
	opt                  Opt
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	c := DefaultConfig()

	flags.StringVar(&cli.opt.TestCMD, "test-command", c.TestCMD, "Command to the Test")
	flags.StringVar(&cli.opt.TestCMD, "t", c.TestCMD, "Command to the Test(Short)")

	flags.StringVar(&cli.opt.CoverageCMD, "coverage-percentage-command", c.CoverageCMD, "Command to get the Coverage Percentage")
	flags.StringVar(&cli.opt.CoverageCMD, "p", c.CoverageCMD, "Command to get the Coverage Percentage(Short)")

	flags.StringVar(&cli.opt.BaseBranch, "base-branch", c.BaseBranch, "Base Branch")
	flags.StringVar(&cli.opt.BaseBranch, "b", c.BaseBranch, "Base Branch(Short)")

	flags.StringVar(&cli.opt.StatusName, "status-name", c.StatusName, "Status Name")
	flags.StringVar(&cli.opt.StatusName, "n", c.StatusName, "Status Name(Short)")

	flags.StringVar(&cli.opt.StatusOK, "status-ok", c.StatusOK, "Status Format for OK")
	flags.StringVar(&cli.opt.StatusNG, "status-ng", c.StatusNG, "Status Format for NG")

	flags.StringVar(&cli.opt.APIEndpoint, "api-endpoint", c.APIEndpoint, "API Endpoint")
	flags.StringVar(&cli.opt.APIEndpoint, "a", c.APIEndpoint, "API Endpoint(Short)")

	flags.StringVar(&cli.opt.AccessToken, "access-token", c.AccessToken, "Access Token")
	flags.StringVar(&cli.opt.AccessToken, "s", c.AccessToken, "Access Token(Short)")

	flags.BoolVar(&cli.opt.Comment, "comment", false, "Comment Coverage to the Pull-Request")
	flags.BoolVar(&cli.opt.Comment, "c", false, "(Short)")

	flags.BoolVar(&cli.opt.Verbose, "verbose", false, "Print verbose log.")
	flags.BoolVar(&cli.opt.Version, "version", false, "Print version information and quit.")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	if cli.opt.Version {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	if testCMD := flags.Args(); len(testCMD) != 0 {
		cli.opt.TestCMD = strings.Join(testCMD, " ")
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
