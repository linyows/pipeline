package pipeline

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

// A Command is an implementation of a pipeline command
type Command struct {
	Run func(args []string) int
	UsageLine string
	Short string
	Long string
	Flag flag.FlagSet
}

// Name returns the command"s name: the first word in the usage line.
func (c *Command) Name() string {
	name := c.UsageLine
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

// Usage returns usage
func (c *Command) Usage() {
	fmt.Fprintf(os.Stderr, "usage: %s\n\n", c.UsageLine)
	fmt.Fprintf(os.Stderr, "%s\n", strings.TrimSpace(c.Long))
	os.Exit(2)
}

// Commands lists the available commands and help topics.
var commands = []*Command{
	cmdInit,
	cmdStart,
	cmdList,
}

var usageTemplate = `
Usage: pipeline [options] CMD

Commands:{{range .}}
  {{.Name | printf "%-11s"}} {{.Short}}{{end}}

Options:
  -c,  --config
	-V,  --verbose

Use "pipeline help [command]" for more information about a command.

`

var helpTemplate = `usage: pipeline {{.UsageLine}}

{{.Long | trim}}
`

// tmpl executes the given template text on data, writing the result to w.
func tmpl(w io.Writer, text string, data interface{}) {
	t := template.New("top")
	t.Funcs(template.FuncMap{"trim": strings.TrimSpace})
	template.Must(t.Parse(text))
	if err := t.Execute(w, data); err != nil {
		panic(err)
	}
}

func printUsage(w io.Writer) {
	bw := bufio.NewWriter(w)
	tmpl(bw, usageTemplate, commands)
	bw.Flush()
}

func usage() {
	printUsage(os.Stderr)
	os.Exit(2)
}

// Help implements the "help" command.
func Help(args []string) {
	if len(args) == 0 {
		printUsage(os.Stdout)
		return
	}
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "usage: pipeline help command\n\nToo many arguments given.\n")
		os.Exit(2)
	}

	arg := args[0]

	for _, cmd := range commands {
		if cmd.Name() == arg {
			tmpl(os.Stdout, helpTemplate, cmd)
			return
		}
	}

	fmt.Fprintf(os.Stderr, "Unknown help topic %#q.  Run \"pipeline help\".\n", arg)
	os.Exit(2)
}
