package main

import (
	"os"

	"github.com/linyows/pipeline"
)

func main() {
	p := pipeline.New(os.Stdin, os.Stdout, os.Stderr)
	os.Exit(p.Run(os.Args))
}
