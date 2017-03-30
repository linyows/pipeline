package main

import (
	"github.com/linyows/pipeline"
)

func main() {
	flag.Usage = usage
	flag.Parse()
	log.SetFlags(0)

	args := flag.Args()
	if len(args) < 1 {
		usage()
	}

	if args[0] == "help" {
		help(args[1:])
		return
	}

	for _, cmd := range commands {
		if cmd.Name() == args[0] {
			cmd.Flag.Usage = func() { cmd.Usage() }

			cmd.Flag.Parse(args[1:])
			args = cmd.Flag.Args()

			os.Exit(cmd.Run(args))
		}
	}

	fmt.Fprintf(os.Stderr, "pipeline: unknown subcommand %q\nRun \" pipeline help\" for usage.\n", args[0])
	os.Exit(2)
}
