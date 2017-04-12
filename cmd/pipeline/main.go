package main

import (
	"os"

	"github.com/linyows/pipeline"
)

func main() {
	cli := pipeline.NewCLI(os.Stdout, os.Stderr, os.Stdin)
	os.Exit(cli.Run(os.Args))
}
